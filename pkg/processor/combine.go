package processor

import (
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("combine", combine)
}

func combine(input Flow, params map[string]interface{}) Flow {

	outData := make(chan interface{})

	go func() {
		outputSlice := []interface{}{}
		for {
			select {
			case newInput, open := <-input.Data:
				if !open {
					close(outData)
					break
				}

				item, ok := newInput.(interface{})
				if !ok {
					logger.Fatalf("unexpected input %#v. Input must be of type interface{}", item)
				}
				outputSlice = append(outputSlice, item)

			case controlData := <-input.Control:
				if controlData == endSplit {
					outData <- outputSlice
					outputSlice = []interface{}{}
				} else {
					go func() {
						input.Control <- controlData // just pass it through
					}()
				}
			}

		}
	}()

	return Flow{
		outData, input.Control,
	}
}
