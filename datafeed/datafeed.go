package datafeed

import (
	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/record"
	rcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors"
)

// FeedInfo provides shape for config data used to create data feed
type FeedInfo struct {
	SourceName string
	// Source is usually URL - other sources like file, etc... are not supported
	Source     string
	SourceType rcollectors.SourceType
	Category   string
	Schedule   string
	Record     record.Info
}

// NewFromFeedInfo returns structured data as per the given data feed settings
func NewFromFeedInfo(filePath string) ([]map[string]interface{}, FeedInfo) {
	feedInfo := readFeedInfo(filePath)
	return NewFromURL(feedInfo.Source, feedInfo.SourceType, feedInfo.Record), feedInfo
}

// NewFromURL returns structured data as per the record info from the given url
func NewFromURL(url string, sourceType rcollectors.SourceType, recordInfo record.Info) []map[string]interface{} {
	data, err := readAsString(url)
	if err != nil {
		glog.Errorf("unable to read from url %s", url)
		return make([]map[string]interface{}, 0)
	}
	return New(data, sourceType, recordInfo)
}

// New returns structured data as per the record setting from the given data string
func New(data string, sourceType rcollectors.SourceType, recordInfo record.Info) []map[string]interface{} {
	return record.Create(data, sourceType, recordInfo)
}
