package cmd

import (
	"fmt"
	"os"
)

func ArgsList() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: seargl help")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("help")
		os.Exit(0)
	case "run":
		StartServer()
	}
}
