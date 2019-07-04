package rssatom

import (
	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/fldcollector"
	"github.com/kodebot/databot/pkg/fldtransformer"
	"github.com/kodebot/databot/pkg/html"
	"github.com/kodebot/databot/pkg/logger"
	"github.com/mmcdole/gofeed"
)

type recordCreator struct {
	cacheManagerFn func() cache.Manager
	docReaderFn    func(string, cache.Manager) html.DocumentReader
}

// NewRecordCreator returns a new record creator that enables creating one or more records using RSS/Atom feed
func NewRecordCreator() databot.RecordCreator {
	return &recordCreator{docReaderFn: html.NewCachedDocumentReader, cacheManagerFn: cache.Current}
}

// Create returns one or more records from given rss/atom record spec
func (r *recordCreator) Create(spec *databot.RecordSpec) []map[string]interface{} {

	sourceURI := spec.CollectorSpec.SourceURI

	docReader := r.docReaderFn(sourceURI, r.cacheManagerFn())
	xml, err := docReader.ReadAsString()
	if err != nil {
		logger.Errorf("unable to retrieve content from URI %s", sourceURI)
	}
	source := Parse(xml)

	recs := collect(source, spec)
	// todo: review whether it is ok to collect all records and transform or we need to collect and transform one record at a time
	// todo: no record transformers are supported now
	// transformed := applyRecTransformers(collected, nil)
	return recs
}

func collect(source *gofeed.Feed, spec *databot.RecordSpec) []map[string]interface{} {
	recs := []map[string]interface{}{}
	for _, item := range source.Items { // nothing to collect at record level - the feed item is already available
		rec := make(map[string]interface{})
		for _, fldSpec := range spec.FieldSpecs {
			normaliseFieldSpec(fldSpec)
			rec[fldSpec.Name] = createField(item, fldSpec)
		}
		recs = append(recs, rec)
	}
	return recs
}

// func applyRecTransformers(rec *map[string]*interface{}, transformers []*databot.TransformerSpec) *map[string]*interface{} {
// 	// todo: apply record transformers if any
// 	return rec
// }

func normaliseFieldSpec(field *databot.FieldSpec) {
	// initialise params if nil
	if params := field.CollectorSpec.Params; params == nil {
		field.CollectorSpec.Params = make(map[string]interface{})
	}

	// set source same as name if missing
	if src := field.CollectorSpec.Params["source"]; src == nil {
		field.CollectorSpec.Params["source"] = field.Name
	}
}

func createField(source *gofeed.Item, spec *databot.FieldSpec) interface{} {
	if source == nil {
		logger.Errorf("Cannot collect field value when RssAtomItem is nil")
		return nil
	}

	collected := collectField(source, spec.CollectorSpec)
	return transformField(collected, spec.TransformerSpecs)
}

func collectField(source *gofeed.Item, spec *databot.FieldCollectorSpec) interface{} {
	collectorType := spec.Type

	// for RSS/Atom feed set the collector type to Pluck if not specified
	if collectorType == "" {
		collectorType = fldcollector.PluckField
	}

	if collector := collectorMap[collectorType]; collector != nil {
		return collector(source, spec.Params)
	}

	if sharedCollector := fldcollector.CollectorMap[collectorType]; sharedCollector != nil {
		var source interface{} = *source
		return sharedCollector(&source, spec.Params)
	}

	logger.Errorf("specified collector %s is not found", collectorType)
	return nil
}

func transformField(value interface{}, specs []*databot.FieldTransformerSpec) interface{} {
	for _, spec := range specs {
		if transformerFn := transformersMap[spec.Type]; transformerFn != nil {
			value = transformerFn(value, spec.Params)
			continue
		}

		if sharedTransformerFn := fldtransformer.TransformersMap[spec.Type]; sharedTransformerFn != nil {
			value = sharedTransformerFn(value, spec.Params)
			continue
		}

		logger.Errorf("transformer %s not found", spec.Type)
	}
	return value
}
