package pipeline

import (
	"reflect"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("split", split)
}

func split(params map[string]interface{}) Operator {
	return func(in <-chan interface{}, out chan<- interface{}) {
		for newInput := range in {
			object := reflect.ValueOf(newInput)

			if object.Kind() != reflect.Slice && object.Kind() != reflect.Array {
				logger.Fatalf("input must be slice is array %+v is neither", newInput)
			}

			if object.Len() > 0 {
				for i := 0; i < object.Len(); i++ {
					out <- object.Index(i).Interface()
				}
			}
		}
	}
}
