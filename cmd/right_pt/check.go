package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rightscale/right_pt/client/policy"
	policytemplate "github.com/rightscale/right_pt/sdk/policy_template"
)

func policyTemplateCheckSyntax(ctx context.Context, cli policy.Client, files []string) error {
	for _, file := range files {
		fmt.Printf("Checking %s... ", file)
		rd, err := os.Open(file)
		if err != nil {
			return err
		}
		srcBytes, err := ioutil.ReadAll(rd)
		if err != nil {
			return err
		}

		err = cli.CompilePolicyTemplate(ctx, file, string(srcBytes))

		if err == nil {
			fmt.Printf("syntax ok\n")
		} else {
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
				fmt.Printf("an unexpected error occurred: %v", err)
			}
		}
	}
	return nil
}
