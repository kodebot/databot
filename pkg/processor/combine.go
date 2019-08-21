package processor

import (
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("combine", combine)
}

func combine(params map[string]interface{}) Processor {
	return func(in <-chan interface{}, out chan<- interface{}) {
		outputSlice := []interface{}{}
		for newInput := range in {
			item, ok := newInput.(interface{})
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type interface{}", item)
			}
			outputSlice = append(outputSlice, item)
		}
		out <- outputSlice
	}
}
