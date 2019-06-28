package rssatom

import (
	"reflect"

	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/mmcdole/gofeed"
)

type collector func(source *gofeed.Item, params map[string]interface{}) interface{}

var collectorMap = map[databot.FieldCollectorType]collector{
	databot.PluckFieldCollector: pluck}

func pluck(source *gofeed.Item, params map[string]interface{}) interface{} {
	if src, ok := params["source"]; ok {
		srcStr := src.(string)
		data := reflect.Indirect(reflect.ValueOf(*source)).FieldByName(srcStr)
		if !data.IsValid() {
			logger.Warnf("the source field %s doesn't exist in the feed item", srcStr)
			return nil
		}
		result := data.Interface()
		return result
	}
	logger.Errorf("source parameter not found for pluck collector")
	return nil
}
