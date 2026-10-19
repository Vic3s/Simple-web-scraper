package unittests

import (
	"context"
	"simple-web-scraper/cmd/internal"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetURLs(t *testing.T) {
	type testCase struct {
		// Inputs
		url     string
		urls    *[]string
		visited *[]string
		ctx     context.Context

		// Outputs
		expectedURLs    []string
		expectedVisited []string
	}

	t.Run("URL collection from HTML docs", func(t *testing.T) {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		tests := []testCase{
			{
				url: "https://example.com/home",
				urls: &[]string{
					"https://nowhere.invalid/imaginary-page-582104",
					"https://example.com/nonexistent-page-73926184",
				},
				visited: &[]string{
					"https://dreamtranslator.com/about",
					"https://invisiblepet.com/pets/i9ga271",
				},
				ctx: timeoutCtx,
				expectedURLs: []string{
					"https://nowhere.invalid/imaginary-page-582104",
					"https://example.com/nonexistent-page-73926184",
				},
				expectedVisited: []string{
					"https://dreamtranslator.com/about",
					"https://invisiblepet.com/pets/i9ga271",
					"https://example.com/home",
				},
			},
			{
				url: "https://example.com/home",
				urls: &[]string{
					"https://nowhere.invalid/imaginary-page-582104",
					"https://example.com/nonexistent-page-73926184",
				},
				visited: &[]string{
					"https://dreamtranslator.com/about",
					"https://invisiblepet.com/pets/i9ga271",
					"https://example.com/home",
				},
				ctx: timeoutCtx,
				expectedURLs: []string{
					"https://nowhere.invalid/imaginary-page-582104",
					"https://example.com/nonexistent-page-73926184",
				},
				expectedVisited: []string{
					"https://dreamtranslator.com/about",
					"https://invisiblepet.com/pets/i9ga271",
					"https://example.com/home",
				},
			},
		}
		for _, test := range tests {
			internal.GetURLs(test.url, test.urls, test.visited, test.ctx)
			assert.Equal(t, test.expectedURLs, *test.urls)
			assert.Equal(t, test.expectedVisited, *test.visited)
		}
	})
	t.Run("Cancels context imediatley and returns", func(t *testing.T) {
		preCanceledCtx, cancel := context.WithCancel(context.Background())
		cancel()

		tests := []testCase{
			{
				url: "https://example.com/home",
				urls: &[]string{
					"https://nowhere.invalid/imaginary-page-582104",
					"https://example.com/nonexistent-page-73926184",
				},
				visited: &[]string{
					"https://dreamtranslator.com/about",
					"https://invisiblepet.com/pets/i9ga271",
				},
				ctx: preCanceledCtx,
				expectedURLs: []string{
					"https://nowhere.invalid/imaginary-page-582104",
					"https://example.com/nonexistent-page-73926184",
				},
				expectedVisited: []string{
					"https://dreamtranslator.com/about",
					"https://invisiblepet.com/pets/i9ga271",
				},
			},
		}

		testCancelCtx := tests[0]

		internal.GetURLs(
			testCancelCtx.url,
			testCancelCtx.urls,
			testCancelCtx.visited,
			testCancelCtx.ctx,
		)
		assert.Equal(t, testCancelCtx.expectedURLs, *testCancelCtx.urls)
	})
}

func TestCrawler(t *testing.T) {
	type testCase struct {
		// Inputs
		urls                     []string
		crawlerGatherTimeSeconds int

		// Outputs the same urls slice
	}

	t.Run("Returns the same length list or bigger(has crawled new URLs)", func(t *testing.T) {

		tests := []testCase{
			{
				urls: []string{
					"https://example.com/home",
				},
				crawlerGatherTimeSeconds: 10,
			},
			{
				urls: []string{
					"https://example.com/home",
				},
				crawlerGatherTimeSeconds: 10,
			},
			{
				urls: []string{
					"https://example.com/home",
				},
				crawlerGatherTimeSeconds: 0,
			},
		}

		for _, test := range tests {
			crawler := internal.Crawler(test.urls, test.crawlerGatherTimeSeconds)
			assert.Subset(t, crawler, test.urls)
		}
	})
}
