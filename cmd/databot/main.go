package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/kodebot/databot/pkg/config"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/exporter"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/record"
	"github.com/kodebot/databot/pkg/toml"
	"github.com/robfig/cron"
)

func main() {
	// todo: keep all feed specs in database
	// todo: change this to be scheduled job

	runonce := flag.Bool("runonce", false, "processes the feeds once outside the schedule and exits")
	flag.Parse()

	confBuilder := config.NewBuilder()
	confBuilder.UseEnv()
	confBuilder.Build()

	processFeeds(*runonce)
}

func processFeeds(runonce bool) {
	feedConfigPath := "./feeds/ready/"
	c := cron.New()
	filepath.Walk(feedConfigPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Errorf("error while processing feed %s. error: %s", path, err.Error())
			return err
		}

		if info.IsDir() {
			return nil
		}

		logger.Infof("processing feed spec %s", path)

		feedSpecReader := toml.FeedSpecReader{}
		feed := feedSpecReader.ReadFile(path)
		var recCreator databot.RecordCreator
		recCreator = record.NewRecordCreator()

		if runonce {
			processFeed(feed, recCreator)
			logger.Infof("feed spec %s is processed successfully", path)
		} else {
			c.AddFunc(feed.Schedule, func() { processFeed(feed, recCreator) })
			logger.Infof("feed spec %s is scheduled successfully", path)
		}
		return nil
	})

	if !runonce {
		logger.Infof("starting feed schedules")
		c.Start()
		logger.Infof("started feed schedules successfully")
		quit := make(chan bool)
		for {
			select {
			case shouldQuit := <-quit:
				if shouldQuit {
					break
				}
			}
		}
	} else {
		logger.Infof("feed specs processed successfully")
	}
}

func processFeed(feedSpec databot.FeedSpec, recCreator databot.RecordCreator) {
	rspec := feedSpec.RecordSpec
	recs := recCreator.Create(rspec)
	exporter.ExportToMongoDB(recs, config.Current().ExportToDBConStr())
}
