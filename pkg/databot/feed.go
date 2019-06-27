package databot

// FeedSpec represents the specification for the feed
type FeedSpec struct {
	Name  string
	Desc  string
	Group string
	// Source is usually URL - other sources like file, etc... are not supported
	SourceURI  string
	SourceType FeedSourceType
	Schedule   string
	RecordSpec *RecordSpec `toml:"record"`
}

// FeedSpecReader reads config from the source and returns Feed
type FeedSpecReader interface {
	Read(source string) *FeedSpec
}

// FeedSourceType represents Feed source type
type FeedSourceType string

const (
	// RssAtomFeedSource represents the rss/atom source
	RssAtomFeedSource FeedSourceType = "rssatom"
	// HTMLSingleFeedSource represents single record html source
	HTMLSingleFeedSource = "htmlSingle"
	// HTMLFeedSource represents multiple record html source
	HTMLFeedSource = "html"
)
