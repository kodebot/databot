package transformers

import (
	"strings"

	"github.com/golang/glog"
)

func trim(value interface{}, parameters map[string]interface{}) interface{} {

	if valueString, ok := value.(string); ok {
		return strings.TrimSpace(valueString)
	}

	glog.Errorf("trim is not allowed on non string type")
	return value
}

func trimLeft(value interface{}, parameters map[string]interface{}) interface{} {

	if valueString, ok := value.(string); ok {
		return strings.TrimPrefix(valueString, " ")
	}

	glog.Errorf("trimLeft is not allowed on non string type")
	return value
}

func trimRight(value interface{}, parameters map[string]interface{}) interface{} {

	if valueString, ok := value.(string); ok {
		return strings.TrimSuffix(valueString, " ")
	}

	glog.Errorf("trimRight is not allowed on non string type")
	return value
}
