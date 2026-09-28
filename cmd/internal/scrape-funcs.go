package internal

import (
	"strings"
	"encoding/json"
	"os"
	"log"
	"time"
	"math/rand"
	"fmt"

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

	data := IdentifiersData{}

	return data
}

func Metadata(document *goquery.Document) MetadataData {
	data := MetadataData{}

	return data
}

func writeToJSON(fileDir string, dataset any){
	// create file first
	fileDir = fileDir + "/" + fmt.Sprintf("scraped-%s-%d", time.Now().Format("2006-01-02_15-04-05"), rand.Int())

	file, err := os.Create(fileDir)
	if err != nil{
		log.Printf("Could not create file: " + err.Error())
		return
	}
	defer file.Close()

	// then write to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(dataset)
	if err != nil{
		log.Printf("Could not encode to json: " + err.Error())
		return
	}
}