package transformer

import (
	"strings"

	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/stringutil"
)

func init() {
	register("html:remove", regexpRemove)
}

func htmlRemove(input <-chan interface{}, params map[string]interface{}) <-chan interface{} {
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

			for _, selector := range selectors {
				matches := stringutil.RegexpMatchAll(block, selector)
				for _, match := range matches {
					block = strings.Replace(block, match, "", -1)
				}
			}
			output <- block
		}
		close(output)
	}()

	return output
}
