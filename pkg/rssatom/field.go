package rssatom

import (
	"reflect"

	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/mmcdole/gofeed"
)

// Field represents config of a field from rss/atom item
type Field struct {
	*databot.Field
	RssAtomItem *gofeed.Item
}

// Collect returns data collected using the given rss/atom Field
func (c *Field) Collect() *interface{} {
	if c.RssAtomItem == nil {
		logger.Errorf("Cannot collect field value when RssAtomItem is nil")
		return nil
	}

	if src, ok := (*c.Collector.Params)["source"]; ok {
		srcStr := (*src).(string)
		data := reflect.Indirect(reflect.ValueOf(c.RssAtomItem)).FieldByName(srcStr)
		if !data.IsValid() {
			logger.Warnf("the source field %s doesn't exist in the input", srcStr)
			return nil
		}
		result := data.Interface()
		return &result
	}
	logger.Errorf("source parameter not found for rssatom collector")
	return nil
}
