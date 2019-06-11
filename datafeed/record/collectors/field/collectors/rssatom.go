package collectors

import (
	"reflect"

	"github.com/mmcdole/gofeed"

	"github.com/golang/glog"
)

func rssatom(source interface{}, parameters map[string]interface{}) interface{} {
	if source == nil {
		return nil
	}

	if _, ok := source.(*gofeed.Item); !ok {
		glog.Errorf("rssatom collector only supports gofeed.Item type input")
		return nil
	}

	if fieldSource, ok := parameters["source"]; ok {
		fieldSourceStr := fieldSource.(string)
		fieldData := reflect.Indirect(reflect.ValueOf(source)).FieldByName(fieldSourceStr)
		if !fieldData.IsValid() {
			glog.Warningf("the source field %s doesn't exist in the input", fieldSourceStr)
			return nil
		}
		return fieldData.Interface()
	}
	glog.Errorf("source parameter not found for rssatom collector")
	return nil
}
