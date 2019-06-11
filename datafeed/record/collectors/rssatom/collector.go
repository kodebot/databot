package rssatom

import (
	field "github.com/kodebot/newsfeed/datafeed/record/collectors/field"
)

// Collect returns collected fields from the given data using given field collector settings
func Collect(data string, fieldsInfo []field.Info) []map[string]interface{} {
	feeds := readFromXML(data)
	var records []map[string]interface{}
	for _, item := range feeds {
		record := map[string]interface{}{}
		for _, fieldInfo := range fieldsInfo {
			record[fieldInfo.Name] = field.Create(item, fieldInfo)
		}
		records = append(records, record)
	}
	return records
}
