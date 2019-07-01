package fieldcollector

import "github.com/kodebot/databot/pkg/databot"

// Collector method signature
type Collector func(source *interface{}, params map[string]interface{}) interface{}

// CollectorMap contains common collectors that can be used for multiple sources
var CollectorMap = map[databot.FieldCollectorType]Collector{
	Value: value}

const (
	// Value returns the specified constant value.
	// All sources are supported
	Value databot.FieldCollectorType = "value"
	// PluckField represents a type of collector that fetches value from property of an instance
	// Supports rssatom source
	PluckField databot.FieldCollectorType = "pluck"
)
