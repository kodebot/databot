package jobs

import (
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/kodebot/newsfeed/articles"
	"github.com/kodebot/newsfeed/data"
	"github.com/kodebot/newsfeed/datafeed"
)

// LoadArticlesFromFeedsJob job
type LoadArticlesFromFeedsJob struct {
	FeedInfoPath string
}

// PruneArticlesJob job
type PruneArticlesJob struct{}

// Run LoadArticlesFromFeedsJob
func (j LoadArticlesFromFeedsJob) Run() {
	articleCollection, err := data.GetCollection("articles")
	if err != nil {
		glog.Errorf("error while loading articles collection %s", err.Error())
		return
	}

	feedConfigPath := j.FeedInfoPath
	if feedConfigPath == "" {
		feedConfigPath = "./conf/feed/ready/"
	}

	filepath.Walk(feedConfigPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			glog.Errorf("error while processing feed %s. error: %s", path, err.Error())
			return err
		}

		if info.IsDir() {
			return nil
		}

		glog.Infof("loading articles using %s", path)
		dataFeed, feedInfo := datafeed.NewFromFeedInfo(path)

		if len(dataFeed) == 0 {
			glog.Warning("no articles found...")
			return nil
		}

		for _, dataFeedItem := range dataFeed {
			newArticle := articles.NewArticle(dataFeedItem)
			newArticle.Source = feedInfo.SourceName
			newArticle.Categories = []string{feedInfo.Category}
			err := newArticle.Store(articleCollection)
			if err != nil {
				glog.Errorf("error while storing article %s", err.Error())
			}
		}
		return nil
	})
	glog.Infoln("finished LoadArticlesFromFeedsJob...")
}

// Run PruneArticlesJob
func (j PruneArticlesJob) Run() {
	// todo:
}
