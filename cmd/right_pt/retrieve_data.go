package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rightscale/right_pt/client/policy"
	"github.com/rightscale/right_pt/sdk/policy_template"
)

func policyTemplateRetrieveData(ctx context.Context, cli policy.Client, file string, names, runOptions []string, outputD string) error {
	pt, err := doUpload(ctx, cli, file)
	if err != nil {
		return err
	}

	// We need to fetch in extended view to get the list of parameters
	pt, err = cli.ShowPolicyTemplate(ctx, pt.ID, "extended")
	if err != nil {
		return err
	}

	options, err := parsePTOptions(pt, runOptions)
	if err != nil {
		return err
	}

	fmt.Printf("Retrieving Data from PolicyTemplate (%s)\n", pt.Href)
	rd, err := cli.RetrieveData(ctx, pt.ID, names, options)
	if err != nil {
		return err
	}
	for _, d := range rd {
		filename := fmt.Sprintf("%s_%s.json", d.Type, d.Name)
		// Prettyprint json
		result, err := json.MarshalIndent(d.Data, "", "\t")
		if err != nil {
			return err
		}
		if outputD != "" {
			filename = fmt.Sprintf("%s/%s", outputD, filename)
		} else {
			d, err := os.Getwd()
			if err != nil {
				return err
			}
			filename = fmt.Sprintf("%s/%s", d, filename)
		}

		err = ioutil.WriteFile(filename, result, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func parsePTOptions(pt *policytemplate.PolicyTemplate, runOptions []string) ([]*policytemplate.ConfigurationOptionCreateType, error) {
	options := []*policytemplate.ConfigurationOptionCreateType{}
	for _, o := range runOptions {
		bits := strings.SplitN(o, "=", 2)
		name := bits[0]
		p, ok := pt.Parameters[name]
		if !ok {
			paramNames := []string{}
			for _, p := range pt.Parameters {
				paramNames = append(paramNames, p.Name)
			}

			return nil, fmt.Errorf("%s does not appear in list of parameters: %s", name, strings.Join(paramNames, ", "))
		}
		var val interface{}
		var err error
		if len(bits) > 1 {
			val, err = coerceOption(name, bits[1], p.Type)
			if err != nil {
				return nil, err
			}
		}
		options = append(options, &policytemplate.ConfigurationOptionCreateType{
			Name:  bits[0],
			Value: val,
		})
	}
	return options, nil
}
