package datafeed

import (
	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/record"
	rcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors"
)

// FeedInfo provides shape for config data used to create data feed
type FeedInfo struct {
	SourceName string
	Source     string
	SourceType rcollectors.SourceType
	Category   string
	Schedule   string
	Record     record.Info
}

// ParseFromFeedInfo returns structured data as per the given data feed settings
func ParseFromFeedInfo(filePath string) ([]map[string]interface{}, FeedInfo) {
	feedInfo := readFeedInfo(filePath)
	return ParseFromURL(feedInfo.Source, feedInfo.SourceType, feedInfo.Record), feedInfo
}

// ParseFromURL returns structured data as per the record info from the given url
func ParseFromURL(url string, sourceType rcollectors.SourceType, recordInfo record.Info) []map[string]interface{} {
	data, err := readAsString(url)
	if err != nil {
		glog.Errorf("unable to read from url %s", url)
		return make([]map[string]interface{}, 0)
	}
	return Parse(data, sourceType, recordInfo)
}

// Parse returns structured data as per the record setting from the given data string
func Parse(data string, sourceType rcollectors.SourceType, recordInfo record.Info) []map[string]interface{} {
	return record.Create(data, sourceType, recordInfo)
}
