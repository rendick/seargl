package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Result struct {
	Name, Url, Text string
}

func Scraper(query string) []string {
	var urls []string

	c := colly.NewCollector()

	c.OnHTML(".links_main", func(h *colly.HTMLElement) {
		url := h.ChildText("a.result__url")
		url = "https://" + url
		urls = append(urls, strings.ReplaceAll(url, `"`, "%22"))
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

	return urls
}
