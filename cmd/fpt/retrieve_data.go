package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/flexera-public/policy_sdk/client/policy"
	appliedpolicy "github.com/flexera-public/policy_sdk/sdk/applied_policy"
	policytemplate "github.com/flexera-public/policy_sdk/sdk/policy_template"
)

func policyTemplateRetrieveData(ctx context.Context, cli policy.Client, file string, runOptions, runCredentials, names []string, outputD string) error {

	// If string header "Meta-Flexera", val($ds_flexera_api_hosts, "path") exists, add '# ' to comment it out
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
		// Instead of returning "not_found" error, return a more descriptive error message
		// "not_found" is usually due to an issue with the credential ID(s) specified
		if err.Error() == "not_found" {
			// Get list of values from credentials map
			var creds []string
			for k := range credentials {
				creds = append(creds, k)
			}
			// Update error message
			err = fmt.Errorf("At least one credential identifier not found -- please check the credential ID(s) specified. " + strings.Join(creds, ", "))
		}
		return err
	}
	for _, d := range rd {
		filename := fmt.Sprintf("%s_%s.json", d.Type, d.Name)
		// Prettyprint json
		if err != nil {
			return err
		}
		if outputD != "" {
			filename = fmt.Sprintf("%s/%s", outputD, filename)
		}

		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		e := json.NewEncoder(f)
		e.SetEscapeHTML(false)
		e.SetIndent("", "  ")

		err = e.Encode(d.Data)
		if err != nil {
			return err
		}

		f.Close()
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
