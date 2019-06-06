package feed

import (
	"reflect"

	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/collectors/common"
	"github.com/kodebot/newsfeed/collectors/model"
)

func Collect(url string, fieldCollectors []model.FieldCollectorSetting) []map[string]interface{} {

	feeds := readFromURL(url)
	var records []map[string]interface{}
	for _, feed := range feeds {
		record := make(map[string]interface{})
		for _, fieldCollector := range fieldCollectors {
			fieldRawValue := reflect.Indirect(reflect.ValueOf(feed)).FieldByName(fieldCollector.Field).String()

			switch fieldCollector.Type {
			case model.VALUE:
				record[fieldCollector.Field] = fieldRawValue
			case model.REGEXP:
				if paramCount := len(fieldCollector.Parameters); paramCount != 1 {
					glog.Errorf("only one regexp is supported per field collector. input has %d", paramCount)
					record[fieldCollector.Field] = ""
					break
				}
				record[fieldCollector.Field] = common.CollectUsingRegexp(fieldRawValue, fieldCollector.Parameters[0])
			default:
				glog.Errorf("collector type %d is not implemented", fieldCollector.Type)
				record[fieldCollector.Field] = ""
				break
			}
		}
	}
	return records
}
