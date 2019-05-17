package main

import (
	"flag"
	"sync"
	"time"

	"github.com/golang/glog"

	"github.com/kodebot/newsorganiser/models"
	"github.com/kodebot/newsorganiser/services"

	"github.com/BurntSushi/toml"
)

// todo: move extractor logic to config file
// todo: use named group for extractors
// todo: refactor
// todo: robust error handling
// todo: test that the links are indeed working
// todo: run it in 10 minutes schedule
// todo: remove items after 3 days of they added
// todo: write solid tests for the extractors
// todo: sometimes link uses feed proxy url - find out when and why this happens (resolve this to original url)

var wg sync.WaitGroup

func main() {
	flag.Parse()
	var feedConfig models.FeedConfig
	_, err := toml.DecodeFile("./feed_config.toml", &feedConfig)
	if err != nil {
		glog.Fatalf("error when loading feed config: %s\n", err.Error())
	}

	// https://stackoverflow.com/a/16466581/3208697
	ticker := time.NewTicker(2 * time.Minute)
	// quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				glog.Infoln("NEW TICK...")
				for _, feed := range feedConfig.Feed {
					services.LoadFeed(feed)
				}
				glog.Infoln("TICK END")
				// case <-quit:
				// 	glog.Infoln("CLOSING...")
				// 	ticker.Stop()
				// 	return
			}
		}
	}()

	wg.Add(1)
	go forever()
	wg.Wait()
}

func forever() {

}
