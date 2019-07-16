package processor

import (
	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("split", split)
}

func split(input <-chan interface{}, params map[string]interface{}) <-chan interface{} {
	output := make(chan interface{})

	go func() {
		for newInput := range input {
			slice, ok := newInput.([]interface{})
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type []interface{}", slice)
			}
			for _, item := range slice {
				output <- item
			}
		}
		close(output)
	}()

	return output
}
