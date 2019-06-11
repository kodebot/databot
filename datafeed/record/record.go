package record

import (
	rcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/field"
)

// Info allows to specified record setting for parsing data
type Info struct {
	Fields []field.Info `toml:"field"`
}

var rcollect = rcollectors.Collect

// Create return one or more records created using the data provided
func Create(data string, sourceType rcollectors.SourceType, recordInfo Info) []map[string]interface{} {
	result := rcollect(data, sourceType, recordInfo.Fields)
	// todo: transformers (record level - not supported at the moment)
	return result
}
