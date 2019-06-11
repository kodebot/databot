package transformers

import "github.com/kodebot/newsfeed/logger"

const (
	// Regexp value transformer
	Regexp string = "regexp"
	// Value transformer returns the original value without changing it
	Value string = "value"
	// Empty transformer returns empty value
	Empty string = "empty"
	// FormatDate transformer
	FormatDate string = "formatDate"
	// ParseDate transformer
	ParseDate string = "parseDate"
	// UTCNow returns current server time in UTC
	UTCNow string = "utcNow"
	// Trim transformer
	Trim string = "trim"
	// TrimLeft transformer
	TrimLeft string = "trimLeft"
	// TrimRight transfromer
	TrimRight string = "trimRight"
)

type transformFuncType func(value interface{}, parameters map[string]interface{}) interface{}

var transformersMap map[string]transformFuncType

func init() {
	transformersMap = map[string]transformFuncType{
		Regexp:     regex,
		Value:      value,
		Empty:      empty,
		FormatDate: formatDate,
		ParseDate:  parseDate,
		UTCNow:     utcNow,
		Trim:       trim,
		TrimLeft:   trimLeft,
		TrimRight:  trimRight}
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
		logger.Warnf("transformer %s is not found", info.Transformer)
	}
	return value
}
