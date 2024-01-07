package controllers

import (
	"log"

	"github.com/gocolly/colly/v2"
)

func GetImageLogo(hostname string) string {
	// Create a new collector
	c := colly.NewCollector()

	// Channel to pass the imageURL from the callback to the main function
	imageURLCh := make(chan string, 1)

	c.OnHTML("head", func(e *colly.HTMLElement) {

		imageURL := e.ChildAttr("link[rel='icon']", "href")
		if imageURL == "" {
			imageURL = e.ChildAttr("link[rel='shortcut icon']", "href")
		}

		// Send the imageURL to the channel
		imageURLCh <- imageURL
	})

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	targetURL := "https://" + hostname
	err := c.Visit(targetURL)
	if err != nil {
		log.Fatal("Error visiting the website:", err)
	}

	// Close the channel to signal that no more values will be sent
	close(imageURLCh)

	imageURL := <-imageURLCh
	return imageURL
}
