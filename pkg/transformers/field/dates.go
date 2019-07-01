package fieldtransformer

import (
	"time"

	"github.com/kodebot/databot/pkg/logger"
)

func formatDate(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}
	switch v := val.(type) {
	case *time.Time:
		var result interface{} = v.String()
		return result
	case time.Time:
		var result interface{} = v.String()
		return &result
	default:
		logger.Errorf("formatDate is not allowed on non time.Time type")
		return nil
	}
}

func parseDate(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
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
		return nil
	}

	result, err := time.ParseInLocation(layoutStr, valStr, loc)

	if err != nil {
		logger.Errorf("parsing date failed with layout %s. error: %s", layoutStr, err.Error())
		return nil
	}

	return result
}
