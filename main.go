package main

import (
	"fmt"
	"runtime"
	"seargl/cmd"
)

func main() {
	os_slice := []string{"linux", "android", "openbsd", "freebsd", "netbsd", "dragonfly", "darwin", "windows"}
	os_type := false

	for _, str := range os_slice {
		if str == runtime.GOOS {
			os_type = true
			break
		}
	}

	if os_type == true {
		cmd.Args()
		return
	} else {
		fmt.Printf("error \n")
		return
	}
}
