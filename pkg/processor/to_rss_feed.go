package processor

import (
	"github.com/kodebot/databot/pkg/rssatom"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("toRssFeed", toRssFeed)
}

func toRssFeed(input Flow, params map[string]interface{}) Flow {
	outData := make(chan interface{})

	go func() {
		for newInput := range input.Data {
			block, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", block)
			}

			rssFeed := rssatom.Parse(block)

			outData <- rssFeed
		}
		close(outData)
	}()

	return Flow{outData, input.Control}
}
