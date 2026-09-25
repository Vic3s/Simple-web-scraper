package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func crawler(urls []string, crawlerGatherTimeSeconds int) {
	fmt.Println("Crawler begins...")
	fmt.Printf("Will crawl for %d seconds\n", crawlerGatherTimeSeconds)

	visited := []string{}

	timer := time.NewTimer(time.Duration(crawlerGatherTimeSeconds) * time.Second)
	defer timer.Stop()

loop:
	for i := 0; i < len(urls); i++ {
		select {
		case <-timer.C:
			fmt.Println("Crawler finished.")
			break loop

		default:
			getURLs(urls[i], &urls, &visited)
		}
	}
	scraper(urls)
}

func getURLs(url string, urls *[]string, visited *[]string) {
	if slices.Contains(*visited, url) {
		return
	}

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

	htmlDoc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if exists {
			if strings.HasPrefix(href, "#") {
				return
			}
			if strings.HasPrefix(href, "//") {
				href = "https:" + href
			}
			if !strings.HasPrefix(href, "http") {
				return
			}
			if strings.Contains(href, "#") {
				href = strings.Split(href, "#")[0]
			}
			if slices.Contains(*urls, href) {
				return
			}

			*urls = append(*urls, string(href))
		}
	})

	*visited = append(*visited, url)
}
