package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

var results []Result

type Result struct {
	Name, Url, Text string
}

func Scraper(query string) string {
	c := colly.NewCollector()

	c.OnHTML(".links_main", func(h *colly.HTMLElement) {
		Name := h.ChildText("h2.result__title")
		Url := h.ChildText("a.result__url")
		Text := h.ChildText("a.result__snippet")

		if Name == "" || Url == "" || Text == "" {
			Name, Url, Text = "nil", "nil", "nil"
		}

		results = append(results, Result{
			Name: Name,
			Url:  strings.ReplaceAll(Url, `"`, "%22"),
			Text: Text,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("%s\n\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		Log = fmt.Sprintf("Request URL: %s failed with response: %d \nError: %s", r.Request.URL, r, err)
		fmt.Println(Log)
	})

	Site := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", strings.ReplaceAll(query, "-", "+"))
	c.Visit(Site)

	var logOutput string

	for i, result := range results {
		logOutput += fmt.Sprintf("Result #%d\n", i+1)
		logOutput += fmt.Sprintf("Name: %s\nUrl: https://%s\nDescription: %s\n\n-------------------\n\n",
			result.Name, result.Url, result.Text[:110])
	}

	return logOutput
}
