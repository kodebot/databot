package rssatom

import (
	fieldcollector "github.com/kodebot/databot/pkg/collectors/field"
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/logger"
	fieldtransformer "github.com/kodebot/databot/pkg/transformers/field"
	"github.com/mmcdole/gofeed"
)

// fieldFactory represents a model that enables getting fields from given RSS/Atom item
type fieldFactory struct {
	*databot.FieldSpec
	RssAtomItem *gofeed.Item
}

// newFieldFactory returns new RSS/Atom field factory
func newFieldFactory(fieldSpec *databot.FieldSpec, rssAtomItem *gofeed.Item) *fieldFactory {
	return &fieldFactory{fieldSpec, rssAtomItem}
}

func (c *fieldFactory) create() interface{} {
	if c.RssAtomItem == nil {
		logger.Errorf("Cannot collect field value when RssAtomItem is nil")
		return nil
	}

	collected := c.collect()
	return applyFieldTransformers(collected, c.TransformerSpecs)
}

func (c *fieldFactory) collect() interface{} {
	collectorType := c.CollectorSpec.Type

	// for RSS/Atom feed set the collector type to Pluck if not specified
	if collectorType == "" {
		collectorType = databot.PluckFieldCollector
	}

	if collector := collectorMap[collectorType]; collector != nil {
		return collector(c.RssAtomItem, c.CollectorSpec.Params)
	}

	if sharedCollector := fieldcollector.CollectorMap[collectorType]; sharedCollector != nil {
		var source interface{} = *c.RssAtomItem
		return sharedCollector(&source, c.CollectorSpec.Params)
	}

	logger.Errorf("specified collector %s is not found", collectorType)
	return nil

}

func applyFieldTransformers(value interface{}, transformerSpecs []*databot.FieldTransformerSpec) interface{} {
	return fieldtransformer.Transform(value, transformerSpecs)
}
