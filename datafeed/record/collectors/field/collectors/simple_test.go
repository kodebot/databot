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

var emptyTests = []CollectorTest{
	{"string", "", nil},
	{1234, "", nil},
	{nil, "", nil}}

func TestEmpty(t *testing.T) {
	for _, test := range emptyTests {
		actual := empty(test.input, test.parameters)
		if test.expected != actual {
			fail(t, "empty collector not working", test.expected, actual)
		}
	}
}
