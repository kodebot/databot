package processor

import (
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("combine", combine)
}

func combine(input <-chan interface{}, params map[string]interface{}) <-chan interface{} {
	output := make(chan interface{})

	go func() {
		outputSlice := []interface{}{}
		for newInput := range input {
			item, ok := newInput.(interface{})
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type interface{}", item)
			}
			outputSlice = append(outputSlice, item)
		}
		output <- outputSlice
		close(output)
	}()

	return output
}
