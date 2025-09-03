package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"

	// This import loads the otto runtime with underscore library enabled.
	_ "github.com/robertkrimen/otto/underscore"
)

var (
	errHalt          = fmt.Errorf("HALT")
	debuglogDataSize = 10 * 1024
	q                = regexp.MustCompile(`\\(.)`)
	qw               = regexp.MustCompile(`"(.*?[^\\])"|'(.*?[^\\])'`)
	scriptStartRe    = regexp.MustCompile(`^\s*script\s+['"]([^'"]+)['"],.*do\s*(?:#.*)?$`)
	scriptEndRe      = regexp.MustCompile(`^\s*end\s*(?:#.*)?$`)
	codeStartRe      = regexp.MustCompile(`^\s*code\s+('.*|".*|<<-?(?:[A-Za-z_]+|'[A-Za-z_]+'|"[A-Za-z_]+")\s*(?:#.*)?)$`)
)

type script struct {
	name   string
	code   string
	line   int
	result string
	params []string
}

type param struct {
	name  string
	value interface{}
}

type scriptErrorList struct {
	errs parser.ErrorList
	file string
	line int
}

func (s *scriptErrorList) Error() string {
	r := regexp.MustCompile(`^(` + regexp.QuoteMeta(s.file) + `: Line )(\d+):`)
	errs := make([]string, len(s.errs))
	for i, err := range s.errs {
		e := err.Error()
		matches := r.FindStringSubmatch(e)
		line, _ := strconv.Atoi(matches[2])
		errs[i] = r.Copy().ReplaceAllString(e, fmt.Sprintf("%v%v:", matches[1], s.line-1+line))
	}
	return strings.Join(errs, "\n")
}

func runScript(ctx context.Context, file, outfile, result, name string, maxExecTime int, options []string, assumeDatasourceJSON bool) error {
	rd, err := os.Open(file)
	if err != nil {
		return err
	}
	srcBytes, err := ioutil.ReadAll(rd)
	if err != nil {
		return err
	}
	src := normalizeLineEndings(string(srcBytes))

	maxExecTimeSeconds := time.Duration(maxExecTime) * time.Second

	// Enhance options with automatic datasource JSON detection if enabled
	if assumeDatasourceJSON {
		options = enhanceOptionsWithDatasourceJSON(src, options)
	}

	params, err := parseParams(options)
	if err != nil {
		return err
	}

	var (
		data interface{}
		line int
	)
	if strings.Contains(src, "rs_pt_ver") {
		scripts := getScripts(src)
		// for i, s := range scripts {
		// 	fmt.Printf("DBG:%d:%+v\n", i, s)
		// }

		if len(scripts) == 0 {
			return fmt.Errorf("no script blocks found")
		} else if len(scripts) == 1 {
			name = scripts[0].name
			src = scripts[0].code
			line = scripts[0].line
			result = scripts[0].result
		} else if len(scripts) > 0 {
			names := []string{}
			found := false
			for _, s := range scripts {
				names = append(names, s.name)
				if s.name == name {
					src = s.code
					line = s.line
					result = s.result
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("multiple script blocks found, pass --name %s to specify which one to run",
					strings.Join(names, " or "))
			}
		}
		if result == "" {
			return fmt.Errorf("no result variable specified in script block %#v", name)
		}
	} else if result == "" {
		return fmt.Errorf("no result variable specified for raw JavaScript, pass --result with the name of the variable to extract")
	} else {
		name = strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
	}
	fmt.Printf("Running script %q from %s and writing %s to %s\n", name, file, result, outfile)

	code, err := parser.ParseFile(nil, file, src, 0)
	if err != nil {
		switch e := err.(type) {
		case parser.ErrorList:
			err = &scriptErrorList{errs: e, file: file, line: line}
		}
		return err
	}

	data, err = execScript(code, params, maxExecTimeSeconds, result)
	if err != nil {
		return err
	}

	wr, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer wr.Close()

	enc := json.NewEncoder(wr)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err = enc.Encode(&data)
	if err != nil {
		return err
	}

	return nil
}

func execScript(code *ast.Program, params []*param, maxExecTime time.Duration, result string) (out interface{}, err error) {
	defer func() {
		if caught := recover(); caught != nil {
			if caught == errHalt {
				err = fmt.Errorf("script execution exceeded %ds", int(maxExecTime.Seconds()))
				return
			}
			err = fmt.Errorf("unknown error: %v", caught)
		}
	}()

	vm := otto.New()
	err = registerJsFunctions(vm)
	if err != nil {
		return nil, err
	}
	stringifyArgs := func(prefix string, args []otto.Value) string {
		output := &strings.Builder{}
		basePrefix := prefix
		separator := " "
		primitive := true
		for _, arg := range args {
			if !arg.IsPrimitive() {
				prefix = basePrefix + " >\n"
				separator = "  "
				primitive = false
				break
			}
		}
		for _, arg := range args {
			v, _ := arg.Export()

			b := &strings.Builder{}
			e := json.NewEncoder(b)
			e.SetEscapeHTML(false)
			e.SetIndent("  ", "  ")
			e.Encode(v)
			s := b.String()
			if primitive {
				s = strings.TrimSuffix(s, "\n")
			}

			output.WriteString(separator)
			output.WriteString(s)
			if output.Len() > debuglogDataSize {
				break
			}
		}
		if output.Len() > debuglogDataSize {
			left := output.Len() - debuglogDataSize
			return fmt.Sprintf("%s%s ... %d bytes omitted", prefix, output.String()[:debuglogDataSize], left)
		}
		return prefix + strings.TrimSuffix(output.String(), "\n")
	}
	logFn := func(kind string) func(call otto.FunctionCall) otto.Value {
		return func(call otto.FunctionCall) otto.Value {
			fmt.Println(stringifyArgs(fmt.Sprintf("console.%s:", kind), call.ArgumentList))
			return otto.UndefinedValue()
		}
	}
	vm.Set("console", map[string]interface{}{
		"log": logFn("log"),
		"dir": logFn("dir"),
	})

	vm.Interrupt = make(chan func(), 1) // The buffer prevents blocking
	start := time.Now()
	timer := time.AfterFunc(maxExecTime, func() {
		vm.Interrupt <- func() {
			panic(errHalt)
		}
	})

	for _, p := range params {
		vm.Set(p.name, p.value)
	}

	_, err = vm.Run(code)
	timer.Stop()
	fmt.Printf("JavaScript finished, duration=%s\n", time.Since(start).String())
	if err != nil {
		return nil, err
	}
	r, err := vm.Get(result)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve value of %q", result)
	}
	return export(r)
}

