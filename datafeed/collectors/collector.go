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
	case model.FEED:
		return feed.Collect(data, fieldSettings)

	default:
		glog.Errorf("collector type %d is not implemented", collector)
		return nil
	}

}
