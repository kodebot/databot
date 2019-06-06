package model

// RecordCollectorType provides available record collector types
type RecordCollectorType int

const (
	// FEED represents rss/atom FEED field group collector
	FEED RecordCollectorType = iota

	// HTML field group collector
	HTML
)
