package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kingpin"

	"github.com/rightscale/right_pt/auth"
	"github.com/rightscale/right_pt/client/policy"
	"github.com/rightscale/right_pt/config"
)

var (
	// ----- Top Level -----
	app = kingpin.New("right_pt", `A command-line application for testing RightScale Policies.

right_pt contains a number of useful commands to help with development of Policies, including a syntax checker and policy runner.`)
	debug      = app.Flag("debug", "Debug mode").Short('d').Bool()
	configFile = app.Flag("config", "Config file path").Short('c').Default(config.DefaultFile(".right_pt.yml")).String()
	account    = app.Flag("account", "Config file account name to use").Short('a').String()

	// ----- Upload -----
	upCmd   = app.Command("upload", `Upload Policy Template.`)
	upFiles = upCmd.Arg("file", "Policy Template file name.").Required().ExistingFiles()

	// ----- Check Syntax -----
	chkCmd   = app.Command("check", "Check syntax for a Policy Template.")
	chkFiles = chkCmd.Arg("file", "Policy Template file name.").Required().ExistingFiles()

	// ----- Run policy template -----
	runCmd = app.Command("run", `Uploads and applies the PolicyTemplate.

Execution of the policy will then be followed. Execution log will be tailed and 
followed and incident printed out.

Options are user-supplied values for "parameters" defined in the PolicyTemplate
language. Options must be in the form of "<name>=<value>". For arrays, values
must be a comma separated list.

Example: right_pt run max_snapshots.pt regions=us-east-1,us-west-2 max_snapshot_count=100`)
	runFile        = runCmd.Arg("file", "Policy Template file name.").Required().ExistingFile()
	runOptions     = runCmd.Arg("options", "Parameter name=value pairs.").Strings()
	runNoLog       = runCmd.Flag("no-log", "Do not print policy execution log.").Short('n').Bool()
	runKeep        = runCmd.Flag("keep", "Keep applied policy running at end, for inspection in UI. Normally policy is terminated at the end.").Short('k').Bool()
	runEscalations = runCmd.Flag("run-escalations", "If set, escalations will be run. Normally dry_run is set to avoid running any escalations.").Short('r').Bool()

	// ----- RetrieveData policy template -----
	rdCmd = app.Command("retrieve_data", `Retrieve data from a Policy Template.

Execute a policy once and retrieve generated datasources, saving them to disk.

Options are user-supplied values for "parameters" defined in the PolicyTemplate
language. Options must be in the form of "<name>=<value>". For arrays, values
must be a comma separated list.

Example: right_pt retrieve_data my_policy.pt instances
		`)
	rdFile    = rdCmd.Arg("file", "Policy Template file name.").Required().ExistingFile()
	rdOptions = rdCmd.Arg("options", "Parameter name=value pairs.").Strings()
	rdNames   = rdCmd.Flag("names", "Names of resources/datasources to retrieve. By default, all datasources will be retrieved.").Short('n').Strings()
	rdOD      = rdCmd.Flag("outputDir", "Directory to store retrieved datasources.").Short('o').String()

	// ----- Run Script -----
	scriptCmd = app.Command("script", `Run the body of a script locally.

Parameters are user-supplied values for "parameters" defined in the script
declarations. Parameters must be of the form "<name>=<value>". To pass in a file
for value, such as a datasource retrieved using the retrieve_data command, use
@<file> as the value.

Example: right_pt script max_snapshots.pt --result snapshots volumes=@ec2_volumes.json instances=@ec2_instances.json max_count=50
`)
	scriptFile    = scriptCmd.Arg("file", "Policy Template file name containing script").Required().ExistingFile()
	scriptOptions = scriptCmd.Arg("parameters", "Script parameters in the same order they appear in args. To specify a file as input, specify @<filename>.").Strings()
	scriptOut     = scriptCmd.Flag("out", "Script output file. Defaults to out.json").Short('o').Default("out.json").String()
	scriptResult  = scriptCmd.Flag("result", "Name of variable holding final result to extract.").Short('r').Required().String()

	// ----- Configuration -----
	configCmd = app.Command("config", "Manage Configuration")

	configAccountCmd     = configCmd.Command("account", "Add or edit configuration for a RightScale API account")
	configAccountName    = configAccountCmd.Arg("name", "Name of RightScale API Account to add or edit").Required().String()
	configAccountDefault = configAccountCmd.Flag("default", "Set the named RightScale API Account as the default").Short('D').Bool()

	configShowCmd = configCmd.Command("show", "Show configuration")

	// ----- Update right_st -----
// 	updateCmd = app.Command("update", "Update "+app.Name+" executable")

// 	updateListCmd = updateCmd.Command("list", "List any available updates for the "+app.Name+" executable")

// 	updateApplyCmd          = updateCmd.Command("apply", "Apply the latest update for the current major version or a specified major version")
// 	updateApplyMajorVersion = updateApplyCmd.Flag("major-version", "Major version to update to").Short('m').Int()
)

