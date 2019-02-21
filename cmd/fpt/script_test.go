package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestGetScripts(t *testing.T) {
	tests := []struct {
		name string
		exp  []*script
	}{
		{
			name: "get_script1.pt",
			exp: []*script{
				{
					name:   "do_filter",
					code:   "var filtered_datacenters = [];\nvar matchers = []\nvar end = [\"a\"]\n\nconsole.log(datacenters.length)\nfiltered_datacenters = [\"foo\"]",
					params: []string{"datacenters", "testing_one"},
					result: "filtered_datacenters",
				},
				{
					name:   "otra",
					code:   "console.log(otra)\nRESVAR=otra[0]",
					params: []string{"otra"},
					result: "RESVAR",
				},
			},
		},
	}

	for _, tt := range tests {
		src := readFile(tt.name)
		got := getScripts(src)
		if len(got) != len(tt.exp) {
			t.Fatalf("Expected %d scripts, got %d", len(tt.exp), len(got))
		}
		for j, expS := range tt.exp {
			gotS := got[j]
			if gotS.name != expS.name {
				t.Errorf("GetScripts(%s)[%d]: expected name %s got %s", tt.name, j, expS.name, gotS.name)
			}
			if gotS.result != expS.result {
				t.Errorf("GetScripts(%s)[%d]: expected result %s got %s", tt.name, j, expS.result, gotS.result)
			}
			if gotS.code != expS.code {
				t.Errorf("GetScripts(%s)[%d]: code differs. expected:\n%q\ngot:\n%q", tt.name, j, expS.code, gotS.code)
			}
			if !reflect.DeepEqual(gotS.params, expS.params) {
				t.Errorf("GetScripts(%s)[%d]: expected params:%v got :%v", tt.name, j, expS.params, gotS.params)
			}

		}
	}
}

func readFile(name string) string {
	rd, _ := os.Open("fixtures/" + name)
	byt, _ := ioutil.ReadAll(rd)
	return string(byt)
}
