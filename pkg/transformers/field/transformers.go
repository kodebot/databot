package field

import (
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/logger"
)

const (
	// Regexp value transformer
	Regexp databot.FieldTransformerType = "regexp"
	// Value transformer returns the original value without changing it
	Value databot.FieldTransformerType = "value"
	// Empty transformer returns empty value
	Empty databot.FieldTransformerType = "empty"
	// FormatDate transformer
	FormatDate databot.FieldTransformerType = "formatDate"
	// ParseDate transformer
	ParseDate databot.FieldTransformerType = "parseDate"
	// UTCNow returns current server time in UTC
	UTCNow databot.FieldTransformerType = "utcNow"
	// Trim transformer
	Trim databot.FieldTransformerType = "trim"
	// TrimLeft transformer
	TrimLeft databot.FieldTransformerType = "trimLeft"
	// TrimRight transfromer
	TrimRight databot.FieldTransformerType = "trimRight"
	// Replace transformer
	Replace databot.FieldTransformerType = "replace"
	// ReplaceAll transformer
	ReplaceAll databot.FieldTransformerType = "replaceAll"
	// Scrape transformer scrapes html
	Scrape databot.FieldTransformerType = "scrape"
	// EnclosureToURL transformer
	EnclosureToURL databot.FieldTransformerType = "enclosureToURL"
)

type transformFuncType func(val interface{}, params map[string]interface{}) interface{}

var transformersMap map[databot.FieldTransformerType]transformFuncType

func init() {
	transformersMap = map[databot.FieldTransformerType]transformFuncType{
		Regexp:         regex,
		Value:          value,
		Empty:          empty,
		FormatDate:     formatDate,
		ParseDate:      parseDate,
		UTCNow:         utcNow,
		Trim:           trim,
		TrimLeft:       trimLeft,
		TrimRight:      trimRight,
		Replace:        replace,
		ReplaceAll:     replaceAll,
		Scrape:         scrape,
		EnclosureToURL: enclosureToURL}
}

// Transform returns transformed data
func Transform(value interface{}, transformerSpecs []*databot.FieldTransformerSpec) interface{} {
	for _, spec := range transformerSpecs {
		transformerFunc := transformersMap[spec.Type]
		if transformerFunc != nil {
			value = transformerFunc(value, spec.Params)
		}
		logger.Warnf("transformer %s is not found", spec.Type)
	}
	return value
}
