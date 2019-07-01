package rssatom

import (
	"github.com/kodebot/databot/pkg/databot"
	fieldtransformer "github.com/kodebot/databot/pkg/transformers/field"
	"github.com/mmcdole/gofeed"
)

var transformersMap = map[databot.FieldTransformerType]fieldtransformer.TransformFuncType{
	fieldtransformer.EnclosureToURL: enclosureToURL}

func enclosureToURL(val interface{}, params map[string]interface{}) interface{} {
	if enclosures, ok := val.([]*gofeed.Enclosure); ok {
		for _, enclosure := range enclosures {
			if enclosure.Type == params["enclosureType"].(string) {
				return enclosure.URL
			}
		}
	}
	return nil
}
