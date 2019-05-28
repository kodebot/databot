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
	Feed []FeedConfigItem
}
