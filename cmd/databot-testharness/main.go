package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/kodebot/databot/pkg/config"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/reccollector"
	"github.com/kodebot/databot/pkg/rssatom"
	"github.com/kodebot/databot/pkg/toml"
)

func handler(w http.ResponseWriter, r *http.Request) {
	feedSpecReader := toml.FeedSpecReader{}
	feed := feedSpecReader.Read("feedconfig.toml")

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
	t, _ := template.ParseFiles("home.html")

	t.Execute(w, recs)
}

func main() {
	confBuilder := config.NewBuilder()
	confBuilder.UseEnv()
	confBuilder.Build()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9022", nil))
}