// enhanceOptionsWithDatasourceJSON checks for datasource_{datasourcename}.json files
// and automatically uses them for script parameters if they exist and no explicit value is provided
func enhanceOptionsWithDatasourceJSON(src string, options []string) []string {
	// Parse existing options to see what's already provided
	providedParams := make(map[string]bool)
	for _, option := range options {
		if strings.Contains(option, "=") {
			paramName := strings.SplitN(option, "=", 2)[0]
			providedParams[paramName] = true
		}
	}

	// Extract script parameters from the policy template
	scriptParams := extractScriptParameters(src)

	// Build enhanced options list
	enhancedOptions := make([]string, len(options))
	copy(enhancedOptions, options)

	// For each script parameter that wasn't explicitly provided,
	// check if a corresponding datasource_{paramname}.json file exists
	for _, paramName := range scriptParams {
		if !providedParams[paramName] {
			dsFilename := fmt.Sprintf("datasource_%s.json", paramName)
			// print some debug information about the path we are expecting
			fmt.Printf("Expecting datasource file: %s\n", dsFilename)
			if _, err := os.Stat(dsFilename); err == nil {
				// File exists, add it to options
				enhancedOption := fmt.Sprintf("%s=@%s", paramName, dsFilename)
				enhancedOptions = append(enhancedOptions, enhancedOption)
				// fmt.Printf("Auto-detected datasource file: %s for parameter '%s'\n", dsFilename, paramName)
			}
		}
	}

	return enhancedOptions
}

