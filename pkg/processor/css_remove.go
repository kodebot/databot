package processor

import (
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/stringutil"
)

func init() {
	register("css:remove", cssRemove)
}

func cssRemove(input Flow, params map[string]interface{}) Flow {
	selectorsParam := params["selectors"]

	if selectorsParam == nil {
		logger.Fatalf("no selectors parameter found.")
	}

	selectorVals, ok := selectorsParam.([]interface{})
	if !ok {
		logger.Fatalf("selector must be specified using slice")
	}

	selectors, ok := stringutil.ToStringSlice(selectorVals)
	if !ok {
		logger.Fatalf("selector must be specified using slice of string")
	}

	outData := make(chan interface{})

	go func() {
		for newInput := range input.Data {
			block, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", block)
			}

			doc := html.NewDocument(block)
			doc.Remove(selectors...)
			outData <- doc.HTML()
		}
		close(outData)
	}()

	return Flow{outData, input.Control}
}
