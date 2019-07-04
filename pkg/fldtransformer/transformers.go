package fldtransformer

import (
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/html"
)

// TransformFuncType is the signature of field transformer
type TransformFuncType func(val interface{}, params map[string]interface{}) interface{}

// TransformersMap contains all the general field transformers
var TransformersMap map[databot.FieldTransformerType]TransformFuncType

const (
	// RegexpSelect value transformer
	RegexpSelect databot.FieldTransformerType = "string:regexp:select"
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
	// EnclosureToURL transformer - only support RSS/ATOM source
	EnclosureToURL databot.FieldTransformerType = "rssatom:enclosure:toURL"
)

func init() {

	httpCtx := httpContext{docReaderFn: html.NewCachedDocumentReader}
	htmlCtx := htmlContext{newDocFn: html.NewDocument}

	TransformersMap = map[databot.FieldTransformerType]TransformFuncType{
		RegexpSelect:                   regexpSelect,
		FormatDate:                     formatDate,
		ParseDate:                      parseDate,
		Trim:                           trim,
		TrimLeft:                       trimLeft,
		TrimRight:                      trimRight,
		Replace:                        replace,
		ReplaceAll:                     replaceAll,
		HTTPGet:                        httpCtx.httpGet,
		RemoveHTMLElements:             htmlCtx.removeElements,
		SelectHTMLElements:             htmlCtx.selectElements,
		RemoveHTMLStyles:               htmlCtx.removeStyles,
		RemoveHTMLScripts:              htmlCtx.removeScripts,
		RemoveNonContentHTMLElements:   htmlCtx.removeNonContentElements,
		RemoveHTMLElementsMatchingText: htmlCtx.removeElementsMatchingText,
		HTMLMetadata:                   htmlCtx.getMetadata}
}
