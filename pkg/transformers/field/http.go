package fieldtransformer

import (
	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
)

func httpGet(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}
	url := val.(string)
	htmlDocReader := html.NewCachedDocumentReader(url, cache.Current())
	result, err := htmlDocReader.ReadAsString()
	if err != nil {
		logger.Errorf("error when feching %s error: %s", url, err.Error())
		return nil
	}
	return result
}
