package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"

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
func policyTemplateRun(ctx context.Context, cli policy.Client, file string, runOptions []string) error {
	fmt.Printf("Running %s\n", file)
	pt, err := doUpload(ctx, cli, file)
	if err != nil {
		return err
	}

	// We need to fetch in extended view to get the list of parameters
	pt, err = cli.ShowPolicyTemplate(ctx, pt.ID, "extended")
	if err != nil {
		return err
	}

	name := pt.Name + " test"
	apList, err := cli.IndexAppliedPolicies(ctx, []string{name}, "default", "")
	if err != nil {
		return err
	}

	if len(apList.Items) > 0 {
		fmt.Printf("Found previous AppliedPolicy named %q, terminating:\n", name)
		for _, ap := range apList.Items {
			err := cli.DeleteAppliedPolicy(ctx, ap.ID)
			if err != nil {
				return err
			}
			fmt.Printf("  Terminated %s\n", ap.Href)
		}
	}
	options, err := parseOptions(pt, runOptions)
	if err != nil {
		return err
	}

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
	fmt.Printf("\nTailing policy logs\n")

	lastLog := ""
	lastEtag := ""
	var lastStatus *appliedpolicy.AppliedPolicyStatus
	for {
		time.Sleep(5 * time.Second)
		status, statusErr := cli.ShowAppliedPolicyStatus(ctx, ap.ID)
		if statusErr != nil {
			continue
		}
		lastStatus = status

		log, logErr := cli.ShowAppliedPolicyLog(ctx, ap.ID, lastEtag)
		if logErr != nil {
			// Most logErrs are of the 304 Not Modified variety and can be ignored
			continue
		}
		if *log.Etag == lastEtag {
			continue
		}
		lastSize := len(lastLog)
		lastEtag = *log.Etag
		lastLog = *log.ResponseBody
		fmt.Printf(lastLog[lastSize:])

		//fmt.Printf("STATUS: %s\n", dump(status))
		if status.LastEvaluationFinish != nil {
			break
		}
	}

	if lastStatus.EvaluationError == nil {
		fmt.Printf("\nPolicy evaluation successful\n")
	} else {
		fmt.Printf("\nPolicy evaluation failed\n")
		return errors.New(*lastStatus.EvaluationError)
	}

	return nil
}

func parseOptions(pt *policytemplate.PolicyTemplate, runOptions []string) ([]*appliedpolicy.ConfigurationOptionCreateType, error) {
	options := []*appliedpolicy.ConfigurationOptionCreateType{}
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
		options = append(options, &appliedpolicy.ConfigurationOptionCreateType{
			Name:  bits[0],
			Value: val,
		})
	}
	return options, nil
}

func coerceOption(name, val, typ string) (interface{}, error) {
	switch typ {
	case "string":
		return val, nil
	case "list":
		return strings.Split(val, ","), nil
	case "number":
		n, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, fmt.Errorf("%s is expected to be a number", name)
		}
		return n, nil
	}
	return nil, fmt.Errorf("unknown option type %q", typ)
}

func dump(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
