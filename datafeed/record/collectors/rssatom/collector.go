package rssatom

import (
	"reflect"

	"github.com/golang/glog"
	field "github.com/kodebot/newsfeed/datafeed/record/collectors/field"
)

// Collect returns collected fields from the given data using given field collector settings
func Collect(data string, fieldsInfo []field.Info) []map[string]interface{} {
	feeds := readFromXML(data)
	var records []map[string]interface{}
	for _, feed := range feeds {
		record := map[string]interface{}{}
		for _, fieldInfo := range fieldsInfo {

			var sourceField string
			if ok := fieldInfo.CollectorInfo.Parameters["source"]; ok != nil {
				sourceField = ok.(string)
			} else {
				sourceField = fieldInfo.Name
			}

			fieldData := reflect.Indirect(reflect.ValueOf(feed)).FieldByName(sourceField)
			var fieldRawValue interface{}
			if !fieldData.IsValid() {
				glog.Warningf("the field %s doesn't exist", sourceField)
				fieldRawValue = nil
			} else {
				fieldRawValue = fieldData.Interface()
			}

			record[fieldInfo.Name] = field.Create(fieldRawValue, fieldInfo)

		}
		records = append(records, record)
	}
	return records
}
