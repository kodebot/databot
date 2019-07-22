package processor

import (
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("combine", combine)
}

func combine(input Flow, params map[string]interface{}) Flow {

	outputData := make(chan interface{})
	outputControl := make(chan ControlMessage)

	go func() {
		outputSlice := []interface{}{}
		for {
			select {
			case newInput, open := <-input.Data:
				if !open {
					close(outputData)
					break
				}

				item, ok := newInput.(interface{})
				if !ok {
					logger.Fatalf("unexpected input %#v. Input must be of type interface{}", item)
				}
				outputSlice = append(outputSlice, item)

			case controlData := <-input.Control:
				if controlData == endSplit {
					outputData <- outputSlice
					outputSlice = []interface{}{}
				} else {
					outputControl <- controlData // just pass it through
				}
			}

		}

	}()

	// go func() { // relay control messages
	// 	for control := range input.Control {
	// 		outputControl <- control
	// 	}
	// }()

	return Flow{
		outputData, outputControl,
	}
}
