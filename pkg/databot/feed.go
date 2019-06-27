package databot

// FeedSpec represents the specification for the feed
type FeedSpec struct {
	Name       string
	Desc       string
	Group      string
	Schedule   string
	RecordSpec *RecordSpec `toml:"record"`
}

// FeedSpecReader reads config from the source and returns Feed
type FeedSpecReader interface {
	Read(source string) *FeedSpec
}
