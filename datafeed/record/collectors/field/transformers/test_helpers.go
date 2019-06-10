package transformers

import "testing"

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
