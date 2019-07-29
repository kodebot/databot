package processor

import (
	"github.com/kodebot/databot/pkg/rssatom"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("toRssFeed", toRssFeed)
}

func toRssFeed(input Flow, params map[string]interface{}) Flow {
	outputData := make(chan interface{})
	outputControl := make(chan ControlMessage)

	go func() {
		for newInput := range input.Data {
			block, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", block)
			}

			rssFeed := rssatom.Parse(block)

			outputData <- rssFeed
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
