package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/flexera-public/policy_sdk/client/policy"
	policytemplate "github.com/flexera-public/policy_sdk/sdk/policy_template"
)

func policyTemplateUpload(ctx context.Context, cli policy.Client, files []string) error {
	success := true
	for _, file := range files {
		_, err := doUpload(ctx, cli, file, false)
		if err != nil {
			success = false
		}
	}
	if !success {
		return fmt.Errorf("Upload errors occurred")
	}
	return nil
}

func doUpload(ctx context.Context, cli policy.Client, file string, disableMetaFix bool) (*policytemplate.PolicyTemplate, error) {
	rd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	srcBytes, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	// Apply meta fix if enabled
	if !disableMetaFix {
		content := string(srcBytes)
		// Identify predictable lines that are used for Meta capability to function on child policy template
		targetLine := `header "Meta-Flexera", val($ds_is_deleted, "path")`
		if strings.Contains(content, targetLine) {
			// Find and comment out the line
			lines := strings.Split(content, "\n")
			for i, line := range lines {
				if strings.Contains(line, targetLine) {
					lines[i] = "# " + line
					// break // TODO: should we only comment out the first occurrence?
				}
			}
			content = strings.Join(lines, "\n")
			srcBytes = []byte(content)
		}
	}

	pt, err := cli.UploadPolicyTemplate(ctx, filepath.Base(file), string(srcBytes))
	verb := "Created"
	if err != nil && errorName(err) == "conflict" {
		errTyped := err.(*policytemplate.ConflictError)
		pt, err = cli.UpdatePolicyTemplate(ctx, idFromHref(errTyped.Location), filepath.Base(file), string(srcBytes))
		verb = "Updated"
	}
	if err != nil {
		printCompileError(err)
		return nil, err
	}
	fmt.Printf("%s PolicyTemplate %q (%s) from %s\n", verb, pt.Name, pt.Href, file)
	return pt, nil
}

func errorName(err error) string {
	type namer interface {
		ErrorName() string
	}
	if n, ok := err.(namer); ok {
		return n.ErrorName()
	}
	return ""
}

func idFromHref(href string) string {
	bits := strings.Split(href, "/")
	return bits[len(bits)-1]
}
