package transformers

import (
	"strings"
	"time"

	"github.com/golang/glog"
)

const (
	// FormatDate transformer
	FormatDate string = "formatDate"

	// Trim transformer
	Trim string = "trim"
)

type transformFuncType func(value interface{}, parameters map[string]interface{}) interface{}

var transformersMap map[string]transformFuncType

func init() {
	transformersMap = map[string]transformFuncType{
		FormatDate: formatDate,
		Trim:       trim}
}

// TransformerInfo provides model to specify transformer settings
type TransformerInfo struct {
	Transformer string
	Parameters  map[string]interface{}
}

// Transform returns transformed data
func Transform(value interface{}, transformersInfo []TransformerInfo) interface{} {

	for _, info := range transformersInfo {
		transformerFunc := transformersMap[info.Transformer]
		if transformerFunc != nil {
			value = transformerFunc(value, info.Parameters)
		}
		glog.Warningf("transformer %s is not found", info.Transformer)
	}
	return value

}

// TransformFormatDate returns formattted date
func formatDate(value interface{}, parameters map[string]interface{}) interface{} {
	if valueTime, ok := value.(*time.Time); ok {
		return valueTime.String()
	}

	glog.Errorf("formatDate is not allowed on non time.Time type")
	return value

}

func trim(value interface{}, parameters map[string]interface{}) interface{} {

	if valueString, ok := value.(string); ok {
		return strings.TrimSpace(valueString)
	}

	glog.Errorf("trim is not allowed on non string type")
	return value
}
