package processor

import (
	"reflect"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("slice:get", sliceGet)
}

func sliceGet(params map[string]interface{}) Processor {

	indexParam := params["index"]

	if indexParam == nil {
		logger.Fatalf("no index parameter found.")
	}

	index := int(indexParam.(int64))

	return func(in <-chan interface{}, out chan<- interface{}) {
		for newInput := range in {
			// input, _ := newInput.(ext.Extensions)

			// logger.Infof("***********************************************************************************")
			// logger.Infof("%+v", input[key])

			// out <- input[key]

			input := reflect.Indirect(reflect.ValueOf(newInput))
			indexFound := false
			if input.Kind() == reflect.Slice || input.Kind() == reflect.Array {

				if input.Len() > index {
					indexFound = true
					out <- input.Index(index).Interface()
				}
			}

			if !indexFound {
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
