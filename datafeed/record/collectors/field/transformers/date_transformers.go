package transformers

import (
	"time"

	"github.com/golang/glog"
)

func formatDate(value interface{}, parameters map[string]interface{}) interface{} {

	if value == nil {
		return value
	}

	if valueTime, ok := value.(*time.Time); ok {
		return valueTime.String()
	}

	if valueTime, ok := value.(time.Time); ok {
		return valueTime.String()
	}

	glog.Errorf("formatDate is not allowed on non time.Time type")
	return value

}
