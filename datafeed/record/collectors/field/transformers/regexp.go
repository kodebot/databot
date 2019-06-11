package transformers

import (
	"regexp"

	"github.com/golang/glog"
)

func regex(source interface{}, parameters map[string]interface{}) interface{} {

	fallback := parameters["fallbackValue"]
	if fallback == nil {
		fallback = source
	}

	if source == nil {
		return fallback
	}

	if sourceString, ok := source.(string); ok {

		glog.Infof("transforming %s using regexp", sourceString)

		var expr string
		if ok := parameters["expr"]; ok != nil {
			expr = ok.(string)
		} else {
			glog.Errorf("no regular expression parameter found")
			return fallback
		}

		re, err := regexp.Compile(expr)
		if err != nil {
			glog.Errorf("invalid regexp: %s error: %s. \n", expr, err.Error())
			return fallback
		}

		requiredMatchIndex := 0
		for i, val := range re.SubexpNames() {
			if val == "data" {
				requiredMatchIndex = i
			}
		}

		if requiredMatchIndex == 0 {
			glog.Errorf("invalid regular expression: %s no named group called 'data' is found. \n", expr)
			return fallback
		}
		matches := re.FindStringSubmatch(sourceString)
		if len(matches) < requiredMatchIndex+1 {
			glog.Warningf("no match found.")
			return fallback
		}

		if found := matches[requiredMatchIndex]; found != "" {
			var result interface{} = found
			return result
		}

	}
	glog.Errorf("input is not string type")
	return fallback
}
