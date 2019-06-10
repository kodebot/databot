package jobs

import (
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/articles"
	"github.com/kodebot/newsfeed/data"
	"github.com/kodebot/newsfeed/datafeed"
)

// LoadArticlesFromFeedsJob job
type LoadArticlesFromFeedsJob struct{}

// PruneArticlesJob job
type PruneArticlesJob struct{}

// Run LoadArticlesFromFeedsJob
func (j LoadArticlesFromFeedsJob) Run() {
	articleCollection, err := data.GetCollection("articles")
	if err != nil {
		glog.Errorf("error while loading articles collection %s", err.Error())
		return
	}

	feedConfigPath := "./conf/feed/ready/"
	files, err := ioutil.ReadDir(feedConfigPath)

	for _, file := range files {
		fullPath := feedConfigPath + file.Name()
		glog.Infof("loading articles using %s", fullPath)
		dataFeed, dataFeedSetting := datafeed.NewFromFeedInfo(fullPath)

		if len(dataFeed) == 0 {
			glog.Warning("no articles found...")
			return
		}

		for _, dataFeedItem := range dataFeed {
			newArticle := articles.NewArticle(dataFeedItem)
			newArticle.Source = dataFeedSetting.SourceName
			newArticle.Categories = []string{dataFeedSetting.Category}
			err := newArticle.Store(articleCollection)
			if err != nil {
				glog.Errorf("error while storing article %s", err.Error())
			}
		}
	}
	glog.Infoln("finished LoadArticlesFromFeedsJob...")
}

// Run PruneArticlesJob
func (j PruneArticlesJob) Run() {
	// todo:
}
