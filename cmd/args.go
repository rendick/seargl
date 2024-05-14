package cmd

import (
	"flag"
	"fmt"
)

var (
	flagRun bool
	flagCli bool
)

func Args() {
	flag.BoolVar(&flagRun, "run", false, "run the web server")
	flag.BoolVar(&flagCli, "cli", false, "run CLI search engine")
	flag.Parse()

	if flagRun {
		StartServer()
		return
	}

	if flagCli {
		CliEngine()
		return
	}

	nonFlagsArgs := flag.Args()
	if len(nonFlagsArgs) == 0 {
		fmt.Printf("seargl: not enough arguments\nTry 'hpm --help' for more information.\n")
		return
	}
}
