package fldtransformer

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/kodebot/databot/pkg/logger"
)

func trim(val interface{}, params map[string]interface{}) interface{} {
	if val, ok := val.(string); ok {
		return strings.TrimSpace(val)
	}

	logger.Errorf("trim is not allowed on non string type")
	return nil
}

func trimLeft(val interface{}, params map[string]interface{}) interface{} {
	if val, ok := val.(string); ok {
		return strings.TrimPrefix(val, " ")
	}

	logger.Errorf("trimLeft is not allowed on non string type")
	return nil
}

func trimRight(val interface{}, params map[string]interface{}) interface{} {
	if val, ok := val.(string); ok {
		return strings.TrimSuffix(val, " ")
	}

	logger.Errorf("trimRight is not allowed on non string type")
	return nil
}

func replace(val interface{}, params map[string]interface{}) interface{} {
	if val, ok := val.(string); ok {
		olds := []string{}

		if oldParam := params["old"]; oldParam != nil {

			if reflect.TypeOf(oldParam).Kind() == reflect.Slice {
				for _, oldItem := range oldParam.([]interface{}) {
					olds = append(olds, oldItem.(string))
				}
			} else {
				if old, ok := oldParam.(string); ok {
					olds = append(olds, old)
				} else {
					logger.Errorf("old param value must be of type string or []string. %T is invalid type", oldParam)
					return nil
				}
			}

			if newParam := params["new"]; newParam != nil {
				if new, ok := newParam.(string); ok {
					n := 1
					if isReplaceAll(params) {
						n = -1
					}

					for _, old := range olds {
						val = strings.Replace(val, old, new, n)
					}
					return val
				}
				logger.Errorf("new param value must be of type string. %T is invalid type", newParam)
				return nil
			}
		}
		logger.Errorf("both new and old params must be present")
		return nil
	}

	logger.Errorf("replace is not allowed on non string type")
	return nil
}

func replaceAll(val interface{}, params map[string]interface{}) interface{} {
	if params != nil {
		params["all"] = true
	}
	return replace(val, params)
}

func regexpSelect(val interface{}, params map[string]interface{}) interface{} {
	if val, ok := val.(string); ok {
		logger.Infof("transforming %s using regexp", val)

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
		matches := re.FindStringSubmatch(val)
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

func isReplaceAll(params map[string]interface{}) bool {
	if allParam := params["all"]; allParam != nil {
		if replaceAll, ok := allParam.(bool); ok && replaceAll {
			return true
		}
	}
	return false
}
