package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	//"github.com/pkg/errors"
	//"regexp"

	"github.com/robertkrimen/otto"
	// This import loads the otto runtime with underscore library enabled.
	_ "github.com/robertkrimen/otto/underscore"
)

var (
	maxExecTime = 600 * time.Second
	errHalt     = fmt.Errorf("HALT")
)

type script struct {
	name   string
	code   string
	result string
	params *param
}

type param struct {
	name  string
	value interface{}
}

func runScript(ctx context.Context, file, outfile, result string, options []string) error {
	fmt.Printf("Running script from %s and writing result to %s\n", file, outfile)
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
		return fmt.Errorf("TBD")
	} else {
		// Assume a javascript was passed in
		data, err = execScript(src, params, result)
	}
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

// func getScripts(src []byte) ([]*script, error) {
// 	return []*script{
// 		{},
// 	}
// 	t := newTokenizer(src)
// 	scripts := []*script{}
// 	tokens, err := t.between("script", "end")
// 	if err == io.EOF {
// 		return scripts, nil
// 	}
// 	strRe := `"([^"]*[^\])"|'([^']*[^\])'|<<-?([A-Z]+).*($4)`

// }

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
		suffix := ""
		for _, arg := range args {
			if !arg.IsPrimitive() {
				prefix = basePrefix + " >\n"
				suffix = "\n"
			}
			v, _ := arg.Export()
			b, _ := json.MarshalIndent(v, "  ", "  ")
			output = append(output, ' ')
			output = append(output, b...)
			// if len(output) > debuglogDataSize {
			// 	break
			// }
		}
		// if len(output) > debuglogDataSize {
		// 	return prefix + string(output[:debuglogDataSize]) + "..." + suffix
		// }
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
	fmt.Printf("JavaScript finished, duration=%s", time.Since(start).String())
	if err != nil {
		return nil, err
	}
	r, err := vm.Get(result)
	if err != nil {
		fmt.Errorf("failed to retrieve value of %q", result)
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
