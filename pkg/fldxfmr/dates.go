package fldxfmr

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
		return result
	default:
		logger.Errorf("formatDate is not allowed on non time.Time type")
		return nil
	}
}

func parseDate(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}

	layout := time.RFC3339
	if params != nil && params["layout"] != nil {
		if l, ok := params["layout"].(string); ok {
			layout = l
		}
	}

	parseLoc := "UTC" // default parse location
	if loc := params["location"]; loc != nil {
		var ok bool
		if parseLoc, ok = loc.(string); !ok {
			logger.Errorf("parsing location should be of type string. %T is invalid type", loc)
			return nil
		}
	}

	loc, err := time.LoadLocation(parseLoc)

	if err != nil {
		logger.Errorf("parsing location specified is not recognised %s", parseLoc)
		return nil
	}

	result, err := time.ParseInLocation(layout, val.(string), loc)

	if err != nil {
		logger.Errorf("parsing date failed with layout %s. error: %s", layout, err.Error())
		return nil
	}

	return result
}
