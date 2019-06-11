package collectors

import (
	"github.com/kodebot/newsfeed/datafeed/record/collectors/field"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/rssatom"
	"github.com/kodebot/newsfeed/logger"
)

// SourceType provides available data feed source types
type SourceType string

const (
	// RssAtom represents rss/atom feed source type
	RssAtom SourceType = "rss/atom"

	// HTML is html feed source type
	HTML SourceType = "html"
)

var rssAtomCollect = rssatom.Collect

// Collect returns one or more articles from given data using given collector type
func Collect(
	data string,
	sourceType SourceType,
	fieldSettings []field.Info) []map[string]interface{} {

	switch sourceType {
	case RssAtom:
		return rssAtomCollect(data, fieldSettings)

	default:
		logger.Errorf("source type %s is not supported", sourceType)
		return nil
	}
}
