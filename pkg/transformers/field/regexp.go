package fieldtransformer

import (
	"regexp"

	"github.com/kodebot/databot/pkg/logger"
)

func regex(val interface{}, params map[string]interface{}) interface{} {

	if valStr, ok := val.(string); ok {

		logger.Infof("transforming %s using regexp", valStr)

		var expr string
		if ok := params["expr"]; ok != nil {
			expr = ok.(string)
		} else {
			logger.Errorf("no regular expression parameter found")
			return nil
		}

		re, err := regexp.Compile(expr)
		if err != nil {
			logger.Errorf("invalid regexp: %s error: %s. \n", expr, err.Error())
			return nil
		}

		requiredMatchIndex := 0
		for i, val := range re.SubexpNames() {
			if val == "data" {
				requiredMatchIndex = i
			}
		}

		if requiredMatchIndex == 0 {
			logger.Errorf("invalid regular expression: %s no named group called 'data' is found. \n", expr)
			return nil
		}
		matches := re.FindStringSubmatch(valStr)
		if len(matches) < requiredMatchIndex+1 {
			logger.Warnf("no match found.")
			return nil
		}

		if found := matches[requiredMatchIndex]; found != "" {
			var result interface{} = found
			return result
		}

	}
	logger.Errorf("input is not string type")
	return nil
}
