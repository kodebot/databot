package processor

import (
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("constant", constant)
}

func constant(params map[string]interface{}) Processor {
	valueParam := params["value"]

	if valueParam == nil {
		logger.Fatalf("value must be specified.")
	}

	value, ok := valueParam.(string)
	if !ok {
		logger.Fatalf("value must be valid string")
	}

	return func(in <-chan interface{}, out chan<- interface{}) {
		for range in {
			out <- value
		}
	}
}
