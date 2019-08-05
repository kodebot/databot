package processor

import (
	"github.com/kodebot/databot/pkg/rssatom"

	"github.com/kodebot/databot/pkg/logger"
)

func init() {
	register("toRssFeed", toRssFeed)
}

func toRssFeed(params map[string]interface{}) Processor {
	return func(in <-chan interface{}, out chan<- interface{}) {
		for newInput := range in {
			block, ok := newInput.(string)
			if !ok {
				logger.Fatalf("unexpected input %#v. Input must be of type string", block)
			}

			rssFeed := rssatom.Parse(block)

			out <- rssFeed
		}
	}
}
