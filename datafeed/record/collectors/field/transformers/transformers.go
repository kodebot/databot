package transformers

import (
	"strings"
	"time"

	"github.com/golang/glog"
)

const (
	// Trim removes whitespaces at start and end
	Trim string = "trim"

	// FormatDate formats the date
	FormatDate string = "formatDate"
)

// TransformerInfo provides model to specify transformer settings
type TransformerInfo struct {
	Transformer string
	Parameters  map[string]interface{}
}

// TransformFormatDate returns formattted date
func TransformFormatDate(value interface{}, parameters map[string]interface{}) interface{} {
	if valueTime, ok := value.(*time.Time); ok {
		return valueTime.String()
	}

	glog.Errorf("formatDate is not allowed on non time.Time type")
	return value

}

// TransformTrim return value without leading and ending whitespaces
func TransformTrim(value interface{}, parameters map[string]interface{}) interface{} {

	if valueString, ok := value.(string); ok {
		return strings.TrimSpace(valueString)
	}

	glog.Errorf("trim is not allowed on non string type")
	return value
}
