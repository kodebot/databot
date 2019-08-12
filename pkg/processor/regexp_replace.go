package processor

import (
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/stringutil"
)

func init() {
	register("regexp:replace", regexpReplace)
}

func regexpReplace(params map[string]interface{}) Processor {
	oldParam := params["old"]

	if oldParam == nil {
		logger.Fatalf("no old parameter found.")
	}

	old, ok := oldParam.(string)
	if !ok {
		logger.Fatalf("old must be string")
	}

	newParam := params["new"]

	if newParam == nil {
		logger.Fatalf("no new parameter found.")
	}

	new, ok := newParam.(string)
	if !ok {
		logger.Fatalf("new must be string")
	}

	return func(in <-chan interface{}, out chan<- interface{}) {
		for newInput := range in {
			block, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", block)
			}

			result := stringutil.RegexpReplaceAll(block, old, new)
			out <- result
		}
	}
}
