package processor

import (
	"reflect"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("split", split)
}

func split(input Input, control Control, params map[string]interface{}) Output {
	output := make(chan interface{})
	go func() {
		for newInput := range input {
			object := reflect.ValueOf(newInput)

			if object.Kind() != reflect.Slice && object.Kind() != reflect.Array {
				logger.Fatalf("input must be slice is array %+v is neither", newInput)
			}

			if object.Len() > 0 {
				for i := 0; i < object.Len(); i++ {
					output <- object.Index(i).Interface()
				}
				control <- endSplit
			}
		}
		close(output)
	}()

	return output
}
