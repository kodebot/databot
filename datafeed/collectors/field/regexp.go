package field

import (
	"regexp"

	"github.com/golang/glog"
)

// RegexpCollector collects data using regexp
type RegexpCollector struct {
	Collector
}

// Collect returns regexp collected value
func (c Collector) Collect() interface{} {

	glog.Infof("collecting from %s using regexp", c.Field)

	fieldRawValueString := c.Field.(string)
	var expr string
	if ok := c.Parameters["expr"]; ok != nil {
		expr = ok.(string)
	} else {
		glog.Errorf("no regular expression parameter found")
		return nil
	}

	re, err := regexp.Compile(expr)
	if err != nil {
		glog.Errorf("invalid regexp: %s error: %s. \n", expr, err.Error())
		return nil
	}

	requiredMatchIndex := 0
	for i, val := range re.SubexpNames() {
		if val == "data" {
			requiredMatchIndex = i
		}
	}

	if requiredMatchIndex == 0 {
		glog.Errorf("invalid regular expression: %s no named group called 'data' is found. \n", expr)
		return nil
	}
	matches := re.FindStringSubmatch(fieldRawValueString)
	if len(matches) < requiredMatchIndex+1 {
		glog.Warningf("no match found.")
		return nil
	}

	if found := matches[requiredMatchIndex]; found != "" {
		var result interface{} = found
		return result
	}
	return nil

}