// extractScriptParameters extracts parameter names from script blocks in a policy template
func extractScriptParameters(src string) []string {
	var allParams []string

	// Check if this is a policy template (contains rs_pt_ver)
	if strings.Contains(src, "rs_pt_ver") {
		scripts := getScripts(src)
		for _, script := range scripts {
			allParams = append(allParams, script.params...)
		}
	} else {
		// For raw JavaScript, we can't easily extract parameter names,
		// so we return an empty list (no auto-detection for raw JS)
		return []string{}
	}

	// Remove duplicates
	paramMap := make(map[string]bool)
	var uniqueParams []string
	for _, param := range allParams {
		if !paramMap[param] {
			paramMap[param] = true
			uniqueParams = append(uniqueParams, param)
		}
	}

	return uniqueParams
}

func parseParams(runOptions []string) ([]*param, error) {
	params := []*param{}
	for _, o := range runOptions {
		bits := strings.SplitN(o, "=", 2)
		name := bits[0]
		valStr := bits[1]
		val, err := parseVal(name, valStr)
		if err != nil {
			return nil, err
		}
		params = append(params, &param{name, val})
	}
	return params, nil
}

func parseVal(name, v string) (interface{}, error) {
	if strings.HasPrefix(v, "@") {
		v = strings.TrimPrefix(v, "@")
		rd, err := os.Open(v)
		if err != nil {
			return nil, err
		}
		var data interface{}
		err = json.NewDecoder(rd).Decode(&data)
		if err != nil {
			return nil, errors.Wrapf(err, "parameter %s: failed to read %q as json", name, v)
		}
		return data, nil
	}
	if looksLikeJSON(v) {
		var valJson interface{}
		err := json.Unmarshal([]byte(v), &valJson)
		if err != nil {
			return nil, errors.Wrapf(err, "parameter %s: failed to parse %q as json", name, v)
		}
		return valJson, nil
	} else if valN, err := coerceOption(name, v, "number"); err == nil {
		return valN, nil
	}
	return v, nil
}

func looksLikeJSON(s string) bool {
	return strings.HasPrefix(s, "[") || strings.HasPrefix(s, "{") || strings.HasPrefix(s, `"`)
}

// export converts an otto.Value into a predictable Go representation. The otto.Value Export function can return more
// types that require more type conversion further on, so we use this instead.
func export(v otto.Value) (interface{}, error) {
	switch {
	case v.IsBoolean():
		// boolean -> bool
		return v.ToBoolean()
	case v.IsNumber():
		// number -> float64
		return v.ToFloat()
	case v.IsString():
		// string -> string
		return v.ToString()
	case v.IsObject():
		o := v.Object()
		switch o.Class() {
		case "Array", "GoArray":
			// array -> []interface{}
			lv, err := o.Get("length")
			if err != nil {
				return nil, err
			}
			l, err := lv.ToInteger()
			if err != nil {
				return nil, err
			}
			// Keys will have at least as many actual elements as are in the array since JavaScript allows sparse arrays
			a := make([]interface{}, 0, len(o.Keys()))
			// this algorithm is NOT optimal for sparse arrays where the length is much greater than the amount of
			// actual elements in the array since it will be doing a lot of IsDefined checks for no good reason.
			// there probably aren't that many use-cases where sparse arrays would make sense for datasource results, so
			// it is unlikely this will be a problem, but if we do run into it, we could come up with a heuristic to
			// determine if it would be better to use the alternative algorithm of sorting the slice of Keys after
			// parsing them as integers and then including values from Keys that parse as integers and are greater than
			// or equal to 0 and less than the array length
			for i := int64(0); i < l; i++ {
				kv, err := o.Get(strconv.FormatInt(i, 10))
				if err != nil {
					return nil, err
				}
				if kv.IsDefined() {
					ekv, err := export(kv)
					if err != nil {
						return nil, err
					}
					a = append(a, ekv)
				}
			}
			return a, nil
		case "Object":
			// object -> map[string]interface{}
			m := make(map[string]interface{})
			for _, k := range o.Keys() {
				kv, err := o.Get(k)
				if err != nil {
					return nil, err
				}
				m[k], err = export(kv)
				if err != nil {
					return nil, err
				}
			}
			return m, nil
		default:
			// function, etc. -> map[string]interface{}
			return map[string]interface{}{}, nil
		}
	default:
		// undefined or null -> nil
		return nil, nil
	}
}

