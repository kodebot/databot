package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kodebot/databot/pkg/config"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/record"
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
	recCreator = record.NewRecordCreator()
	rspec := feed.RecordSpec
	recs := recCreator.Create(rspec)

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
