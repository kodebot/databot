package transformers

import "testing"

func fail(t *testing.T, message string, expected interface{}, actual interface{}) {
	t.Helper()
	t.Fatalf("%s. EXPECTED: >>%s<<, ACTUAL: >>%s<<", message, expected, actual)
}
