package main

import (
	"testing"

	"github.com/kodebot/databot/pkg/config"
	"github.com/kodebot/databot/pkg/exporter"
	"github.com/kodebot/databot/pkg/record"

	"github.com/kodebot/databot/pkg/databot"

	"github.com/kodebot/databot/pkg/toml"
)

func Test(t *testing.T) {

	confBuilder := config.NewBuilder()
	confBuilder.UseEnv()
	confBuilder.Build()

	feedSpecReader := toml.FeedSpecReader{}
	feed := feedSpecReader.ReadFile("feedconfig copy.toml")

	var recCreator databot.RecordCreator
	recCreator = record.NewRecordCreator()

	recs := recCreator.Create(feed.RecordSpec)
	outputPath := "./result.txt"
	exporter.ExportToTextFile(recs, outputPath)
	exporter.ExportToMongoDB(recs, config.Current().ExportToDBConStr())
}
