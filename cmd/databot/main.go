package main

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/kodebot/databot/pkg/config"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/exporter"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/reccollector"
	"github.com/kodebot/databot/pkg/rssatom"
	"github.com/kodebot/databot/pkg/toml"
)

func main() {
	// todo: keep all feed specs in database
	// todo: change this to be scheduled job

	runonce := flag.Bool("runonce", false, "processes the feeds once outside the schedule and exits")
	flag.Parse()

	confBuilder := config.NewBuilder()
	confBuilder.UseEnv()
	confBuilder.Build()
	if *runonce {
		logger.Infof("processing feeds only once outside the schedule")
		processFeeds()
	} else {
		logger.Infof("scheduling feeds for processing")
		schedule()
	}
}

func schedule() {
	ticker := time.NewTicker(30 * time.Minute)
	quit := make(chan bool)

	for {
		select {
		case <-ticker.C:
			print("process the files")

		case shouldQuit := <-quit:
			if shouldQuit {
				break
			}
		}
	}
}

func processFeeds() {

	feedConfigPath := "./feeds/ready/"

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
		switch feed.RecordSpec.CollectorSpec.Type {
		case reccollector.RssAtom:
			recCreator = rssatom.NewRecordCreator()
		case reccollector.HTMLSingle:
			panic(errors.New("HTMLSingle record collector is not implemented"))
		case reccollector.HTML:
			panic(errors.New("HTMLMultiple record collector is not implemented"))
		default:
			panic(errors.New("Unsupported record collector"))
		}

		recs := recCreator.Create(feed.RecordSpec)
		exporter.ExportToMongoDB(recs, config.Current().ExportToDBConStr())
		logger.Infof("feed spec %s is processed successfully", path)
		return nil
	})
	logger.Infof("feed specs processed successfully")
}
