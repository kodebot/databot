package feed

import (
	"reflect"

	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/collectors/common"
	"github.com/kodebot/newsfeed/datafeed/collectors/model"
)

// Collect returns collected fields from the given data using given field collector settings
func Collect(data string, fieldCollectors []model.FieldCollectorSetting) []map[string]*interface{} {
	feeds := readFromXML(data)
	var records []map[string]*interface{}
	for _, feed := range feeds {
		record := map[string]*interface{}{}
		for _, fieldCollector := range fieldCollectors {

			var sourceField string
			if ok := fieldCollector.Parameters["Source"]; ok != nil {
				sourceField = ok.(string)
			} else {
				sourceField = fieldCollector.Field
			}

			fieldRawValue := reflect.Indirect(reflect.ValueOf(feed)).FieldByName(sourceField).Interface()

			switch fieldCollector.Type {
			case model.VALUE:
				record[fieldCollector.Field] = &fieldRawValue
			case model.REGEXP:
				fieldRawValueString := fieldRawValue.(string) // todo: check if this works for all the types
				var expr string
				if ok := fieldCollector.Parameters["Expr"]; ok != nil {
					expr = ok.(string)
				} else {
					glog.Errorf("no regular expression parameter found")
					record[fieldCollector.Field] = nil
					break
				}
				var result interface{} = common.CollectUsingRegexp(fieldRawValueString, expr)
				record[fieldCollector.Field] = &result
			default:
				glog.Errorf("collector type %d is not implemented", fieldCollector.Type)
				record[fieldCollector.Field] = nil
				break
			}
		}
		records = append(records, record)
	}
	return records
}
