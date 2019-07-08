package main

import (
	"errors"
	"testing"

	"github.com/kodebot/databot/pkg/config"
	"github.com/kodebot/databot/pkg/exporter"

	"github.com/kodebot/databot/pkg/databot"

	"github.com/kodebot/databot/pkg/reccollector"
	"github.com/kodebot/databot/pkg/rssatom"
	"github.com/kodebot/databot/pkg/toml"
)

func Test(t *testing.T) {

	confBuilder := config.NewBuilder()
	confBuilder.UseEnv()
	confBuilder.Build()

	feedSpecReader := toml.FeedSpecReader{}
	feed := feedSpecReader.ReadFile("feedconfig.toml")

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
	outputPath := "./result.txt"
	exporter.ExportToTextFile(recs, outputPath)
	exporter.ExportToMongoDB(recs, config.Current().ExportToDBConStr())
}
