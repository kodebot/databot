package processor

import (
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("css:select", cssSelect)
}

func cssSelect(input <-chan interface{}, params map[string]interface{}) <-chan interface{} {
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

			doc := html.NewDocument(block)
			doc.Select(selectors...)
			output <- doc.HTML()
		}
		close(output)
	}()

	return output
}
