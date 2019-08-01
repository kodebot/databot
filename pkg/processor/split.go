package processor

import (
	"reflect"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("split", split)
}

func split(input Flow, params map[string]interface{}) Flow {
	outputData := make(chan interface{})
	// outputControl := make(chan ControlMessage)
	go func() {
		for newInput := range input.Data {
			object := reflect.ValueOf(newInput)

			if object.Kind() != reflect.Slice && object.Kind() != reflect.Array {
				logger.Fatalf("input must be slice is array %+v is neither", newInput)
			}

			if object.Len() > 0 {
				for i := 0; i < object.Len(); i++ {
					outputData <- object.Index(i).Interface()
				}
				input.Control <- endSplit
			}
		}
		close(outputData)
	}()

	return Flow{outputData, input.Control}
}
