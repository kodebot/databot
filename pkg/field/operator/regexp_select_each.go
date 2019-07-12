package operator

import (
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/stringutil"
)

func init() {
	register("regexp:selectEach", regexpRemove)
}

func regexpSelectEach(input <-chan interface{}, params map[string]interface{}) <-chan interface{} {
	selectorsParam := params["selectors"]

	if selectorsParam == nil {
		logger.Fatalf("no selectors parameter found.")
	}

	selectors, ok := selectorsParam.([]string)
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
			result := []interface{}{}
			for _, selector := range selectors {
				matches := stringutil.RegexpMatchAll(block, selector)
				for _, match := range matches {
					result = append(result, match)
				}
			}
			output <- result
		}
		close(output)
	}()

	return output
}
