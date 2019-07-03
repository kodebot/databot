package fldtransformer

import "testing"

func TestTrim(t *testing.T) {
	tests := []TransformerTest{
		{" ", "", nil},
		{"test ", "test", nil},
		{134, nil, nil},
		{nil, nil, nil}}

	for _, test := range tests {
		actual := trim(test.input, test.params)
		if test.expected != actual {
			fail(t, "trim not working", test.expected, actual)
		}
	}
}

func TestLeftTrim(t *testing.T) {
	tests := []TransformerTest{
		{" ", "", nil},
		{" test", "test", nil},
		{" test ", "test ", nil},
		{134, nil, nil},
		{nil, nil, nil}}

	for _, test := range tests {
		actual := trimLeft(test.input, test.params)
		if test.expected != actual {
			fail(t, "trimLeft not working", test.expected, actual)
		}
	}
}

func TestRightTrim(t *testing.T) {
	tests := []TransformerTest{
		{" ", "", nil},
		{"test ", "test", nil},
		{" test ", " test", nil},
		{134, nil, nil},
		{nil, nil, nil}}

	for _, test := range tests {
		actual := trimRight(test.input, test.params)
		if test.expected != actual {
			fail(t, "trimRight not working", test.expected, actual)
		}
	}
}

func TestReplace(t *testing.T) {
	tests := []TransformerTest{
		{" ", nil, nil},
		{"", nil, nil},
		{nil, nil, nil},
		{"foobarbaz", nil, nil},
		{"foobarbaz", nil, map[string]interface{}{}},
		{"foobarbaz", nil, map[string]interface{}{"old": "bar"}},
		{"foobarbaz", nil, map[string]interface{}{"new": "qux"}},
		{"foobarbaz", nil, map[string]interface{}{"old": 1, "new": "qux"}},
		{"foobarbaz", nil, map[string]interface{}{"old": "bar", "new": 1}},
		{"foobarbaz", "fooquxbaz", map[string]interface{}{"old": "bar", "new": "qux"}},
		{"foobarbaz", "foobaz", map[string]interface{}{"old": "bar", "new": ""}},
		{"foobarbazfoobarbaz", "fooquxbazfoobarbaz", map[string]interface{}{"old": "bar", "new": "qux"}},
		{"foobarbazfoobarbaz", "fooquxquxfoobarbaz", map[string]interface{}{"old": []interface{}{"bar", "baz"}, "new": "qux"}},
		{"foobarbazfoobarbaz", "fooquxbazfooquxbaz", map[string]interface{}{"old": "bar", "new": "qux", "all": true}},
		{"foobarbazfoobarbaz", "fooquxquxfooquxqux", map[string]interface{}{"old": []interface{}{"bar", "baz"}, "new": "qux", "all": true}}}

	for _, test := range tests {
		actual := replace(test.input, test.params)
		if test.expected != actual {
			fail(t, "replace not working", test.expected, actual)
		}
	}
}

func TestReplaceAll(t *testing.T) {
	tests := []TransformerTest{
		{" ", nil, nil},
		{"", nil, nil},
		{nil, nil, nil},
		{"foobarbaz", nil, nil},
		{"foobarbaz", nil, map[string]interface{}{}},
		{"foobarbaz", nil, map[string]interface{}{"old": "bar"}},
		{"foobarbaz", nil, map[string]interface{}{"new": "qux"}},
		{"foobarbaz", nil, map[string]interface{}{"old": 1, "new": "qux"}},
		{"foobarbaz", nil, map[string]interface{}{"old": "bar", "new": 1}},
		{"foobarbaz", "fooquxbaz", map[string]interface{}{"old": "bar", "new": "qux"}},
		{"foobarbaz", "foobaz", map[string]interface{}{"old": "bar", "new": ""}},
		{"foobarbazfoobarbaz", "fooquxbazfooquxbaz", map[string]interface{}{"old": "bar", "new": "qux"}},
		{"foobarbazfoobarbaz", "fooquxquxfooquxqux", map[string]interface{}{"old": []interface{}{"bar", "baz"}, "new": "qux"}}}

	for _, test := range tests {
		actual := replaceAll(test.input, test.params)
		if test.expected != actual {
			fail(t, "replace not working", test.expected, actual)
		}
	}
}

func TestRegexpSelect(t *testing.T) {
	regexpTests := []TransformerTest{
		{nil, nil, nil},
		{"no regex", nil, nil},
		{"regex without data group", nil, map[string]interface{}{"expr": ".*result.*"}},
		{"invalid regex", nil, map[string]interface{}{"expr": "??..*result.*"}},
		{"get result from this", "result", map[string]interface{}{"expr": ".*(?P<data>result).*"}},
		{"foo", nil, map[string]interface{}{"expr": ".*(?P<data>result).*"}},
		{nil, nil, map[string]interface{}{"expr": ".*(?P<data>result).*"}}}

	for _, test := range regexpTests {
		actual := regexpSelect(test.input, test.params)
		if test.expected != actual {
			fail(t, "regexp collector not working", test.expected, actual)
		}
	}
}
