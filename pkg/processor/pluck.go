package processor

import (
	"reflect"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("pluck", pluck)
}

func pluck(params map[string]interface{}) Processor {

	fieldParam := params["field"]

	if fieldParam == nil {
		logger.Fatalf("no field parameter found.")
	}

	field, ok := fieldParam.(string)
	if !ok {
		logger.Fatalf("field must be string type")
	}

	return func(in <-chan interface{}, out chan<- interface{}) {
		for newInput := range in {
			fieldData := reflect.Indirect(reflect.ValueOf(newInput)).FieldByName(field)
			if !fieldData.IsValid() {
				logger.Fatalf("the field %s doesn't exist in the input", field)
			}

			object := reflect.ValueOf(fieldData.Interface())
			if object.Kind() == reflect.Slice || object.Kind() == reflect.Array {
				slice := []interface{}{}
				if object.Len() > 0 {
					for i := 0; i < object.Len(); i++ {
						slice = append(slice, object.Index(i).Interface())
					}
				}

				out <- slice
			} else {
				out <- fieldData.Interface()
			}
		}
	}
}
