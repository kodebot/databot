package processor

import "time"

func init() {
	register("utcnow", utcnow)
}

func utcnow(params map[string]interface{}) Processor {
	return func(in <-chan interface{}, out chan<- interface{}) {
		for range in {
			out <- time.Now().UTC()
		}
	}
}
