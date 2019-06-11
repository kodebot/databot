package transformers

import "testing"

var valueTests = []TransformerTest{
	{"string", "string", nil},
	{1234, 1234, nil},
	{nil, nil, nil}}

func TestValue(t *testing.T) {
	for _, test := range valueTests {
		actual := value(test.input, test.parameters)
		if test.expected != actual {
			fail(t, "value transformer not working", test.expected, actual)
		}
	}
}

var emptyTests = []TransformerTest{
	{"string", "", nil},
	{1234, "", nil},
	{nil, "", nil}}

func TestEmpty(t *testing.T) {
	for _, test := range emptyTests {
		actual := empty(test.input, test.parameters)
		if test.expected != actual {
			fail(t, "empty transformer not working", test.expected, actual)
		}
	}
}
