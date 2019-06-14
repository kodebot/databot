package scraper

import "github.com/kodebot/newsfeed/logger"

/*Scrape returns scrapped string from html
source - can be url or html string
params - can have
			sourceType - default is url. allowed values are url and html
			selectorType - default is css. allowed values are css, xpath and custom
			selector - selector for chosen selectorType
			all the params for custom selector is prefixed with custom:

*/
func Scrape(source string, params map[string]interface{}) string {

	if source == "" {
		return source
	}

	var sourceType string

	// use url when for source type when type param is missing
	// allowed types are url and html
	if params["sourceType"] == nil {
		sourceType = "url"
	} else {
		sourceType = params["sourceType"].(string)
	}

	// css, xpath and custom are supported
	var selectorType string
	if params["selectorType"] == nil { // default to css
		selectorType = "css"
	} else {
		selectorType = params["selectorType"].(string)
	}

	switch selectorType {
	case "custom":
		var selector string
		if params["custom:selector"] != nil {
			selector = params["custom:selector"].(string)
		}
		switch selector {
		case "content":
			var focusSelectors []string
			if params["custom:focusSelectors"] != nil {
				for _, focusSelector := range params["custom:focusSelectors"].([]interface{}) {
					focusSelectors = append(focusSelectors, focusSelector.(string))
				}
			}

			var blacklistedSelectors []string
			if params["custom:blacklistedSelectors"] != nil {
				for _, blacklistedSelector := range params["custom:blacklistedSelectors"].([]interface{}) {
					blacklistedSelectors = append(blacklistedSelectors, blacklistedSelector.(string))
				}
			}

			var fallbackImageSelector string
			if params["custom:fallbackImageSelector"] != nil {
				fallbackImageSelector = params["custom:fallbackImageSelector"].(string)
			}

			return extractContent(source, sourceType, focusSelectors, blacklistedSelectors, fallbackImageSelector)
		default:
			logger.Errorf("the custom selector type %s is not implemented", selector)
		}

	default:
		logger.Errorf("the selector type %s is not implemented", selectorType)
	}

	return ""
}
