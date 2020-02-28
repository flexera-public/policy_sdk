package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rightscale/policy_sdk/client/policy"
	"github.com/rightscale/policy_sdk/sdk/applied_policy"
	"github.com/rightscale/policy_sdk/sdk/policy_template"
)

func policyTemplateRetrieveData(ctx context.Context, cli policy.Client, file string, runOptions, runCredentials, names []string, outputD string) error {
	pt, err := doUpload(ctx, cli, file)
	if err != nil {
		return err
	}

	// We need to fetch in extended view to get the list of parameters
	pt, err = cli.ShowPolicyTemplate(ctx, pt.ID, "extended")
	if err != nil {
		return err
	}

	credentials := parseCredentials(runCredentials)

	apO, err := parseOptions(pt, runOptions)
	if err != nil {
		return err
	}

	options := apOptionsToptOptions(apO)

	fmt.Printf("Retrieving Data from PolicyTemplate (%s)\n", pt.Href)
	rd, err := cli.RetrieveData(ctx, pt.ID, names, options, credentials)
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
		fmt.Printf("Wrote %s\n", filename)
	}

	return nil
}

func apOptionsToptOptions(apO []*appliedpolicy.ConfigurationOptionCreateType) []*policytemplate.ConfigurationOptionCreateType {
	options := make([]*policytemplate.ConfigurationOptionCreateType, len(apO))
	for i, o := range apO {
		options[i] = &policytemplate.ConfigurationOptionCreateType{
			Name:  o.Name,
			Value: o.Value,
		}
	}
	return options
}
