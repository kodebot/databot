package services

import (
	"flag"
	"sync"
	"time"

	"github.com/golang/glog"

	"github.com/kodebot/newsfeed/models"
	"github.com/kodebot/newsfeed/services"

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
// todo: record last loaded time for each feed
// todo: move scheduling to config

var wg sync.WaitGroup

func test() {
	flag.Parse()
	defer crashHandler()
	var feedConfig models.FeedConfig
	_, err := toml.DecodeFile("./feed_config.toml", &feedConfig)
	if err != nil {
		glog.Fatalf("error when loading feed config: %s\n", err.Error())
	}

	// https://stackoverflow.com/a/16466581/3208697
	loadArticlesTicker := time.NewTicker(30 * time.Minute)
	cleanArticlesTicker := time.NewTicker(6 * time.Hour)

	loadArticles := func() {
		for _, feed := range feedConfig.Feed {
			services.LoadArticlesFromFeed(feed)
		}
	}

	go scheduler("Feed Loader Task", loadArticlesTicker, loadArticles)
	go scheduler("Clean News Item", cleanArticlesTicker, services.PruneArticles)

	wg.Add(1)
	go forever()
	wg.Wait()
}

func forever() {

}

func scheduler(name string, ticker *time.Ticker, task func()) {
	for {
		<-ticker.C
		glog.Infof("new tick received on %s...\n", name)
		task()
		glog.Infof("tick completed on %s...\n", name)
	}
}

func crashHandler() {
	if r := recover(); r != nil {
		glog.Warningf("unhandled panic %s. recovering to keep the process alive...", r)
	}
}
