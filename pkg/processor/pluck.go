package processor

import (
	"reflect"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("pluck", pluck)
}

func pluck(input Flow, params map[string]interface{}) Flow {

	fieldParam := params["field"]

	if fieldParam == nil {
		logger.Fatalf("no field parameter found.")
	}

	field, ok := fieldParam.(string)
	if !ok {
		logger.Fatalf("field must be string type")
	}

	outData := make(chan interface{})

	go func() {
		for newInput := range input.Data {

			fieldData := reflect.Indirect(reflect.ValueOf(newInput)).FieldByName(field)
			if !fieldData.IsValid() {
				logger.Fatalf("the field %s doesn't exist in the input", field)
			}
			outData <- fieldData.Interface()
		}
		close(outData)
	}()

	return Flow{outData, input.Control}
}
