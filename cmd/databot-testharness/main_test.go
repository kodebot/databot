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
	rspec := feed.RecordSpec
	recs := recCreator.Create(rspec)
	outputPath := "./result.txt"
	exporter.ExportToTextFile(recs, outputPath)
	exporter.ExportToMongoDB(recs, config.Current().ExportToDBConStr())
}

// func TestGoroutine(t *testing.T) {
// 	preprocessors := []*databot.ProcessorSpec{
// 		{
// 			Name:   "http:get",
// 			Params: map[string]interface{}{},
// 		},
// 		{
// 			Name:   "css:select",
// 			Params: map[string]interface{}{"selectors": []string{"html"}},
// 		},
// 	}

// 	input := make(chan interface{})
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		var pipeline <-chan interface{}
// 		pipeline = input
// 		for _, spec := range preprocessors {
// 			nextProcessor := processor.Get(spec.Name)
// 			pipeline = nextProcessor(pipeline, spec.Params)
// 		}
// 		wg.Done()
// 		for out := range pipeline {
// 			t.Errorf("output: %+v", out)
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	wg.Add(1)
// 	input <- "https://www.digitalocean.com/pricing/"
// 	close(input)
// 	wg.Wait()
// 	t.FailNow()
// }
