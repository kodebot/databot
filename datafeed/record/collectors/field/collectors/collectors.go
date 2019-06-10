package collectors

import (
	"regexp"

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

// CollectValue returns source value without any changes
func value(source interface{}, parameters map[string]interface{}) interface{} {
	return source
}

// CollectRegexp returns regexp collected value
func regex(source interface{}, parameters map[string]interface{}) interface{} {

	glog.Infof("collecting from %s using regexp", source)

	fieldRawValueString := source.(string)
	var expr string
	if ok := parameters["expr"]; ok != nil {
		expr = ok.(string)
	} else {
		glog.Errorf("no regular expression parameter found")
		return nil
	}

	re, err := regexp.Compile(expr)
	if err != nil {
		glog.Errorf("invalid regexp: %s error: %s. \n", expr, err.Error())
		return nil
	}

	requiredMatchIndex := 0
	for i, val := range re.SubexpNames() {
		if val == "data" {
			requiredMatchIndex = i
		}
	}

	if requiredMatchIndex == 0 {
		glog.Errorf("invalid regular expression: %s no named group called 'data' is found. \n", expr)
		return nil
	}
	matches := re.FindStringSubmatch(fieldRawValueString)
	if len(matches) < requiredMatchIndex+1 {
		glog.Warningf("no match found.")
		return nil
	}

	if found := matches[requiredMatchIndex]; found != "" {
		var result interface{} = found
		return result
	}
	return nil

}
