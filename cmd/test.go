package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	p2ort = "8080"
)

func Test() {
	output_default_browser, err := exec.Command("xdg-settings", "get", "default-web-browser").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	trim := strings.TrimSpace(string(output_default_browser))
	fmt.Println("Replace .desktop with blank char")

	replacer := strings.ReplaceAll(trim, ".desktop", "")
	replacer = strings.TrimSpace(replacer)

	fmt.Println(replacer)

	launch_browser, err := exec.Command(replacer, "http://localhost:8000").CombinedOutput()
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(string(launch_browser))
}
