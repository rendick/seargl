package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

const (
	port = ":8080"
)

var SearchInput string
var Log string

func StartServer() {
	fmt.Print("Your input: ")
	fmt.Scan(&SearchInput)
	Log = fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), SearchInput)

	outputDefaultBrowser, err := exec.Command("xdg-settings", "get", "default-web-browser").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	trim := strings.TrimSpace(string(outputDefaultBrowser))

	replacer := strings.ReplaceAll(trim, ".desktop", "")
	replacer = strings.TrimSpace(replacer)

	fmt.Println(replacer)

	launchBrowser, err := exec.Command(replacer, "http://localhost:8080").CombinedOutput()
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(string(launchBrowser))

	fmt.Printf("Server started.\nhttp://localhost%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./www/index.html"))

		data := struct {
			TswOutput string
		}{
			TswOutput: Scraper(SearchInput),
		}

		tmpl.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
