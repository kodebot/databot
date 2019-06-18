package transformers

import (
	"strings"

	"github.com/kodebot/newsfeed/logger"
)

func trim(value interface{}, parameters map[string]interface{}) interface{} {

	if valueString, ok := value.(string); ok {
		return strings.TrimSpace(valueString)
	}

	logger.Errorf("trim is not allowed on non string type")
	return value
}

func trimLeft(value interface{}, parameters map[string]interface{}) interface{} {

	if valueString, ok := value.(string); ok {
		return strings.TrimPrefix(valueString, " ")
	}

	logger.Errorf("trimLeft is not allowed on non string type")
	return value
}

func trimRight(value interface{}, parameters map[string]interface{}) interface{} {

	if valueString, ok := value.(string); ok {
		return strings.TrimSuffix(valueString, " ")
	}

	logger.Errorf("trimRight is not allowed on non string type")
	return value
}

func replace(value interface{}, parameters map[string]interface{}) interface{} {
	if valueString, ok := value.(string); ok {

		if old := parameters["old"]; old != nil {
			if new := parameters["new"]; new != nil {
				return strings.Replace(valueString, old.(string), new.(string), 1)
			}
		}
	}

	logger.Errorf("replace is not allowed on non string type")
	return value
}
