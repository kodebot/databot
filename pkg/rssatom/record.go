package rssatom

import (
	"github.com/kodebot/databot/pkg/databot"
	"github.com/mmcdole/gofeed"
)

// RecordEngine represents a model that enables retrieving record(s) from RSS/Atom feed
type RecordEngine struct {
	*databot.RecordSpec
	RssAtomFeed *gofeed.Feed
}

// NewRecordEngine return new rss/atom record type
func NewRecordEngine(recordSpec *databot.RecordSpec, rssAtomFeed *gofeed.Feed) *RecordEngine {
	return &RecordEngine{recordSpec, rssAtomFeed}
}

// CreateRecords returns one or more records from given rss/atom record spec
func (r *RecordEngine) CreateRecords() []*map[string]*interface{} {
	recs := []*map[string]*interface{}{}

	for _, item := range r.RssAtomFeed.Items {
		// todo: collector guards
		collected := r.collect(item)
		// todo: no record transformers are supported now
		transformed := applyRecTransformers(collected, nil)
		recs = append(recs, transformed)
	}
	return recs
}

func (r *RecordEngine) collect(item *gofeed.Item) *map[string]*interface{} {
	rec := make(map[string]*interface{})
	for _, field := range r.RecordSpec.FieldSpecs {
		normalise(field)
		src := (*(*field.CollectorSpec.Params)["source"]).(string)
		f := newFieldEngine(field, item)
		rec[src] = f.createField()
	}
	return &rec
}

func applyRecTransformers(rec *map[string]*interface{}, transformers []*databot.TransformerSpec) *map[string]*interface{} {
	// todo: apply record transformers if any
	return rec
}

func normalise(field *databot.FieldSpec) {
	// initialise params if nil
	if params := field.CollectorSpec.Params; params == nil {
		prm := (make(map[string]*interface{}))
		field.CollectorSpec.Params = &prm
	}

	// set source same as name if missing
	if src := (*field.CollectorSpec.Params)["source"]; src == nil {
		var fieldName interface{} = field.Name
		(*field.CollectorSpec.Params)["source"] = &fieldName
	}
}
