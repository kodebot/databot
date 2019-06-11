package collectors

import (
	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/field"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/rssatom"
)

// SourceType provides available data feed source types
type SourceType string

const (
	// RssAtom represents rss/atom feed source type
	RssAtom SourceType = "rss/atom"

	// HTML is html feed source type
	HTML SourceType = "html"
)

// Collect returns one or more articles from given data using given collector type
func Collect(
	data string,
	sourceType SourceType,
	fieldSettings []field.Info) []map[string]interface{} {

	switch sourceType {
	case RssAtom:
		return rssatom.Collect(data, fieldSettings)

	default:
		glog.Errorf("source type %s is not supported", sourceType)
		return nil
	}
}
