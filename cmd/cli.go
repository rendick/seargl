package cmd

import (
	"fmt"
	"os"
	"time"
)

var SearchInputCli string

func CliEngine() {
	fmt.Print("Your input: ")
	fmt.Scan(&SearchInputCli)
	Log = fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), SearchInputCli)
	WriteLogs()
	output := Scraper(SearchInputCli)
	fmt.Println(output)
	os.Exit(0)
}
