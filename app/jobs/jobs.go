package jobs

import (
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/models"
	"github.com/kodebot/newsfeed/services"
)

// LoadArticlesFromFeedsJob job
type LoadArticlesFromFeedsJob struct{}

// PruneArticlesJob job
type PruneArticlesJob struct{}

// Run LoadArticlesFromFeedsJob
func (j LoadArticlesFromFeedsJob) Run() {
	var feedConfig models.FeedConfig
	_, err := toml.DecodeFile("./feed_config.toml", &feedConfig)
	if err != nil {
		glog.Errorf("error when loading feed config: %s\n", err.Error())
	}

	glog.Infoln("running LoadArticlesFromFeedsJob...")
	for _, feed := range feedConfig.Feed {
		glog.Infof("processing feed from %s \n", feed.URL)
		result := services.ParseFeed(feed)
		if result == nil {
			glog.Errorln("feed skipped...")
			continue
		}

		articles := services.CreateArticles(result, feed)
		if len(articles) == 0 {
			glog.Warning("no articles found...")
			continue
		}
		services.LoadArticles(articles)
	}

	glog.Infoln("finished LoadArticlesFromFeedsJob...")
}

// Run PruneArticlesJob
func (j PruneArticlesJob) Run() {
	services.PruneArticles()
}
