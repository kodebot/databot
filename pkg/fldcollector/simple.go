package fldcollector

import (
	"github.com/kodebot/databot/pkg/logger"
)

func value(source *interface{}, params map[string]interface{}) interface{} {
	if val, ok := params["value"]; ok {
		return val
	}
	logger.Errorf("value parameter not found for value collector")
	return nil
}
