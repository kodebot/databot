package processor

import (
	"reflect"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("map:get", mapGet)
}

func mapGet(params map[string]interface{}) Processor {

	keyParam := params["key"]

	if keyParam == nil {
		logger.Fatalf("no key parameter found.")
	}

	key, ok := keyParam.(string)
	if !ok {
		logger.Fatalf("key must be string type")
	}

	return func(in <-chan interface{}, out chan<- interface{}) {
		for newInput := range in {
			// input, _ := newInput.(ext.Extensions)

			// logger.Infof("***********************************************************************************")
			// logger.Infof("%+v", input[key])

			// out <- input[key]

			input := reflect.Indirect(reflect.ValueOf(newInput))
			keyFound := false
			if input.Kind() == reflect.Map {
				for _, e := range input.MapKeys() {
					x := e.String()
					if x == key {
						keyFound = true
						y := input.MapIndex(e)
						out <- y.Interface()
						break
					}
				}
			}

			if !keyFound {
				out <- nil
			}

			// fieldData := reflect.Indirect(reflect.ValueOf(newInput)).FieldByName(field)
			// if !fieldData.IsValid() {
			// 	logger.Fatalf("the field %s doesn't exist in the input", field)
			// }

			// object := reflect.ValueOf(fieldData.Interface())
			// if object.Kind() == reflect.Slice || object.Kind() == reflect.Array {
			// 	slice := []interface{}{}
			// 	if object.Len() > 0 {
			// 		for i := 0; i < object.Len(); i++ {
			// 			slice = append(slice, object.Index(i).Interface())
			// 		}
			// 	}

			// 	out <- slice
			// } else {
			// 	out <- fieldData.Interface()
			// }
		}
	}
}
