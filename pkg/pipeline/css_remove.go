package pipeline

import (
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/stringutil"
)

func init() {
	register("css:remove", cssRemove)
}

func cssRemove(params map[string]interface{}) Operator {
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

	return func(in <-chan interface{}, out chan<- interface{}) {
		for newInput := range in {
			block, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", block)
			}

			doc := html.NewDocument(block)
			doc.Remove(selectors...)
			out <- doc.HTML()
		}
	}
}
