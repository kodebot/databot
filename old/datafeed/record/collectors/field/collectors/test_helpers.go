package collectors

import "testing"

// CollectorTest helps to create table driven tests for collectors
type CollectorTest struct {
	input      interface{}
	expected   interface{}
	parameters map[string]interface{}
}

func fail(t *testing.T, message string, expected interface{}, actual interface{}) {
	t.Helper()
	t.Fatalf("%s. EXPECTED: >>%s<<, ACTUAL: >>%s<<", message, expected, actual)
}
