package processor

import (
	"strings"

	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/stringutil"
)

func init() {
	register("regexp:select", regexpSelect)
}

func regexpSelect(input Input, control Control, params map[string]interface{}) Output {
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

	output := make(chan interface{})

	go func() {
		for newInput := range input {
			block, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", block)
			}
			for _, selector := range selectors {
				matches := stringutil.RegexpMatchAll(block, selector)
				block = strings.Join(matches, "")
			}
			output <- block
		}
		close(output)
	}()

	return output
}