func main() {
	app.Writer(os.Stdout)
	app.Version(VV)
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')
	ctx := context.Background()

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	err := config.ReadConfig(*configFile, *account)
	var client policy.Client
	if !strings.HasPrefix(command, "config") && !strings.HasPrefix(command, "update") {
		acct := config.Config.Account
		// Makes sure the config file structure is valid
		if err != nil {
			fatalError("%s: Error reading config file: %s\n", filepath.Base(os.Args[0]), err.Error())
		}

		ts, err := auth.NewOAuthAuthenticator(acct.AuthHost(), acct.RefreshToken)
		if err != nil {
			fatalError("Configuration error: %s", err.Error())
		}
		_, err = ts.TokenString()
		if err != nil {
			fatalError("Authentication error: %s", err.Error())
		}
		client = policy.NewClient(acct.Host, uint(acct.ID), ts, *debug)
	}
	switch command {
	case upCmd.FullCommand():
		files, err := walkPaths(*upFiles)
		if err != nil {
			fatalError("%s\n", err.Error())
		}
		err = policyTemplateUpload(ctx, client, files)
		if err != nil {
			fatalError("%s\n", err.Error())
		}
	case chkCmd.FullCommand():
		files, err := walkPaths(*chkFiles)
		if err != nil {
			fatalError("%s\n", err.Error())
		}
		err = policyTemplateCheckSyntax(ctx, client, files)
		if err != nil {
			fatalError("%s\n", err.Error())
		}
	case runCmd.FullCommand():
		err = policyTemplateRun(ctx, client, *runFile, *runOptions, *runKeep, !*runEscalations, *runNoLog)
		if err != nil {
			fatalError("%s\n", err.Error())
		}
	case rdCmd.FullCommand():
		err = policyTemplateRetrieveData(ctx, client, *rdFile, *rdNames, *rdOptions, *rdOD)
		if err != nil {
			fatalError("%s\n", err.Error())
		}
	case scriptCmd.FullCommand():
		err = runScript(ctx, *scriptFile, *scriptOut, *scriptResult, *scriptOptions)
		if err != nil {
			fatalError("%s\n", err.Error())
		}
	case configAccountCmd.FullCommand():
		err := config.Config.SetAccount(*configAccountName, *configAccountDefault, os.Stdin, os.Stdout)
		if err != nil {
			fatalError("%s\n", err.Error())
		}
	case configShowCmd.FullCommand():
		err := config.Config.ShowConfiguration(os.Stdout)
		if err != nil {
			fatalError("%s\n", err.Error())
		}
	}
}

// Turn a mixed array of directories and files into a linear list of files
func walkPaths(paths []string) ([]string, error) {
	files := []string{}
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return files, err
		}
		if info.IsDir() {
			err = filepath.Walk(path, func(p string, f os.FileInfo, err error) error {
				files = append(files, p)
				_, e := os.Stat(p)
				return e
			})
			if err != nil {
				return files, err
			}
		} else {
			files = append(files, path)
		}
	}
	return files, nil
}

func fatalError(format string, v ...interface{}) {
	msg := fmt.Sprintf("ERROR: "+format, v...)
	fmt.Fprintf(os.Stderr, "%s\n", msg)

	os.Exit(1)
}

func strsVal(in *[]string) []string {
	if in == nil {
		return []string{}
	}
	return *in
}
