package fieldtransformer

import (
	"github.com/mmcdole/gofeed"
)

// todo: move this transformer to rssatom package
func enclosureToURL(val interface{}, params map[string]interface{}) interface{} {

	// todo: make sure value is slice to start with

	if enclosures, ok := val.([]*gofeed.Enclosure); ok {

		for _, enclosure := range enclosures {
			if enclosure.Type == params["enclosureType"].(string) {
				return enclosure.URL
			}
		}

	}
	return nil
}
