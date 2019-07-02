package fldxfmr

import (
	"testing"
	"time"
)

// TransformerTest helps to create table driven tests for transformers
type TransformerTest struct {
	input    interface{}
	expected interface{}
	params   map[string]interface{}
}

func fail(t *testing.T, msg string, expected interface{}, actual interface{}) {
	t.Helper()
	t.Fatalf("%#v. EXPECTED: >>%s<<, ACTUAL: >>%#v<<", msg, expected, actual)
}

func compareDateStr(expected interface{}, actual interface{}) bool {
	if (expected == nil || actual == nil) && expected == actual {
		return true
	}
	if actualTime, ok := actual.(time.Time); ok {
		x := actualTime.String()
		return expected == x
	}

	if actualTimePtr, ok := actual.(*time.Time); ok {
		return expected == actualTimePtr.String()
	}
	return expected == actual // non dates
}
