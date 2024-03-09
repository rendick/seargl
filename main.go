package main

import (
	"log"
	"runtime"
	"seargl/cmd"
)

func main() {
	if runtime.GOOS == "linux" {
		cmd.ArgsList()
	} else {
		log.Fatal("You are not running Linux right now!")
	}
}
