package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func CoreComponents(document *goquery.Document) CoreComponentsData {
	titles := []string{}
	descriptions := []string{}
	categories := []string{}
	dates := []string{}
	prices := []string{}
	availability := []string{}

	document.Find("title").Each(func(i int, s *goquery.Selection) {
		titles = append(titles, s.Text())
	})
	document.Find(`meta[name="description"]`).Each(func(i int, s *goquery.Selection) {
		descriptions = append(descriptions, s.AttrOr("content", ""))
	})
	document.Find(".category, .categories, .tag, .tags").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			categories = append(categories, text)
		}
	})
	document.Find("time").Each(func(i int, s *goquery.Selection) {
		date := s.AttrOr("datetime", "")
		if date == "" {
			date = strings.TrimSpace(s.Text())
		}
		if date != "" {
			dates = append(dates, date)
		}
	})
	document.Find(`[itemprop="price"]`).Each(func(i int, s *goquery.Selection) {
		value := s.AttrOr("content", "")
		if value == "" {
			value = strings.TrimSpace(s.Text())
		}
		prices = append(prices, value)
	})
	document.Find(`[itemprop="availability"]`).Each(func(i int, s *goquery.Selection) {
		value := s.AttrOr("href", "")
		if value == "" {
			value = s.AttrOr("content", "")
		}
		if value == "" {
			value = strings.TrimSpace(s.Text())
		}
		availability = append(availability, value)
	})

	data := CoreComponentsData{
		titles,
		descriptions,
		categories,
		dates,
		prices,
		availability,
	}
	return data
}
func Identifiers(document *goquery.Document) IdentifiersData {
	productIds := []string{}
	articleIds := []string{}
	userIds := []string{}
	skus := []string{}
	externalIds := []string{}
	embeddedIds := []string{}

	for _, attr := range []string{
		"data-product-id",
		"data-productid",
		"product-id",
		"product_id",
		"data-id",
		"data-product",
	} {
		selector := fmt.Sprintf("[%s]", attr)

		document.Find(selector).Each(func(i int, s *goquery.Selection) {
			if value, exists := s.Attr(attr); exists && value != "" {
				productIds = append(productIds, value)
			}
		})
	}
	document.Find("article").Each(func(i int, s *goquery.Selection) {
		for _, attr := range []string{
			"data-post-id",
			"data-article-id",
			"data-content-id",
			"data-item-id",
		} {
			if value, exists := s.Attr(attr); exists && value != "" {
				articleIds = append(articleIds, value)
			}
		}
	})
	// Article IDs might be found in the URL as well - might implement functionality to get from the URL as well

	for _, attr := range []string{
		"data-user-id",
		"data-userid",
		"user-id",
		"userid",
	} {
		selector := fmt.Sprintf("[%s]", attr)

		document.Find(selector).Each(func(i int, s *goquery.Selection) {
			if value, exists := s.Attr(attr); exists && value != "" {
				userIds = append(userIds, value)
			}
		})
	}
	// User IDs can also be found in the URL - again i might add functionality to extract from there.

	document.Find("[data-sku]").Each(func(i int, s *goquery.Selection) {
		if value, exists := s.Attr("data-sku"); exists && value != "" {
			skus = append(skus, value)
		}
	})
	document.Find(".sku").Each(func(i int, s *goquery.Selection) {
		skus = append(skus, s.Text())
	})
	for _, attr := range []string{
		"stripeCustomerId",
		"stripePaymentIntentId",
		"googleUserId",
		"githubUserId",
		"externalOrderId",
		"externalProductId",
		"facebookPageId",
		"youtubeChannelId",
		"gaClientId",
		"adAccountId",
		"crmContactId",
	} {
		selector := fmt.Sprintf("[%s]", attr)

		document.Find(selector).Each(func(i int, s *goquery.Selection) {
			if value, exists := s.Attr(attr); exists && value != "" {
				externalIds = append(externalIds, value)
			}
		})
	}
	document.Find("*").Each(func(i int, s *goquery.Selection) {
		if len(s.Nodes) == 0 {
			return
		}

		for _, attr := range s.Nodes[0].Attr {
			if attr.Key == "id" || strings.HasSuffix(attr.Key, "Id") || strings.HasSuffix(attr.Key, "-id") {
				embeddedIds = append(embeddedIds, attr.Val)
			}
		}
	})
	document.Find("script").Each(func(i int, s *goquery.Selection) {
		script := s.Text()
		var data map[string]interface{}

		err := json.Unmarshal([]byte(script), &data)
		if err == nil {
			for key, value := range data {
				lower := strings.ToLower(key)
				if lower == "id" || strings.HasSuffix(lower, "Id") || strings.HasSuffix(lower, "-id") || strings.HasSuffix(lower, "-sku") {
					if str, ok := value.(string); ok && str != "" {
						embeddedIds = append(embeddedIds, str)
					}
				}
			}
		}
	})

	data := IdentifiersData{
		productIds,
		articleIds,
		userIds,
		skus,
		externalIds,
		embeddedIds,
	}

	return data
}

func SiteMetadata(document *goquery.Document) SiteMetadataData {

	title := strings.TrimSpace(document.Find("title").First().Text())

	description := GetMetaHelper(document, "name", "description")

	canonical, _ := document.Find(`link[rel="canonical"]`).Attr("href")

	robots := GetMetaHelper(document, "name", "robots")

	language, _ := document.Find("html").Attr("lang")

	author := GetAuthorHelper(document)

	data := SiteMetadataData{
		title,
		description,
		strings.TrimSpace(canonical),
		robots,
		strings.TrimSpace(language),
		author,
	}

	return data
}

func writeToJSON(fileDir string, dataset any) {
	// create file first
	fileDir = fileDir + "/" + fmt.Sprintf("scraped-%s-%d.json", time.Now().Format("2006-01-02_15-04-05"), rand.Int())

	file, err := os.Create(fileDir)
	if err != nil {
		log.Printf("Could not create file: %s", err.Error())
		return
	}
	defer file.Close()

	// then write to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(dataset)
	if err != nil {
		log.Printf("Could not encode to json: %s", err.Error())
		return
	}
	log.Printf("Scraped data saved at %s.", fileDir)
}
