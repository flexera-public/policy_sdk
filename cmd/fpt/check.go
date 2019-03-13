package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rightscale/policy_sdk/client/policy"
	policytemplate "github.com/rightscale/policy_sdk/sdk/policy_template"
)

func policyTemplateCheckSyntax(ctx context.Context, cli policy.Client, files []string) error {
	success := true
	for _, file := range files {
		fmt.Printf("Checking %s\n", file)
		err := doCheck(ctx, cli, file)
		if err != nil {
			success = false
		}
	}
	if !success {
		return fmt.Errorf("compilation errors occurred")
	}
	return nil
}

func doCheck(ctx context.Context, cli policy.Client, file string) error {
	rd, err := os.Open(file)
	if err != nil {
		return err
	}
	srcBytes, err := ioutil.ReadAll(rd)
	if err != nil {
		return err
	}

	err = cli.CompilePolicyTemplate(ctx, filepath.Base(file), string(srcBytes))

	if err != nil {
		printCompileError(err)
		return err
	}
	return nil
}

func printCompileError(err error) {
	switch t := err.(type) {
	case *policytemplate.CompilationErrors:
		s := "s"
		if len(t.Errors) == 1 {
			s = ""
		}
		fmt.Printf("%d syntax error%s found:\n", len(t.Errors), s)
		for _, syntaxErr := range t.Errors {
			fmt.Printf("  origin: %s\n  problem: %s\n  summary: %s\n  resolution: %s\n\n",
				syntaxErr.Origin, syntaxErr.Problem, syntaxErr.Summary, syntaxErr.Resolution)
		}
	default:
		fmt.Printf("An unexpected error occurred: %v", err)
	}

}
