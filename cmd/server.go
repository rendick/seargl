package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"text/template"
)

const (
	port = ":8080"
)

var SearchInput string

func StartServer() {
	fmt.Print("Your input: ")
	fmt.Scan(&SearchInput)
	fmt.Printf("Server started.\nhttp://localhost%s\n", port)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./www/style"))))

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./www/index.html"))
		tsw, err := exec.Command("sh", "-c", "tsw ddg "+SearchInput).Output()
		if err != nil {
			fmt.Println(err)
		}

		data := struct {
			TswOutput string
		}{
			TswOutput: string(tsw),
		}

		tmpl.Execute(w, data)
	}

	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(port, nil))
}
