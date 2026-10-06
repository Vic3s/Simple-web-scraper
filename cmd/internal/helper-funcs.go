package internal

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getMetaHelper(document *goquery.Document, attr, value string) string {
	var result string

	document.Find(fmt.Sprintf(`meta[%s="%s"]`, attr, value)).First().Each(func(_ int, s *goquery.Selection) {
		result = strings.TrimSpace(s.AttrOr("content", ""))
	})

	return result
}

func getAuthorHelper(document *goquery.Document) string {
	if author, exists := document.Find(`meta[name="author"]`).Attr("content"); exists {
		return author
	}
	if author := strings.TrimSpace(document.Find("[rel='author']").First().Text()); author != "" {
		return author
	}
	if author := strings.TrimSpace(document.Find("[itemprop='author']").First().Text()); author != "" {
		return author
	}

	return ""
}
