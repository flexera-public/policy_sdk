package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/robertkrimen/otto"
	// This import loads the otto runtime with underscore library enabled.
	_ "github.com/robertkrimen/otto/underscore"
)

var (
	maxExecTime      = 600 * time.Second
	errHalt          = fmt.Errorf("HALT")
	debuglogDataSize = 10 * 1024
)

type script struct {
	name   string
	code   string
	result string
	params []string
}

type param struct {
	name  string
	value interface{}
}

func runScript(ctx context.Context, file, outfile, result, name string, options []string) error {
	rd, err := os.Open(file)
	if err != nil {
		return err
	}
	srcBytes, err := ioutil.ReadAll(rd)
	if err != nil {
		return err
	}
	src := string(srcBytes)

	wr, err := os.Create(outfile)
	if err != nil {
		return err
	}

	params, err := parseParams(options)
	if err != nil {
		return err
	}

	var data interface{}
	if strings.Contains(src, "rs_pt_ver") {
		scripts := getScripts(src)
		// for i, s := range scripts {
		// 	fmt.Printf("DBG:%d:%+v\n", i, s)
		// }

		if len(scripts) == 0 {
			return fmt.Errorf("no script blocks found")
		} else if len(scripts) == 1 {

			src = scripts[0].code
			result = scripts[0].result
		} else if len(scripts) > 0 {
			names := []string{}
			found := false
			for _, s := range scripts {
				names = append(names, s.name)
				if s.name == name {
					src = s.code
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
	}
	if name != "" {
		name += " "
	}
	fmt.Printf("Running %sscript from %s and writing %s to %s\n", name, file, result, outfile)

	data, err = execScript(src, params, result)

	if err != nil {
		return err
	}

	enc := json.NewEncoder(wr)
	enc.SetIndent("", "  ")
	err = enc.Encode(&data)
	if err != nil {
		return err
	}

	return nil
}

func execScript(code string, params []*param, result string) (out interface{}, err error) {
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
	stringifyArgs := func(prefix string, args []otto.Value) string {
		output := []byte{}
		basePrefix := prefix
		suffix := "\n"
		for _, arg := range args {
			if !arg.IsPrimitive() {
				prefix = basePrefix + " >\n"
				suffix = "\n"
			}
			v, _ := arg.Export()
			b, _ := json.MarshalIndent(v, "  ", "  ")
			output = append(output, ' ')
			output = append(output, b...)
			if len(output) > debuglogDataSize {
				break
			}
		}
		if len(output) > debuglogDataSize {
			left := len(output) - debuglogDataSize
			return fmt.Sprintf("%s%s ... %d bytes omitted %s",
				prefix, string(output[:debuglogDataSize]), left, suffix)
		}
		return prefix + string(output) + suffix
	}
	logFn := func(kind string) func(call otto.FunctionCall) otto.Value {
		return func(call otto.FunctionCall) otto.Value {
			fmt.Printf(stringifyArgs("console.%s:", call.ArgumentList), kind)
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

func parseParams(runOptions []string) ([]*param, error) {
	params := []*param{}
	for _, o := range runOptions {
		bits := strings.SplitN(o, "=", 2)
		name := bits[0]
		valStr := bits[1]
		val, err := parseVal(valStr)
		if err != nil {
			return nil, err
		}
		params = append(params, &param{name, val})
	}
	return params, nil
}

func parseVal(v string) (interface{}, error) {
	if strings.HasPrefix(v, "@") {
		v = strings.TrimPrefix(v, "@")
		rd, err := os.Open(v)
		if err != nil {
			return nil, err
		}
		var data interface{}
		err = json.NewDecoder(rd).Decode(&data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	if valN, err := coerceOption("", v, "number"); err != nil {
		return valN, nil
	}
	return v, nil
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
			// it is unlikely this will be a problem, but if we do run into it, we could come up with a hueristic to
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
	qw := regexp.MustCompile(`"(.*?[^\\])"|'(.*?[^\\])'`)
	scriptStartRe := regexp.MustCompile(`^\s*script ['"]([^'"]+)['"],.*do\s*$`)
	scriptEndRe := regexp.MustCompile(`^\s*end\s*$`)
	codeStartRe := regexp.MustCompile(`^\s*code ('.*|".*|<<-?[A-Z_]+\s*)$`)
	var codeEndRe *regexp.Regexp

	scriptLines := []string{}
	codeLines := []string{}
	name := ""
	for _, line := range lines {
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
					params = append(params, m[1])
				}
				scripts = append(scripts, &script{
					name:   name,
					code:   unquote(strings.Join(codeLines, "\n")),
					result: unquote(result),
					params: params,
				})
			} else if matches := codeStartRe.FindStringSubmatch(line); len(matches) > 0 {
				//fmt.Println("DBG HERE_CODE START")
				inCode = true
				if strings.HasPrefix(matches[1], "'") {
					codeEndRe = regexp.MustCompile(`^(.*?[^\\]')`)
					codeLines = append(codeLines, matches[1])
				} else if strings.HasPrefix(matches[1], `"`) {
					codeEndRe = regexp.MustCompile(`^(.*?[^\\]")`)
					codeLines = append(codeLines, matches[1])
				} else {
					end := strings.TrimPrefix(strings.TrimPrefix(matches[1], "<<"), "-")
					codeEndRe = regexp.MustCompile(`^\s*` + end)
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
	keyRe := regexp.MustCompile(fmt.Sprintf(`\s*%s\s+(.*)\s*$`, key))
	for _, line := range lines {
		if matches := keyRe.FindStringSubmatch(line); len(matches) > 0 {
			return matches[1]
		}
	}
	return ""
}

func unquote(s string) string {
	q := regexp.MustCompile(`\\(.)`)
	if strings.HasPrefix(s, "'") {
		s = s[1 : len(s)-1]
		s = q.ReplaceAllString(s, "$1")
	} else if strings.HasPrefix(s, `"`) {
		s = s[1 : len(s)-1]
		s = q.ReplaceAllString(s, "$1")
	}
	return s
}
