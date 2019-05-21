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
	URL                     string
	Origin                  string
	DefaultCategory         string
	ItemThumbImageExtractor FeedDataExtractorConfig
	ItemURLExtractor        FeedDataExtractorConfig
	PublishedDateExtractors []FeedDataExtractorConfig
	DateLayouts             []string
}

// FeedConfig model
type FeedConfig struct {
	Feed []FeedConfigItem
}
