package processor

import (
	"time"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("date:parse", parseDate)
}

func parseDate(params map[string]interface{}) Processor {
	layout := time.RFC3339
	layoutParam := params["layout"]
	if layoutParam != nil {
		var ok bool
		layout, ok = layoutParam.(string)
		if !ok {
			logger.Fatalf("layout must be string")
		}
	}

	location, _ := time.LoadLocation("UTC")
	locationParam := params["location"]

	if locationParam != nil {
		loc, ok := locationParam.(string)
		if !ok {
			logger.Fatalf("location must be string")
		}
		var err error
		location, err = time.LoadLocation(loc)
		if err != nil {
			logger.Fatalf("location is invalid %+v", loc)
		}
	}

	return func(in <-chan interface{}, out chan<- interface{}) {
		for newInput := range in {
			block, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", block)
			}

			result, err := time.ParseInLocation(layout, block, location)
			if err != nil {
				logger.Warnf("date parsing failed layout:%+v, value:%+v, location:%+v", layout, block, location)
				out <- nil

			} else {
				out <- result
			}
		}
	}
}
