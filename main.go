package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/rightscale/right_pt/config"
)

var (
	app        = kingpin.New("right_pt", "A command-line application for managing RightScale Policies")
	debug      = app.Flag("debug", "Debug mode").Short('d').Bool()
	configFile = app.Flag("config", "Config file path").Short('c').Default(config.DefaultFile(".right_pt.yml")).String()
	account    = app.Flag("account", "Config file account name to use").Short('a').String()
)

func main() {
	app.Writer(os.Stdout)
	app.Version(VV)
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	fmt.Println(command)
}
