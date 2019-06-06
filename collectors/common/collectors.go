package common

import (
	"regexp"

	"github.com/golang/glog"
)

// CollectUsingRegexp collects field data using regexp
// the expression must have only one named group called 'data'
func CollectUsingRegexp(source string, expr string) string {

	glog.Infof("collecting from %s using regexp selector %s", source, expr)
	re, err := regexp.Compile(expr)
	if err != nil {
		glog.Errorf("invalid regexp: %s error: %s. \n", expr, err.Error())
		return ""
	}

	requiredMatchIndex := 0
	for i, val := range re.SubexpNames() {
		if val == "data" {
			requiredMatchIndex = i
		}
	}

	if requiredMatchIndex == 0 {
		glog.Errorf("invalid regular expression: %s no named group called 'data' is found. \n", expr)
		return ""
	}
	matches := re.FindStringSubmatch(source)
	if len(matches) < requiredMatchIndex+1 {
		glog.Warningf("no match found.")
		return ""
	}

	return matches[requiredMatchIndex]
}
