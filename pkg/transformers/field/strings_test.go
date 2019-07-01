package fieldtransformer

import "testing"

var trimTests = []TransformerTest{
	{" ", "", nil},
	{"test ", "test", nil},
	{134, 134, nil},
	{nil, nil, nil}}

func TestTrim(t *testing.T) {
	for _, test := range trimTests {
		actual := trim(test.input, test.params)
		if test.expected != actual {
			fail(t, "trim not working", test.expected, actual)
		}
	}
}

var trimLeftTests = []TransformerTest{
	{" ", "", nil},
	{" test", "test", nil},
	{" test ", "test ", nil},
	{134, 134, nil},
	{nil, nil, nil}}

func TestLeftTrim(t *testing.T) {
	for _, test := range trimLeftTests {
		actual := trimLeft(test.input, test.params)
		if test.expected != actual {
			fail(t, "trimLeft not working", test.expected, actual)
		}
	}
}

var trimRightTests = []TransformerTest{
	{" ", "", nil},
	{"test ", "test", nil},
	{" test ", " test", nil},
	{134, 134, nil},
	{nil, nil, nil}}

func TestRightTrim(t *testing.T) {
	for _, test := range trimRightTests {
		actual := trimRight(test.input, test.params)
		if test.expected != actual {
			fail(t, "trimRight not working", test.expected, actual)
		}
	}
}
