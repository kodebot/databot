package processor

import (
	"time"

	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/stringutil"
)

func init() {
	register("date:parse", parseDate)
}

func parseDate(params map[string]interface{}) Processor {
	layouts := []string{time.RFC3339}
	layoutParam := params["layout"]
	if layoutParam != nil {
		layoutVals, ok := layoutParam.([]interface{})
		if ok {
			layouts, ok = stringutil.ToStringSlice(layoutVals)
			if !ok {
				logger.Fatalf("layout specified using slice must be slice of string")
			}
		} else {
			val, ok := layoutParam.(string)
			if ok {
				layouts = []string{val}
			} else {
				logger.Fatalf("layout must be specified using string or slice of string")
			}
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

			parsingSuccess := false
			for _, layout := range layouts {
				result, err := time.ParseInLocation(layout, block, location)
				if err != nil {
					logger.Warnf("date parsing failed layout:%+v, value:%+v, location:%+v", layout, block, location)
					continue
				} else {
					parsingSuccess = true
					out <- result
					break
				}
			}

			if !parsingSuccess {
				logger.Warnf("unable to parse the date at all...")
				out <- nil
			}
		}
	}
}
