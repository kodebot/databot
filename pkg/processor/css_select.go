package processor

import (
	"fmt"

	"github.com/kodebot/databot/pkg/stringutil"

	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("css:select", cssSelect)
}

func cssSelect(input Flow, params map[string]interface{}) Flow {
	fmt.Printf("%+v", params)
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

	outputData := make(chan interface{})
	outputControl := make(chan ControlMessage)

	go func() {
		for newInput := range input.Data {
			block, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", block)
			}

			doc := html.NewDocument(block)
			doc.Select(selectors...)
			outputData <- doc.HTML()
		}
		close(outputData)
	}()

	go func() { // relay control messages
		for control := range input.Control {
			outputControl <- control
		}
	}()

	return Flow{outputData, outputControl}
}
