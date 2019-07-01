package recordcollector

import "github.com/kodebot/databot/pkg/databot"

const (
	// RssAtom represents the rss/atom record collector
	RssAtom databot.RecordCollectorType = "rssatom"
	// HTMLSingle represents single html record collector
	HTMLSingle = "htmlSingle"
	// HTML represents multiple html record collector
	HTML = "html"
)
