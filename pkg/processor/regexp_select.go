package processor

import (
	"strings"

	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/stringutil"
)

func init() {
	register("regexp:select", regexpSelect)
}

func regexpSelect(input Flow, params map[string]interface{}) Flow {
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
			for _, selector := range selectors {
				matches := stringutil.RegexpMatchAll(block, selector)
				block = strings.Join(matches, "")
			}
			outData <- block
		}
		close(outData)
	}()

	return Flow{outData, input.Control}
}
