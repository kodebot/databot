package collectors

import (
	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/collectors/model"
	"github.com/kodebot/newsfeed/datafeed/collectors/record/rssAtom"
	dmodel "github.com/kodebot/newsfeed/datafeed/model"
)

// Collect returns one or more articles from given data using given collector type
func Collect(
	data string,
	sourceType dmodel.DataFeedSourceType,
	fieldSettings []model.FieldCollectorSetting) []map[string]*interface{} {

	switch sourceType {
	case dmodel.RssAtom:
		return rssAtom.Collect(data, fieldSettings)

	default:
		glog.Errorf("source type %s is not supported", sourceType)
		return nil
	}

}
