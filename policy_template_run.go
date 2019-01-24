package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rightscale/right_pt/client/policy"
	"github.com/rightscale/right_pt/sdk/applied_policy"
	"github.com/rightscale/right_pt/sdk/policy_template"
)

// Steps:
//   1. Upload policy template. If it exists already, update it.
//   2. Apply policy template.
//   3. Print log as we go (if --tail option)
//   4. Print escalation status as we go (if --tail option)
//   5. Cleanup (stop applied policy, delete policy template)
func policyTemplateRun(ctx context.Context, cli policy.Client, file string) error {
	fmt.Printf("Running %s\n", file)
	rd, err := os.Open(file)
	if err != nil {
		return err
	}
	srcBytes, err := ioutil.ReadAll(rd)
	if err != nil {
		return err
	}

	pt, err := cli.UploadPolicyTemplate(ctx, file, string(srcBytes))
	verb := "Created"
	if err != nil && errorName(err) == "conflict" {
		errTyped := err.(*policytemplate.ConflictError)
		pt, err = cli.UpdatePolicyTemplate(ctx, idFromHref(errTyped.Location), file, string(srcBytes))
		verb = "Updated"
		if err != nil {
			return err
		}
	}
	fmt.Printf("%s PolicyTemplate %q (%s) from %s\n", verb, pt.Name, pt.Href, file)

	name := pt.Name + " test"
	apList, err := cli.IndexAppliedPolicies(ctx, []string{name}, "default", "")
	if err != nil {
		return err
	}

	if len(apList.Items) > 0 {
		fmt.Printf("Found previously applied policies named %q, terminating them\n", name)
		for _, ap := range apList.Items {
			err := cli.DeleteAppliedPolicy(ctx, ap.ID)
			if err != nil {
				return err
			}
			fmt.Printf("  Terminated %s\n", ap.Href)
		}
	}
	options := []*appliedpolicy.ConfigurationOptionCreateType{}

	p := &appliedpolicy.CreatePayload{
		Name:         name,
		Description:  nil,
		DryRun:       false,
		TemplateHref: pt.Href,
		Frequency:    "hourly",
		Options:      options,
	}
	ap, err := cli.CreateAppliedPolicy(ctx, p)
	if err != nil {
		return err
	}
	fmt.Printf("Created AppliedPolicy %q (%s)\n", name, ap.Href)

	return nil
}
