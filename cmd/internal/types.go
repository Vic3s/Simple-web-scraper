package internal

type ScrapedData struct {
	URL            string
	CoreComponents CoreComponentsData
	Identifiers    IdentifiersData
	Metadata       MetadataData
}

type CoreComponentsData struct {
	Titles       []string `json:"titles"`
	Descriptions []string `json:"descriptions"`
	Categories   []string `json:"categories"`
	Dates        []string `json:"dates"`
	Prices       []string `json:"prices"`
	Availability []string `json:"availability"`
}

type IdentifiersData struct {
	ProductIDs  []string `json:"product_ids"`
	ArticleIDs  []string `json:"article_ids"`
	UserIDs     []string `json:"user_ids"`
	SKUs        []string `json:"skus"`
	ExternalIDs []string `json:"external_ids"`
	EmbededIDs  []string `json:"embeded_ids"`
}

type MetadataData struct {
	Author          []string `json:"author"`
	Publisher       []string `json:"publisher"`
	Location        []string `json:"location"`
	PublicationDate []string `json:"publication_date"`
	LastUpdatedDate []string `json:"last_updated_date"`
	Language        []string `json:"language"`
	PageCategory    []string `json:"page_category"`
	Tags            []string `json:"tags"`
	ImageURLs       []string `json:"image_urls"`
	CanonicalURLs   []string `json:"canonical_urls"`
}
