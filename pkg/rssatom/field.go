package rssatom

import (
	"github.com/kodebot/databot/pkg/databot"
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

func (c *fieldFactory) create() *interface{} {
	if c.RssAtomItem == nil {
		logger.Errorf("Cannot collect field value when RssAtomItem is nil")
		return nil
	}

	collected := c.collect()
	return applyFieldTransformers(collected, c.TransformerSpecs)
}

func (c *fieldFactory) collect() *interface{} {
	collectorType := c.CollectorSpec.Type

	// for RSS/Atom feed set the collector type to Pluck if not specified
	if collectorType == "" {
		collectorType = databot.PluckFieldCollector
	}

	collector := collectorMap[collectorType]
	if collector == nil {
		logger.Errorf("specified collector %s is missing implementation", collectorType)
		return nil
	}

	return collector(c.RssAtomItem, c.CollectorSpec.Params)
}

func applyFieldTransformers(source *interface{}, transformerSpecs []*databot.FieldTransformerSpec) *interface{} {
	// todo
	return source
}
