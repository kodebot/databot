package fldxfmr

import (
	"testing"

	"github.com/kodebot/databot/pkg/html"
)

func TestRemoveHTMLElements(t *testing.T) {
	negativeTests := []struct {
		name     string
		input    interface{}
		expected interface{}
		params   map[string]interface{}
	}{
		{"input val is nil", nil, nil, nil},
		{"input val is not string", 1, nil, nil},
		{"params is nil", "foo", nil, nil},
		{"params has no selectors", "foo", nil, map[string]interface{}{}},
		{"params has empty selectors", "foo", nil, map[string]interface{}{"selectors": []interface{}{}}},
	}

	for _, test := range negativeTests {
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return html.NewMockDocument()
		}}
		actual := htmlCtx.removeHTMLElements(test.input, test.params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	positiveTests := []struct {
		name      string
		input     interface{}
		selectors []string
		expected  interface{}
	}{{"remove selectors and return doc html", "foo", []string{"bar", "baz"}, "qux"}}

	for _, test := range positiveTests {
		selectors := make([]interface{}, len(test.selectors))
		for i, v := range test.selectors {
			selectors[i] = v
		}
		params := map[string]interface{}{"selectors": selectors}
		mockDocument := html.NewMockDocument()
		mockDocument.On("Remove", test.selectors).Return()
		mockDocument.On("HTML").Return(test.expected)
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		actual := htmlCtx.removeHTMLElements(test.input, params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}
}
