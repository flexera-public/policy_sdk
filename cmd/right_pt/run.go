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
	"github.com/rightscale/right_pt/sdk/incident"
	"github.com/rightscale/right_pt/sdk/policy_template"
)

// Steps:
//   1. Upload policy template. If it exists already, update it.
//   2. Apply policy template.
//   3. Print log as we go (if --tail option)
//   4. Print escalation status as we go (if --tail option)
//   5. Cleanup (stop applied policy, delete policy template)
func policyTemplateRun(ctx context.Context, cli policy.Client, file string, runOptions []string, keep bool, dryRun bool, noLog bool) error {
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

	// Handle Applied Policy
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
		DryRun:       dryRun,
		TemplateHref: pt.Href,
		Frequency:    "hourly",
		Options:      options,
	}
	ap, err := cli.CreateAppliedPolicy(ctx, p)
	if err != nil {
		return err
	}
	if !keep {
		defer cleanupRun(ctx, cli, ap)
	}

	fmt.Printf("Created AppliedPolicy %q (%s)\n", name, ap.Href)
	if !noLog {
		fmt.Printf("\nTailing policy logs\n")
	}

	lastLog := ""
	lastEtag := ""
	var lastStatus *appliedpolicy.AppliedPolicyStatus
	for {
		time.Sleep(3 * time.Second)
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
		if !noLog {
			fmt.Printf(lastLog[lastSize:])
		}

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

	if checksPassed(lastLog) {
		what := "were"
		if dryRun {
			what = "would have been"
		}
		fmt.Printf("All validations passed successfully, no incidents %s created\n", what)
		return nil
	}

	// Handle Incidents
	incList, err := cli.IndexIncidents(ctx, ap.ID, nil, "extended", "")
	if err != nil {
		return err
	}
	fmt.Printf("%d validations failed and created incidents:\n", len(incList.Items))

	for _, inc := range incList.Items {
		fmt.Printf("Incident %s\n", *inc.Href)
		fmt.Printf("Severity: %s\n", *inc.Severity)
		fmt.Printf("Category: %s\n", *inc.Category)
		fmt.Printf("Items: %d\n", *inc.ViolationDataCount)
		fmt.Printf("Summary: %s\n", *inc.Summary)
		fmt.Printf("Detail:\n%s\n\n", *inc.Detail)
	}
	if dryRun {
		return nil
	}

	finishedCount := 0
	byIncident := make(map[string]*incident.Escalations)
	for {
		for _, inc := range incList.Items {
			// Escalation states: "queued", "aborted", "pending", "running", "completed", "failed", "denied"
			escalations, err := cli.IndexEscalations(ctx, inc.ID)
			if err != nil {
				continue
			}
			if escalations.Status != "running" && escalations.Status != "queued" {
				finishedCount++
			}
			byIncident[*inc.Href] = escalations
		}
		if finishedCount == len(incList.Items) {
			break
		}
	}

	for _, inc := range incList.Items {
		fmt.Printf("Incident %s final escalation status:\n%s\n", *inc.Href, dump(byIncident[*inc.Href]))
	}

	return nil
}

func checksPassed(log string) bool {
	return strings.Contains(log, "Total items failing checks: 0")
}

func parseOptions(pt *policytemplate.PolicyTemplate, runOptions []string) ([]*appliedpolicy.ConfigurationOptionCreateType, error) {
	options := []*appliedpolicy.ConfigurationOptionCreateType{}
	var errors []string
	var seen = map[string]bool{}

	for _, o := range runOptions {
		bits := strings.SplitN(o, "=", 2)
		name := bits[0]
		p, ok := pt.Parameters[name]
		if !ok {
			paramNames := []string{}
			for _, p := range pt.Parameters {
				paramNames = append(paramNames, p.Name)
			}
			errors = append(errors, fmt.Sprintf("%s does not appear in list of parameters: %s", name, strings.Join(paramNames, ", ")))
		}
		var val interface{}
		var err error
		if len(bits) > 1 {
			val, err = coerceOption(name, bits[1], p.Type)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
		options = append(options, &appliedpolicy.ConfigurationOptionCreateType{
			Name:  bits[0],
			Value: val,
		})
		seen[bits[0]] = true
	}
	for name, _ := range pt.Parameters {
		if !seen[name] {
			errors = append(errors, fmt.Sprintf("%s is required", name))
		}
	}

	if len(errors) != 0 {
		return nil, fmt.Errorf("Parameter errors: \n  %s", strings.Join(errors, "\n  "))
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

func cleanupRun(ctx context.Context, cli policy.Client, ap *appliedpolicy.AppliedPolicy) {
	fmt.Println("Cleaning up")
	fmt.Printf("  Deleting AppliedPolicy %q (%s)\n", ap.Name, ap.Href)
	cli.DeleteAppliedPolicy(ctx, ap.ID)
}

func dump(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
