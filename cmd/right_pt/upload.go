package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rightscale/right_pt/client/policy"
	policytemplate "github.com/rightscale/right_pt/sdk/policy_template"
)

func policyTemplateUpload(ctx context.Context, cli policy.Client, files []string) error {
	success := true
	for _, file := range files {
		rd, err := os.Open(file)
		if err != nil {
			return err
		}
		srcBytes, err := ioutil.ReadAll(rd)
		if err != nil {
			return err
		}

		verb := "Created"
		pt, err := cli.UploadPolicyTemplate(ctx, file, string(srcBytes))
		if err != nil && errorName(err) == "conflict" {
			verb = "Updated"
			errTyped := err.(*policytemplate.ConflictError)
			pt, err = cli.UpdatePolicyTemplate(ctx, idFromHref(errTyped.Location), file, string(srcBytes))
		}

		if err == nil {
			fmt.Printf("%s %q (%s) from %s\n", verb, pt.Name, pt.Href, file)
		} else {
			success = false
			fmt.Printf("Failed to upload %s: %v", file, err)
		}
	}
	if !success {
		fmt.Errorf("compilation errors occurred")
	}
	return nil
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
