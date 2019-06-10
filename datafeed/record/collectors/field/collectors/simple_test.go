package collectors

import "testing"


var valueTests = []CollectorTest{
	{"string", "string", nil},
	{1234, 1234, nil},
	{nil, nil, nil}}

func TestValue(t *testing.T) {
	for _, test := range valueTests {
		actual := value(test.input, test.parameters)
		if test.expected != actual {
			fail(t, "value collector not working", test.expected, actual)
		}
	}
}
