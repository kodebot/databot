package databot

// Feed represents source of Record s and config
type Feed struct {
	Name  string
	Desc  string
	Group string
	// Source is usually URL - other sources like file, etc... are not supported
	SourceURI  string
	SourceType SourceType
	Schedule   string
	Record     *Record
}

// SourceType represents Feed source type
type SourceType int

const (
	// RssAtom represents the rss/atom source
	RssAtom SourceType = iota + 1
	// HTMLSingle represents single record html source
	HTMLSingle
	// HTMLMultiple represents multiple record html source
	HTMLMultiple
)

// FeedConfigReader reads config from the source and returns Feed
type FeedConfigReader interface {
	Get(source string) *Feed
}