// getScripts extracts the script blocks from a PolicyTemplate. Would be nice to
// get this from compilation service as well but not sure if there's an easy way
// to do that.
func getScripts(src string) []*script {
	scripts := []*script{}
	lines := strings.Split(src, "\n")
	inScript := false
	inCode := false
	unescapeCode := true
	var codeEndRe *regexp.Regexp
	var codeLine int

	scriptLines := []string{}
	codeLines := []string{}
	name := ""
	for i, line := range lines {
		if inCode {
			//fmt.Println("DBG HERE_CODE", line)
			if matches := codeEndRe.FindStringSubmatch(line); len(matches) > 0 {
				if len(matches) > 1 {
					codeLines = append(codeLines, matches[1])
				}
				//fmt.Println("DBG HERE_CODE END")
				inCode = false
			} else {
				codeLines = append(codeLines, line)
			}
		} else if inScript {
			//fmt.Println("DBG HERE_SCR", line)
			if scriptEndRe.MatchString(line) {
				//fmt.Println("DBG HERE_SCR END")
				inScript = false
				result := getKey(scriptLines, "result")
				rawParams := getKey(scriptLines, "parameters")
				matches := qw.FindAllStringSubmatch(rawParams, -1)
				params := []string{}
				for _, m := range matches {
					if m[1] != "" {
						params = append(params, m[1])
					} else {
						params = append(params, m[2])
					}
				}
				scripts = append(scripts, &script{
					name:   name,
					code:   unquote(strings.Join(codeLines, "\n"), unescapeCode),
					line:   codeLine,
					result: unquote(result, true),
					params: params,
				})
			} else if matches := codeStartRe.FindStringSubmatch(line); len(matches) > 0 {
				//fmt.Println("DBG HERE_CODE START")
				inCode = true
				codeLine = i + 1
				if strings.HasPrefix(matches[1], "'") {
					codeEndRe = regexp.MustCompile(`^(.*?[^\\]')`)
					codeLines = append(codeLines, matches[1])
				} else if strings.HasPrefix(matches[1], `"`) {
					codeEndRe = regexp.MustCompile(`^(.*?[^\\]")`)
					codeLines = append(codeLines, matches[1])
				} else {
					identifier := strings.TrimSpace(strings.SplitN(strings.TrimPrefix(strings.TrimPrefix(matches[1], "<<"), "-"), "#", 2)[0])
					unescapeCode = identifier[:1] != "'"
					end := strings.Trim(identifier, `'"`)
					codeEndRe = regexp.MustCompile(`^\s*` + end)
					codeLine++
				}
			} else {
				scriptLines = append(scriptLines, line)
			}
		} else {
			if matches := scriptStartRe.FindStringSubmatch(line); len(matches) > 0 {
				name = matches[1]
				inScript = true
				scriptLines = []string{}
				codeLines = []string{}
			}
		}
	}
	return scripts
}

func getKey(lines []string, key string) string {
	keyRe := regexp.MustCompile(fmt.Sprintf(`\s*%s\s+([^#]*)(?:#.*)?$`, key))
	for _, line := range lines {
		if matches := keyRe.FindStringSubmatch(line); len(matches) > 0 {
			return strings.TrimSpace(matches[1])
		}
	}
	return ""
}

func unquote(s string, unescape bool) string {
	if strings.HasPrefix(s, "'") {
		s = s[1 : len(s)-1]
	} else if strings.HasPrefix(s, `"`) {
		s = s[1 : len(s)-1]
	}
	if unescape {
		s = q.ReplaceAllString(s, "$1")
	}
	return s
}

func normalizeLineEndings(in string) string {
	out := strings.Replace(in, "\r\n", "\n", -1)
	return strings.Replace(out, "\r", "\n", -1)
}

func jsBtoa(b string) string {
	return base64.StdEncoding.EncodeToString([]byte(b))
}

func jsAtob(str string) string {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		v, _ := otto.ToValue(fmt.Sprintf("DOMExpection [InvalidCharacterError]: Invalid character: %v", err))
		panic(v)
	}
	return string(b)
}

func registerJsFunctions(vm *otto.Otto) (err error) {
	err = vm.Set("btoa", jsBtoa)
	if err != nil {
		return err
	}
	err = vm.Set("atob", jsAtob)
	if err != nil {
		return err
	}
	return
}
