package transformers

import (
	"strings"
	"time"
)

func formatDate(value time.Time, parameters map[string]string) string {
	// todo:
	return value.String()
}

func trim(value string, parameters map[string]string) string {

	return strings.TrimSpace(value)
}
