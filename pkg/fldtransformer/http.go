package fldtransformer

import (
	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
)

type httpContext struct {
	docReaderFn func(string, cache.Manager) html.DocumentReader
}

func (ctx *httpContext) httpGet(val interface{}, params map[string]interface{}) interface{} {
	if val == nil {
		return nil
	}
	if url, ok := val.(string); ok {
		htmlDocReader := ctx.docReaderFn(url, cache.Current())
		result, err := htmlDocReader.ReadAsString()
		if err != nil {
			logger.Errorf("error when feching %s error: %s", url, err.Error())
			return nil
		}
		return result
	}
	logger.Errorf("%v is not a valid url for http:get transformer", val)
	return nil
}
