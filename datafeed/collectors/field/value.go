package field

import (
	"github.com/golang/glog"
)

// ValueCollector collects data using regexp
type ValueCollector struct {
	Collector
}

// Collect returns regexp collected value
func (c *ValueCollector) Collect() interface{} {
	glog.Infof("collecting from %s using value collector from", c.Field)
	return c.Field
}
