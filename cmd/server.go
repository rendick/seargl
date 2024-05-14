package cmd

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

const (
	port = ":8080"
)

var (
	SearchInput string
	Log         string
)

type Link struct {
	Name string
}

func Html(name string, html string) {
	http.HandleFunc("/"+name, func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./www/" + html + ".html"))

		links := Scraper(SearchInput)

		tmpl.Execute(w, links)
	})
}

func StartServer() {
	fmt.Print("Your input: ")
	fmt.Scan(&SearchInput)
	Log = fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), SearchInput)

	Html("", "index")
	Html("wiki", "wiki")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./www/static"))))
	fmt.Printf("Server started.\nhttp://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
