package transformers

import (
	"testing"
	"time"
)

// TransformerTest helps to create table driven tests for transformers
type TransformerTest struct {
	input      interface{}
	expected   interface{}
	parameters map[string]interface{}
}

func fail(t *testing.T, message string, expected interface{}, actual interface{}) {
	t.Helper()
	t.Fatalf("%s. EXPECTED: >>%s<<, ACTUAL: >>%s<<", message, expected, actual)
}

func compareDateStr(expected interface{}, actual interface{}) bool {
	if (expected == nil || actual == nil) && expected == actual {
		return true
	}
	if actualTime, ok := actual.(time.Time); ok {
		return expected == actualTime.String()
	}

	if actualTimePointer, ok := actual.(*time.Time); ok {
		return expected == actualTimePointer.String()
	}
	return expected == actual // non dates
}
