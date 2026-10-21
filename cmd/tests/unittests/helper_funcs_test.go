package unittests

import (
	"log"
	"simple-web-scraper/cmd/internal"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestGetMetaHelper(t *testing.T) {
	type testCase struct {
		// Inputs
		document *goquery.Document
		attr     string
		value    string

		// Outputs
		result string
	}

	t.Run("Collect metatags content value", func(t *testing.T) {
		doc1, err := goquery.NewDocumentFromReader(strings.NewReader(`
		<html> 
			<head>
				<meta name="description" content="Description of a page here">
			</head>
			<body>
				<h1>Heading1</h1>
			</bidy>
		</html>
		`))
		if err != nil {
			log.Printf("An error has occured while testing: %s", err)
		}
		doc2, err := goquery.NewDocumentFromReader(strings.NewReader(`
		<html> 
			<head>
				<meta name="robots" content="index, follow">
			</head>
			<body>
				<h1>Heading1</h1>
			</bidy>
		</html>
		`))
		if err != nil {
			log.Printf("An error has occured while testing: %s", err)
		}

		tests := []testCase{
			{
				document: doc1,
				attr:     "name",
				value:    "description",
				result:   "Description of a page here",
			},
			{
				document: doc2,
				attr:     "name",
				value:    "robots",
				result:   "index, follow",
			},
		}

		for _, test := range tests {
			metaContent := internal.GetMetaHelper(test.document, test.attr, test.value)
			assert.Equal(t, metaContent, test.result)
		}
	})

	t.Run("Returns nothing when no metatag was found", func(t *testing.T) {
		doc1, err := goquery.NewDocumentFromReader(strings.NewReader(`
		<html> 
			<head>
			</head>
			<body>
				<h1>Heading1</h1>
			</bidy>
		</html>
		`))
		if err != nil {
			log.Printf("An error has occured while testing: %s", err)
		}

		tests := []testCase{
			{
				document: doc1,
				attr:     "name",
				value:    "description",
			},
		}
		test := tests[0]

		metaContent := internal.GetMetaHelper(test.document, test.attr, test.value)
		assert.Equal(t, metaContent, "")
	})
}

func TestGetAuthorHelper(t *testing.T) {
	type testCase struct {
		// Inputs
		document *goquery.Document

		// Outputs
		author string
	}

	t.Run("Returns the author meta tag content if it exists", func(t *testing.T) {
		doc1, err := goquery.NewDocumentFromReader(strings.NewReader(`
		<html> 
			<head>
				<meta name="author" content="Author name here">
			</head>
			<body>
				<h1>Heading1</h1>
			</bidy>
		</html>
		`))
		if err != nil {
			log.Printf("An error has occured while testing: %s", err)
		}

		doc2, err := goquery.NewDocumentFromReader(strings.NewReader(`
		<html> 
			<head>
			</head>
			<body>
				<h1>Heading1</h1>
				<a href="https://example.com/author" rel="author">Author name here</a>
			</bidy>
		</html>
		`))
		if err != nil {
			log.Printf("An error has occured while testing: %s", err)
		}
		doc3, err := goquery.NewDocumentFromReader(strings.NewReader(`
		<html>
			<head>
			</head>
			<body>
				<h1>Heading1</h1>
				<span itemprop="author">Author name here</span>
			</bidy>
		</html>
		`))
		if err != nil {
			log.Printf("An error has occured while testing: %s", err)
		}

		tests := []testCase{
			{
				document: doc1,
				author:   "Author name here",
			},
			{
				document: doc2,
				author:   "Author name here",
			},
			{
				document: doc3,
				author:   "Author name here",
			},
		}

		for _, test := range tests {
			author := internal.GetAuthorHelper(test.document)
			assert.Equal(t, author, test.author)
		}
	})

	t.Run("Returns the author meta tag content if it exists", func(t *testing.T) {
		doc1, err := goquery.NewDocumentFromReader(strings.NewReader(`
		<html> 
			<head>
			</head>
			<body>
				<h1>Heading1</h1>
			</bidy>
		</html>
		`))
		if err != nil {
			log.Printf("An error has occured while testing: %s", err)
		}

		tests := []testCase{
			{
				document: doc1,
				author:   "",
			},
		}
		test := tests[0]

		author := internal.GetAuthorHelper(test.document)
		assert.Equal(t, author, test.author)
	})
}