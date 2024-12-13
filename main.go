package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Initialize the collector
	c := colly.NewCollector(
		// Limit the scraper to the domain of the target site
		colly.AllowedDomains("example.com"),
	)

	// Define what to do when a specific element is found
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("Title found:", e.Text)
	})

	// Log errors during scraping
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %v failed with response: %v\nError: %v", r.Request.URL, r, err)
	})

	// Log each request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	// Start scraping
	err := c.Visit("https://example.com")
	if err != nil {
		log.Fatalf("Failed to visit site: %v", err)
	}
}
