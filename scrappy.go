package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	count := 10

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		count -= 1
		if count > 0 {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.digitec.ch/")
}
