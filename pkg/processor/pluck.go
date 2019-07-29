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

	outputData := make(chan interface{})
	outputControl := make(chan ControlMessage)

	go func() {
		for newInput := range input.Data {

			fieldData := reflect.Indirect(reflect.ValueOf(newInput)).FieldByName(field)
			if !fieldData.IsValid() {
				logger.Fatalf("the field %s doesn't exist in the input", field)
			}
			outputData <- fieldData.Interface()
		}
		close(outputData)
	}()

	go func() { // relay control messages
		for control := range input.Control {
			outputControl <- control
		}
	}()

	return Flow{outputData, outputControl}
}
