package model

// RecordCollectorType provides available record collector types
type RecordCollectorType string

const (
	// Feed represents rss/atom FEED field group collector
	Feed RecordCollectorType = "feed"

	// HTML field group collector
	HTML RecordCollectorType = "html"
)
