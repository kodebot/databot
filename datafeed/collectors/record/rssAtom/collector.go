package rssatom

import (
	"reflect"

	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed/collectors/record/fields"
)

// Collect returns collected fields from the given data using given field collector settings
func Collect(data string, fieldCollectors []fields.CollectorInfo) []map[string]interface{} {
	feeds := readFromXML(data)
	var records []map[string]interface{}
	for _, feed := range feeds {
		record := map[string]interface{}{}
		for _, fieldCollector := range fieldCollectors {

			var sourceField string
			if ok := fieldCollector.Parameters["source"]; ok != nil {
				sourceField = ok.(string)
			} else {
				sourceField = fieldCollector.Field
			}

			fieldData := reflect.Indirect(reflect.ValueOf(feed)).FieldByName(sourceField)
			var fieldRawValue interface{}
			if !fieldData.IsValid() {
				glog.Warningf("the field %s doesn't exist", sourceField)
				fieldRawValue = nil
			} else {
				fieldRawValue = fieldData.Interface()
			}

			switch fieldCollector.Type {
			case fields.Value:
				record[fieldCollector.Field] = fields.CollectValue(fieldRawValue, fieldCollector.Parameters)
			case fields.Regexp:
				record[fieldCollector.Field] = fields.CollectRegexp(fieldRawValue, fieldCollector.Parameters)
			default:
				glog.Errorf("collector type %s is not implemented", fieldCollector.Type)
				record[fieldCollector.Field] = nil
				break
			}
		}
		records = append(records, record)
	}
	return records
}
