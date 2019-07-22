package processor

import (
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/stringutil"
)

func init() {
	register("regexp:selectEach", regexpSelectEach)
}

func regexpSelectEach(input Flow, params map[string]interface{}) Flow {
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
			result := []interface{}{}
			for _, selector := range selectors {
				matches := stringutil.RegexpMatchAll(block, selector)
				for _, match := range matches {
					result = append(result, match)
				}
			}
			outputData <- result
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
