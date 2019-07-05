package fldcollector

import (
	"github.com/kodebot/databot/pkg/logger"
)

func value(src interface{}, params map[string]interface{}) interface{} {
	if val, ok := params["value"]; ok {
		return val
	}
	logger.Errorf("value parameter not found for value collector")
	return nil
}

func source(src interface{}, params map[string]interface{}) interface{} {
	return src
}
