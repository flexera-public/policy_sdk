package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kingpin"
	//"github.com/inconshreveable/log15"
	//"github.com/mattn/go-colorable"

	"github.com/rightscale/right_pt/auth"
	"github.com/rightscale/right_pt/client/policy"
	"github.com/rightscale/right_pt/config"
)

var (
	// ----- Top Level -----
	app        = kingpin.New("right_pt", "A command-line application for managing RightScale Policies")
	debug      = app.Flag("debug", "Debug mode").Short('d').Bool()
	configFile = app.Flag("config", "Config file path").Short('c').Default(config.DefaultFile(".right_pt.yml")).String()
	account    = app.Flag("account", "Config file account name to use").Short('a').String()

	// ----- Upload -----
	upCmd   = app.Command("upload", "Upload Policy Template")
	upFiles = upCmd.Arg("file", "Policy Template file name.").Required().ExistingFiles()

	// ----- Check Syntax -----
	chkCmd   = app.Command("check", "Check syntax for Policy Template")
	chkFiles = chkCmd.Arg("file", "Policy Template file name.").Required().ExistingFiles()

	// ----- Run policy template -----
	runCmd         = app.Command("run", "Run a Policy Template once, streaming back results.")
	runFile        = runCmd.Arg("file", "Policy Template file name.").Required().ExistingFile()
	runOptions     = runCmd.Arg("options", "Parameter values.").Strings()
	runKeep        = runCmd.Flag("keep", "Keep applied policy running at end, for inspection in UI. Normally policy is terminated at the end.").Short('k').Bool()
	runEscalations = runCmd.Flag("run-escalations", "If set, escalations will be run. Normally dry_run is set to avoid running any escalations.").Short('r').Bool()

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

	//fmt.Println(command)

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

		// // Make sure the config file auth token is valid. Check now so we don't have to
		// // keep rechecking in code.
		// api, err = config.Config.Account.PolicyClient()
	}

	// // Handle logging
	// logLevel := log15.LvlInfo

	// if *debug {
	// 	log.Logger.SetHandler(
	// 		log15.LvlFilterHandler(
	// 			log15.LvlDebug,
	// 			log15.StderrHandler))
	// 	httpclient.DumpFormat = httpclient.Debug
	// 	logLevel = log15.LvlDebug
	// }
	// handler := log15.LvlFilterHandler(logLevel, log15.StreamHandler(colorable.NewColorableStdout(), log15.TerminalFormat()))
	// log15.Root().SetHandler(handler)

	// if config.Config.GetBool("update.check") && !strings.HasPrefix(command, "update") {
	// 	defer UpdateCheck(VV, os.Stderr)
	// }

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
		err = policyTemplateRun(ctx, client, *runFile, *runOptions, *runKeep, !*runEscalations)
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
