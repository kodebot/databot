package rssatom

import (
	"github.com/kodebot/databot/pkg/databot"
	"github.com/kodebot/databot/pkg/fldcollector"
	"github.com/kodebot/databot/pkg/fldtransformer"
	"github.com/kodebot/databot/pkg/logger"
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
	return transform(collected, c.TransformerSpecs)
}

func (c *fieldFactory) collect() interface{} {
	collectorType := c.CollectorSpec.Type

	// for RSS/Atom feed set the collector type to Pluck if not specified
	if collectorType == "" {
		collectorType = fldcollector.PluckField
	}

	if collector := collectorMap[collectorType]; collector != nil {
		return collector(c.RssAtomItem, c.CollectorSpec.Params)
	}

	if sharedCollector := fldcollector.CollectorMap[collectorType]; sharedCollector != nil {
		var source interface{} = *c.RssAtomItem
		return sharedCollector(&source, c.CollectorSpec.Params)
	}

	logger.Errorf("specified collector %s is not found", collectorType)
	return nil

}

func transform(value interface{}, transformerSpecs []*databot.FieldTransformerSpec) interface{} {
	for _, spec := range transformerSpecs {
		if transformerFn := transformersMap[spec.Type]; transformerFn != nil {
			value = transformerFn(value, spec.Params)
			continue
		}

		if sharedTransformerFn := fldtransformer.TransformersMap[spec.Type]; sharedTransformerFn != nil {
			value = sharedTransformerFn(value, spec.Params)
			continue
		}

		logger.Errorf("transform %s not found", spec.Type)
	}
	return value
}
