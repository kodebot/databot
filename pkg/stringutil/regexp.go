package stringutil

import (
	"regexp"

	"github.com/kodebot/databot/pkg/logger"
)

// RegexpMatchAll returns all the matches from val using expr
// when captured group name called 'data' is present, the only the 'data' sub-group is returned
func RegexpMatchAll(val string, expr string) []string {
	result := []string{}
	if val != "" {
		if expr == "" {
			logger.Errorf("no regular expression found")
			return result
		}

		re, err := regexp.Compile(expr)
		if err != nil {
			logger.Errorf("invalid regexp: %s error: %s. \n", expr, err.Error())
			return result
		}

		requiredMatchIndex := -1
		for i, val := range re.SubexpNames() {
			if val == "data" {
				requiredMatchIndex = i
			}
		}

		if requiredMatchIndex > -1 {
			matches := re.FindAllStringSubmatch(val, -1)
			if matches == nil || len(matches) < 1 {
				logger.Warnf("no match found.")
			}

			for _, m := range matches {
				if len(m) < requiredMatchIndex+1 {
					logger.Warnf("no match found.")
					return result
				}
				result = append(result, m[requiredMatchIndex])
			}
		} else { // when there is no captured group 'data' - just return the whole match
			result = re.FindAllString(val, -1)
		}
	}
	return result
}
