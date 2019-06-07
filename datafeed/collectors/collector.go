package collectors

import (
	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/collectors/feed"
	"github.com/kodebot/newsfeed/datafeed/collectors/model"
)

// Collect returns one or more articles from given data using given collector type
func Collect(
	data string,
	collector model.RecordCollectorType,
	fieldSettings []model.FieldCollectorSetting) []map[string]*interface{} {

	switch collector {
	case model.Feed:
		return feed.Collect(data, fieldSettings)

	default:
		glog.Errorf("collector type %s is not implemented", collector)
		return nil
	}

}
