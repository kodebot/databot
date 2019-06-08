package transformers

import (
	"strings"
	"time"
)

func formatDate(value time.Time, parameters map[string]interface{}) string {
	// todo:
	return value.String()
}

func trim(value string, parameters map[string]interface{}) string {

	return strings.TrimSpace(value)
}
