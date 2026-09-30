package main

import (
	"simple-web-scraper/cmd/internal"
	"context"
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

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(crawlerGatherTimeSeconds)*time.Second,
	)
	defer cancel()

	for i := 0; i < len(urls); i++ {
		if ctx.Err() != nil {
			break
		}
		getURLs(urls[i], &urls, &visited, ctx)
	}
	internal.Scraper(urls)
}

func getURLs(url string, urls *[]string, visited *[]string, ctx context.Context) {
	if slices.Contains(*visited, url) {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
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
