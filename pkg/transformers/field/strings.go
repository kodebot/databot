package fieldtransformer

import (
	"reflect"
	"strings"

	"github.com/kodebot/databot/pkg/logger"
)

func trim(val interface{}, params map[string]interface{}) interface{} {

	if valStr, ok := val.(string); ok {
		return strings.TrimSpace(valStr)
	}

	logger.Errorf("trim is not allowed on non string type")
	return nil
}

func trimLeft(val interface{}, params map[string]interface{}) interface{} {

	if valStr, ok := val.(string); ok {
		return strings.TrimPrefix(valStr, " ")
	}

	logger.Errorf("trimLeft is not allowed on non string type")
	return nil
}

func trimRight(val interface{}, params map[string]interface{}) interface{} {

	if valStr, ok := val.(string); ok {
		return strings.TrimSuffix(valStr, " ")
	}

	logger.Errorf("trimRight is not allowed on non string type")
	return nil
}

func replace(val interface{}, params map[string]interface{}) interface{} {
	if valStr, ok := val.(string); ok {

		olds := []string{}

		if old := params["old"]; old != nil {

			if reflect.TypeOf(old).Kind() == reflect.Slice {
				for _, oldItem := range old.([]interface{}) {
					olds = append(olds, oldItem.(string))
				}
			} else {
				olds = append(olds, old.(string))
			}

			if new := params["new"]; new != nil {
				newstr := new.(string)
				for _, oldstr := range olds {
					valStr = strings.Replace(valStr, oldstr, newstr, 1)
				}
			}
		}
		return valStr
	}

	logger.Errorf("replace is not allowed on non string type")
	return nil
}

func replaceAll(val interface{}, params map[string]interface{}) interface{} {
	if valStr, ok := val.(string); ok {
		if old := params["old"]; old != nil {
			if new := params["new"]; new != nil {
				return strings.Replace(valStr, old.(string), new.(string), -1)
			}
		}
	}

	logger.Errorf("replace is not allowed on non string type")
	return nil
}
