package collectors

import (
	"reflect"

	"github.com/kodebot/newsfeed/logger"
	"github.com/mmcdole/gofeed"
)

func rssatom(source interface{}, parameters map[string]interface{}) interface{} {
	if source == nil {
		return nil
	}

	if _, ok := source.(*gofeed.Item); !ok {
		logger.Errorf("rssatom collector only supports gofeed.Item type input")
		return nil
	}

	if fieldSource, ok := parameters["source"]; ok {
		fieldSourceStr := fieldSource.(string)
		fieldData := reflect.Indirect(reflect.ValueOf(source)).FieldByName(fieldSourceStr)
		if !fieldData.IsValid() {
			logger.Warnf("the source field %s doesn't exist in the input", fieldSourceStr)
			return nil
		}
		return fieldData.Interface()
	}
	logger.Errorf("source parameter not found for rssatom collector")
	return nil
}
