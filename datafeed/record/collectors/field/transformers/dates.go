package transformers

import (
	"time"

	"github.com/kodebot/newsfeed/logger"
)

func formatDate(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	if valTimePtr, ok := val.(*time.Time); ok {
		return valTimePtr.String()
	}

	if valTime, ok := val.(time.Time); ok {
		return valTime.String()
	}

	logger.Errorf("formatDate is not allowed on non time.Time type")
	return val
}

func parseDate(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	layoutStr := time.RFC3339
	if params != nil && params["layout"] != nil {
		layoutStr = params["layout"].(string)
	}

	valStr := val.(string)
	result, err := time.Parse(layoutStr, valStr)

	if err != nil {
		logger.Errorf("parsing date failed with layout %s", layoutStr)
		return val
	}

	return result
}

func utcNow(val interface{}, params map[string]interface{}) interface{} {
	return time.Now().UTC()
}
