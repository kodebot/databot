package collectors

import "github.com/kodebot/newsfeed/logger"

// CollectorInfo settings for collecting field
type CollectorInfo struct {
	Type       CollectorType
	Parameters map[string]interface{}
}

// CollectorType provides available field collector types
type CollectorType string

const (
	// RssAtomField collects value from rss/atom field
	RssAtomField CollectorType = "rssAtomField"

	// CSS field collects values using css selector from html document
	CSS CollectorType = "css"

	// Unknown field collector
	Unknown CollectorType = "unknown"
)

type collectorFuncType func(value interface{}, parameters map[string]interface{}) interface{}

var collectorsMap map[CollectorType]collectorFuncType

func init() {
	collectorsMap = map[CollectorType]collectorFuncType{
		RssAtomField: rssatom}
}

// Collect value from the source
func Collect(source interface{}, info CollectorInfo) interface{} {
	collector := collectorsMap[info.Type]

	if collector != nil {
		return collector(source, info.Parameters)
	}

	logger.Warnf("invalid collector type %s", info.Type)
	return nil
}
