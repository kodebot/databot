package processor

import (
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("combine", combine)
}

func combine(input Flow, params map[string]interface{}) Flow {
	output := make(chan interface{})
	go func() {
		outputSlice := []interface{}{}
		select {
		case newInput := <-input:
			item, ok := newInput.(interface{})
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type interface{}", item)
			}
			outputSlice = append(outputSlice, item)

		case controlData := <-control:
			if controlData == endSplit {
				output <- outputSlice
				outputSlice = []interface{}{}
			}
		}

		close(output)
	}()

	return output
}
