package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kodebot/databot/pkg/config"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/exporter"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/kodebot/databot/pkg/record"
	"github.com/kodebot/databot/pkg/toml"
	"github.com/robfig/cron"
)

type previewResp struct {
	Config string
	Recs   []map[string]interface{}
}

func previewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	if r.Method == "GET" {
		config, _ := ioutil.ReadFile("feedconfig.toml")
		resp := previewResp{Config: string(config), Recs: []map[string]interface{}{}}
		t.Execute(w, resp)
		return
	}

	config := r.FormValue("config")

	feedSpecReader := toml.FeedSpecReader{}
	feed := feedSpecReader.Read(config)

	var recCreator databot.RecordCreator
	recCreator = record.NewRecordCreator()
	rspec := feed.RecordSpec
	recs := recCreator.Create(rspec)

	resp := previewResp{config, recs}
	t.Execute(w, resp)
}

func main() {
	// todo: keep all feed specs in database

	runonce := flag.Bool("runonce", false, "processes the feeds once outside the schedule and exits")
	feedConfigPath := flag.String("feedconfigpath", "./feeds/ready/", "specifies the location of config files to process. The processes all the config files in the specified directory and its subdirectories recursively.")

	if !strings.HasSuffix(*feedConfigPath, "/") {
		x := (*feedConfigPath) + "/"
		feedConfigPath = &x
	}

	flag.Parse()

	confBuilder := config.NewBuilder()
	confBuilder.UseEnv()
	confBuilder.Build()

	processFeeds(*runonce, *feedConfigPath)

	if !*runonce {
		http.HandleFunc("/", previewHandler)
		logger.Infof("Launching web UI on port 9025")
		logger.Fatalf(http.ListenAndServe(":9025", nil).Error())
	}
}

func processFeeds(runonce bool, feedConfigPath string) {
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
	} else {
		logger.Infof("all feed specs processed successfully")
	}
}

func processFeed(feedSpec databot.FeedSpec, recCreator databot.RecordCreator) {
	rspec := feedSpec.RecordSpec
	recs := recCreator.Create(rspec)
	exporter.ExportToMongoDB(recs, config.Current().ExportToDBConStr())
}
