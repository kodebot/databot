package fieldcollector

import "github.com/kodebot/databot/pkg/databot"

// Collector method signature
type Collector func(source *interface{}, params map[string]interface{}) interface{}

// CollectorMap contains common collectors that can be used for multiple sources
var CollectorMap = map[databot.FieldCollectorType]Collector{
	databot.ValueCollector: value}
