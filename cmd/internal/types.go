package internal

type ScrapedData struct {
	URL            string             `json:"url"`
	CoreComponents CoreComponentsData `json:"core_components"`
	Identifiers    IdentifiersData    `json:"identifiers"`
	SiteMetadata   SiteMetadataData   `json:"metadata"`
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
	EmbeddedIDs []string `json:"embedded_ids"`
}

type SiteMetadataData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Canonical   string `json:"canonical"`
	Robots      string `json:"robots"`
	Language    string `json:"language"`
	Author      string `json:"author"`
}
