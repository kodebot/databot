package fieldtransformer

import (
	"github.com/kodebot/databot/pkg/databot"
)

// TransformFuncType is the signature of field transformer
type TransformFuncType func(val interface{}, params map[string]interface{}) interface{}

// TransformersMap contains all the general field transformers
var TransformersMap map[databot.FieldTransformerType]TransformFuncType

const (
	// Regexp value transformer
	Regexp databot.FieldTransformerType = "string:regexp:select"
	// FormatDate transformer
	FormatDate databot.FieldTransformerType = "date:format"
	// ParseDate transformer
	ParseDate databot.FieldTransformerType = "date:parse"
	// Trim transformer
	Trim databot.FieldTransformerType = "string:trim"
	// TrimLeft transformer
	TrimLeft databot.FieldTransformerType = "string:trimLeft"
	// TrimRight transfromer
	TrimRight databot.FieldTransformerType = "string:trimRight"
	// Replace transformer
	Replace databot.FieldTransformerType = "string:replace"
	// ReplaceAll transformer
	ReplaceAll databot.FieldTransformerType = "string:replaceAll"
	// HTTPGet transformer scrapes html
	HTTPGet databot.FieldTransformerType = "http:get"
	// RemoveHTMLElements removes elements matching the selector
	RemoveHTMLElements databot.FieldTransformerType = "html:elements:remove"
	// SelectHTMLElements only keep the elements matching the selector
	SelectHTMLElements databot.FieldTransformerType = "html:elements:select"
	// RemoveHTMLStyles remove all html styles
	RemoveHTMLStyles databot.FieldTransformerType = "html:styles:remove"
	// RemoveHTMLScripts remove all html scripts
	RemoveHTMLScripts databot.FieldTransformerType = "html:scripts:remove"
	// RemoveNonContentHTMLElements removes all empty elements including comments
	RemoveNonContentHTMLElements databot.FieldTransformerType = "html:elements:non-content:remove"
	// RemoveHTMLElementsMatchingText removes all the elements matching given text matcher
	RemoveHTMLElementsMatchingText databot.FieldTransformerType = "html:elements:text:match:remove"
	// HTMLMetadata gets the value from specified meta element
	HTMLMetadata databot.FieldTransformerType = "html:metadata:select"
	// EnclosureToURL transformer
	EnclosureToURL databot.FieldTransformerType = "rssatom:enclosure:toURL"
)

func init() {
	TransformersMap = map[databot.FieldTransformerType]TransformFuncType{
		Regexp:                         regex,
		FormatDate:                     formatDate,
		ParseDate:                      parseDate,
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
		HTMLMetadata:                   htmlMetadata}
}
