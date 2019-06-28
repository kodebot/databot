package rssatom

import (
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/mmcdole/gofeed"
)

// RecordFactory represents a model that enables retrieving record(s) from RSS/Atom feed
type RecordFactory struct {
	*databot.RecordSpec
	RssAtomFeed *gofeed.Feed
}

// NewRecordFactory returns a new record factory that enables creating one or more records using RSS/Atom feed
func NewRecordFactory(recordSpec *databot.RecordSpec) *RecordFactory {
	sourceURI := recordSpec.CollectorSpec.SourceURI

	xml, err := html.ReadAsString(sourceURI)
	if err != nil {
		logger.Errorf("unable to retrieve content from URI %s", sourceURI)
	}
	rssAtomFeed := Parse(xml)
	return &RecordFactory{recordSpec, rssAtomFeed}
}

// Create returns one or more records from given rss/atom record spec
func (r *RecordFactory) Create() []map[string]interface{} {
	recs := r.collect()
	// todo: review whether it is ok to collect all records and transform or we need to collect and transform one record at a time
	// todo: no record transformers are supported now
	// transformed := applyRecTransformers(collected, nil)
	return recs
}

func (r *RecordFactory) collect() []map[string]interface{} {
	recs := []map[string]interface{}{}
	for _, item := range r.RssAtomFeed.Items { // nothing to collect at record level - the feed item is already available
		rec := make(map[string]interface{})
		for _, field := range r.RecordSpec.FieldSpecs {
			normalise(field)
			f := newFieldFactory(field, item)
			rec[field.Name] = f.create()
		}
		recs = append(recs, rec)
	}
	return recs
}

// func applyRecTransformers(rec *map[string]*interface{}, transformers []*databot.TransformerSpec) *map[string]*interface{} {
// 	// todo: apply record transformers if any
// 	return rec
// }

func normalise(field *databot.FieldSpec) {
	// initialise params if nil
	if params := field.CollectorSpec.Params; params == nil {
		field.CollectorSpec.Params = make(map[string]interface{})
	}

	// set source same as name if missing
	if src := field.CollectorSpec.Params["source"]; src == nil {
		field.CollectorSpec.Params["source"] = field.Name
	}
}
