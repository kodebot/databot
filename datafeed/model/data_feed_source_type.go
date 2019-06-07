package model

// DataFeedSourceType provides available data feed source types
type DataFeedSourceType string

const (
	// RssAtom represents rss/atom feed source type
	RssAtom DataFeedSourceType = "rss/atom"

	// HTML is html feed source type
	HTML DataFeedSourceType = "html"
)
