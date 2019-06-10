package collectors

import (
	"github.com/golang/glog"
)

// CollectorInfo settings for collecting field
type CollectorInfo struct {
	Type       CollectorType
	Parameters map[string]interface{}
}

// CollectorType provides available field collector types
type CollectorType string

const (
	// Value field collector - the result will be same as source field
	Value CollectorType = "value"
	// Regexp field collector
	Regexp CollectorType = "regexp"

	// CSS field collector
	CSS CollectorType = "css"

	// Unknown field collector
	Unknown CollectorType = "unknown"
)

type collectorFuncType func(value interface{}, parameters map[string]interface{}) interface{}

var collectorsMap map[CollectorType]collectorFuncType

func init() {
	collectorsMap = map[CollectorType]collectorFuncType{
		Value:  value,
		Regexp: regex}
}

// Collect value from the source
func Collect(source interface{}, info CollectorInfo) interface{} {
	collector := collectorsMap[info.Type]

	if collector != nil {
		return collector(source, info.Parameters)
	}

	glog.Warningf("invalid collector type %s", info.Type)
	return nil
}
