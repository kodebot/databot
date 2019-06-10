package transformers

import "testing"

type StringTransformerTest struct {
	input      interface{}
	expected   interface{}
	parameters map[string]interface{}
}

var trimTests = []StringTransformerTest{
	{" ", "", nil},
	{"test ", "test", nil},
	{134, 134, nil},
	{nil, nil, nil}}

func TestTrim(t *testing.T) {
	for _, test := range trimTests {
		actual := trim(test.input, test.parameters)
		if test.expected != actual {
			fail(t, "trim not working", test.expected, actual)
		}
	}
}

var trimLeftTests = []StringTransformerTest{
	{" ", "", nil},
	{" test", "test", nil},
	{" test ", "test ", nil},
	{134, 134, nil},
	{nil, nil, nil}}

func TestLeftTrim(t *testing.T) {
	for _, test := range trimLeftTests {
		actual := trimLeft(test.input, test.parameters)
		if test.expected != actual {
			fail(t, "trimLeft not working", test.expected, actual)
		}
	}
}

var trimRightTests = []StringTransformerTest{
	{" ", "", nil},
	{"test ", "test", nil},
	{" test ", " test", nil},
	{134, 134, nil},
	{nil, nil, nil}}

func TestRightTrim(t *testing.T) {
	for _, test := range trimRightTests {
		actual := trimRight(test.input, test.parameters)
		if test.expected != actual {
			fail(t, "trimRight not working", test.expected, actual)
		}
	}
}
