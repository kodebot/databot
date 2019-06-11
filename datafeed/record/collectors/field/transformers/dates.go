package transformers

import (
	"strings"
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

	fallbackVal := params["fallbackValue"]

	if fallbackVal != nil {
		if str, ok := fallbackVal.(string); ok {
			if strings.HasPrefix(str, "transformer:") {
				transformer := strings.Replace(str, "transformer:", "", 1)
				if transFunc := transformersMap[transformer]; transFunc != nil {
					fallbackVal = transFunc(val, params)
				} else {
					logger.Errorf("invalid nested transformer %s", fallbackVal)
					fallbackVal = val
				}

			}
		}
	} else {
		fallbackVal = val
	}

	if val == nil {
		return fallbackVal
	}

	layoutStr := time.RFC3339
	if params != nil && params["layout"] != nil {
		layoutStr = params["layout"].(string)
	}

	valStr := val.(string)

	parseLocStr := "UTC"
	if parseLoc := params["location"]; parseLoc != nil {
		parseLocStr = parseLoc.(string)
	}

	loc, err := time.LoadLocation(parseLocStr)

	if err != nil {
		logger.Errorf("parsing location specified is not recognised %s", parseLocStr)
		return fallbackVal
	}

	result, err := time.ParseInLocation(layoutStr, valStr, loc)

	if err != nil {
		logger.Errorf("parsing date failed with layout %s", layoutStr)
		return fallbackVal
	}

	return result
}

func utcNow(val interface{}, params map[string]interface{}) interface{} {
	return time.Now().UTC()
}
