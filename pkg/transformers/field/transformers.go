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
	// HTTPGet transformer scrapes html
	HTTPGet databot.FieldTransformerType = "httpGet"
	// RemoveHTMLElements removes elements matching the selector
	RemoveHTMLElements databot.FieldTransformerType = "removeHTMLElements"
	// SelectHTMLElements only keep the elements matching the selector
	SelectHTMLElements databot.FieldTransformerType = "selectHTMLElements"
	// RemoveHTMLStyles remove all html styles
	RemoveHTMLStyles databot.FieldTransformerType = "removeHTMLStyles"
	// RemoveHTMLScripts remove all html scripts
	RemoveHTMLScripts databot.FieldTransformerType = "removeHTMLScripts"
	// RemoveNonContentHTMLElements removes all empty elements including comments
	RemoveNonContentHTMLElements databot.FieldTransformerType = "removeNonContentHTMLElements"
	// RemoveHTMLElementsMatchingText removes all the elements matching given text matcher
	RemoveHTMLElementsMatchingText databot.FieldTransformerType = "removeHTMLElementsMatchingText"
	// EnclosureToURL transformer
	EnclosureToURL databot.FieldTransformerType = "enclosureToURL"
)

type transformFuncType func(val interface{}, params map[string]interface{}) interface{}

var transformersMap map[databot.FieldTransformerType]transformFuncType

func init() {
	transformersMap = map[databot.FieldTransformerType]transformFuncType{
		Regexp:                         regex,
		Value:                          value,
		Empty:                          empty,
		FormatDate:                     formatDate,
		ParseDate:                      parseDate,
		UTCNow:                         utcNow,
		Trim:                           trim,
		TrimLeft:                       trimLeft,
		TrimRight:                      trimRight,
		Replace:                        replace,
		ReplaceAll:                     replaceAll,
		HTTPGet:                        httpGet,
		RemoveHTMLElements:             removeHTMLElements,
		SelectHTMLElements:             selectHTMLElements,
		RemoveHTMLStyles:               removeHTMLStyles,
		RemoveHTMLScripts:              removeHTMLScripts,
		RemoveNonContentHTMLElements:   removeNonContentHTMLElements,
		RemoveHTMLElementsMatchingText: removeHTMLElementsMatchingText,
		EnclosureToURL:                 enclosureToURL}
}

// Transform returns transformed data
func Transform(value interface{}, transformerSpecs []*databot.FieldTransformerSpec) interface{} {
	for _, spec := range transformerSpecs {
		transformerFunc := transformersMap[spec.Type]
		if transformerFunc != nil {
			value = transformerFunc(value, spec.Params)
		} else {
			logger.Warnf("transformer %s is not found", spec.Type)
		}
	}
	return value
}
