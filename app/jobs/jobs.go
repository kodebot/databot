package jobs

import (
	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/datafeed"
	"github.com/kodebot/newsfeed/services"
)

// LoadArticlesFromFeedsJob job
type LoadArticlesFromFeedsJob struct{}

// PruneArticlesJob job
type PruneArticlesJob struct{}

// Run LoadArticlesFromFeedsJob
func (j LoadArticlesFromFeedsJob) Run() {
	dataFeed, dataFeedSetting := datafeed.ParseFromDataFeedSetting("./conf/new_feed_configs/001_dinamalar_politics.toml")
	articles := services.CreateArticles(dataFeed)
	if len(articles) == 0 {
		glog.Warning("no articles found...")
	} else {

		for _, article := range articles {
			article.Source = dataFeedSetting.SourceName
			article.Categories = []string{dataFeedSetting.Category}
		}

		services.LoadArticles(articles)
	}

	glog.Infoln("finished LoadArticlesFromFeedsJob...")
}

// Run PruneArticlesJob
func (j PruneArticlesJob) Run() {
	services.PruneArticles()
}
