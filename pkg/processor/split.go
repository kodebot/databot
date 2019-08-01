package processor

import (
	"reflect"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("split", split)
}

func split(input Flow, params map[string]interface{}) Flow {
	outData := make(chan interface{})

	go func() {
		for newInput := range input.Data {
			object := reflect.ValueOf(newInput)

			if object.Kind() != reflect.Slice && object.Kind() != reflect.Array {
				logger.Fatalf("input must be slice is array %+v is neither", newInput)
			}

			if object.Len() > 0 {
				for i := 0; i < object.Len(); i++ {
					outData <- object.Index(i).Interface()
				}
				input.Control <- endSplit
			}
		}
		close(outData)
	}()

	return Flow{outData, input.Control}
}
