package field

import (
	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
)

func httpGet(val interface{}, params map[string]interface{}) interface{} {
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
