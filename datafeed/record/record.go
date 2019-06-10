package record

import (
	"github.com/golang/glog"
	rcollectors "github.com/kodebot/newsfeed/datafeed/record/collectors"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/field"
	"github.com/kodebot/newsfeed/datafeed/record/collectors/rssatom"
)

// Info allows to specified record setting for parsing data
type Info struct {
	Fields []field.Info `toml:"field"`
}

// Create return one or more records created using the data provided
func Create(data string, sourceType rcollectors.SourceType, recordInfo Info) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)

	switch sourceType {
	case rcollectors.RssAtom:
		result = rssatom.Collect(data, recordInfo.Fields)
	default:
		glog.Errorf("source type %s is not implemented", sourceType)
	}

	// todo: transformers

	return result
}
