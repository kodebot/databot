package field

import (
	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
)

func removeHTMLElements(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}

	if htmlStr, ok := val.(string); ok {
		if found := params["selectors"]; found != nil {
			if selectors, ok := found.([]string); ok {
				doc := html.NewDocument(htmlStr)
				doc.Remove(selectors...)
				return doc.HTML()
			}
			logger.Errorf("selectors is not valid")
			return val
		}
		logger.Errorf("no selectors found")
		return val
	}

	logger.Errorf("input is not a string")
	return val

}

func keepHTMLElements(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return val
	}
	// todo: replace this to get the cache initialised when the app started using the caching strategy
	url := val.(string)
	htmlDocReader := html.NewCachedDocumentReader(url, cache.NewMemCache())
	result, err := htmlDocReader.ReadAsString()
	if err != nil {
		logger.Errorf("error when feching %s error: %s", url, err.Error())
	}
	return result
}
