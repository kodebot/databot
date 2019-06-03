package models

// FeedDataExtractorConfig model
type FeedDataExtractorConfig struct {
	SourceField      string
	ScrapingRequired bool
	SelectorType     string
	Selector         string
}

// FeedConfigItem model
type FeedConfigItem struct {
	Name                                 string
	URL                                  string
	Origin                               string
	DefaultCategory                      string
	ItemThumbImageExtractor              FeedDataExtractorConfig
	ItemURLExtractor                     FeedDataExtractorConfig
	ShortContentExtractor                FeedDataExtractorConfig
	PublishedDateExtractors              []FeedDataExtractorConfig
	ForceReparsePublishedDate            bool
	ReparsePublishedDateDefaultLocation  string
	DateLayouts                          []string
	UseCurrentDateTimeWhenPubDateMissing bool
}

// FeedConfig model
type FeedConfig struct {
	AllowedOrigins []string
	Feed           []FeedConfigItem
}

type ArticleSource struct {
	ID                  int
	Source              string
	SourceImageURL      string
	SourceDisplayName   string
	AvailableCategories []int
}

type ArticleCategory struct {
	ID                  int
	Category            string
	CategoryDisplayName string
	IsPublic            bool
}

type AppSettings struct {

	// ArticleCategory contains all the article categories
	ArticleCategory []ArticleCategory

	// ArticleSource contains all the article sources
	ArticleSource []ArticleSource
}
