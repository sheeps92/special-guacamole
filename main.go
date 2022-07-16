package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	var cookies []string
	var desc []string

	// Find and visit all links
	c.OnHTML("h3", func(e *colly.HTMLElement) {
		if e.Index < 6 {
			cookies = append(cookies, e.Text)
		}
	})
	c.OnHTML("P", func(e *colly.HTMLElement) {
		if e.Index > 2 && e.Index < 9 {
			desc = append(desc, e.Text)
		}
	})
	c.Visit("http://crumblcookies.com/")
	fmt.Println("Crumbl cookies of the week")
	for idx, cookie := range cookies {
		fmt.Printf("%v: %v\n", cookie, desc[idx])
	}
}
