package rssatom

import (
	"github.com/kodebot/databot/pkg/databot"
	"github.com/mmcdole/gofeed"
)

// Record represents the config to create record from rss/atom feed item
type Record struct {
	*databot.Record
	RssAtomFeed *gofeed.Feed
}

// Collect returns one or more records from given rss data
func (r *Record) Collect() []*map[string]*interface{} {
	recs := []*map[string]*interface{}{}
	for _, item := range r.RssAtomFeed.Items {
		rec := make(map[string]*interface{})
		for _, field := range r.Record.Fields {
			normalise(field)
			src := (*(*field.Collector.Params)["source"]).(string)
			f := Field{field, item}
			rec[src] = f.Collect()
		}

		recs = append(recs, &rec)
	}
	return recs
}

func normalise(field *databot.Field) {
	// initialise params if nil
	if params := field.Collector.Params; params == nil {
		prm := (make(map[string]*interface{}))
		field.Collector.Params = &prm
	}

	// set source same as name if missing
	if src := (*field.Collector.Params)["source"]; src == nil {
		var fieldName interface{} = field.Name
		(*field.Collector.Params)["source"] = &fieldName
	}
}
