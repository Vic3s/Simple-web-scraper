package main

import (
	"fmt"
	"net/http"
	"log"
	"slices"
	"strings"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func scraper(urls []string) {
	fmt.Println("Scraper starting...")

	for i := 0; i < len(urls); i++ {
		scrape(urls[i])
	}
}

func scrape (url string) {
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
	coreComponents(htmlDoc)
	
}

// initial scraping data methods

func coreComponents(document *goquery.Document) {}
func identifiers(document *goquery.Document) {}
func metadata(document *goquery.Document) {}
func relationships(document *goquery.Document) {}
