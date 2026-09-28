package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Scraper(urls []string) {
	fmt.Println("Scraper starting...")

	for i := 0; i < len(urls); i++ {
		scrape(urls[i])
	}
}

func scrape(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Could not reach %s, skipping.", url)
		return
	}
	defer resp.Body.Close()

	htmlDoc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Could not read HTML Body of page %s, skipping.", url)
		return
	}
	coreComponents := CoreComponents(htmlDoc)
	identifiers := Identifiers(htmlDoc)
	metadata := Metadata(htmlDoc)

	scrapedData := ScrapedData{
		url,
		coreComponents,
		identifiers,
		metadata,
	}
	writeToJSON(".", scrapedData)
}
