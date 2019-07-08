package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kodebot/databot/pkg/config"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/reccollector"
	"github.com/kodebot/databot/pkg/rssatom"
	"github.com/kodebot/databot/pkg/toml"
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
	resp := previewResp{config, recs}
	t.Execute(w, resp)
}

func main() {
	confBuilder := config.NewBuilder()
	confBuilder.UseEnv()
	confBuilder.Build()

	http.HandleFunc("/", previewHandler)
	log.Fatal(http.ListenAndServe(":9022", nil))
}
