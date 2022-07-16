package main

import (
	"testing"

	"github.com/gocolly/colly"
)

func TestMain(t *testing.T) {
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
	if cookies == nil || desc == nil {
		t.Fatalf("Cookies or description is nil\nCookies: %v\nDescription: %v\n", cookies, desc)
	}
}
