package fldxfmr

import (
	"fmt"
	"testing"

	"github.com/kodebot/databot/pkg/html"
	"github.com/stretchr/testify/mock"
	gohtml "golang.org/x/net/html"
)

func TestRemoveElements(t *testing.T) {
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
		actual := htmlCtx.removeElements(test.input, test.params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	functionalTests := []struct {
		name      string
		input     interface{}
		selectors []string
		expected  interface{}
	}{{"remove elements matching selectors and return doc html", "foo", []string{"bar", "baz"}, "qux"}}

	for _, test := range functionalTests {
		params := make(map[string]interface{})
		addSelectorsToParam(params, test.selectors)
		mockDocument := html.NewMockDocument()
		mockDocument.On("Remove", test.selectors).Return()
		mockDocument.On("HTML").Return(test.expected)
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		actual := htmlCtx.removeElements(test.input, params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}
}

func TestSelectElements(t *testing.T) {
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
		actual := htmlCtx.selectElements(test.input, test.params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	functionalTests := []struct {
		name      string
		input     interface{}
		selectors []string
		expected  interface{}
	}{{"only keep elements matching selectors and return doc html", "foo", []string{"bar", "baz"}, "qux"}}

	for _, test := range functionalTests {
		params := make(map[string]interface{})
		addSelectorsToParam(params, test.selectors)
		mockDocument := html.NewMockDocument()
		mockDocument.On("Select", test.selectors).Return()
		mockDocument.On("HTML").Return(test.expected)
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		actual := htmlCtx.selectElements(test.input, params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}
}

func TestRemoveStyles(t *testing.T) {
	negativeTests := []struct {
		name     string
		input    interface{}
		expected interface{}
		params   map[string]interface{}
	}{
		{"input val is nil", nil, nil, nil},
		{"input val is not string", 1, nil, nil}}

	for _, test := range negativeTests {
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return html.NewMockDocument()
		}}
		actual := htmlCtx.removeStyles(test.input, test.params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	functionalTests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{{"remove html styles and return doc html", "foo", "qux"}}

	for _, test := range functionalTests {
		mockDocument := html.NewMockDocument()
		mockDocument.On("Remove", []string{"style"}).Return()
		mockDocument.On("RemoveAttrs", []string{"style", "class"}).Return()
		mockDocument.On("HTML").Return(test.expected)
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		actual := htmlCtx.removeStyles(test.input, nil)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}
}

func TestRemoveScripts(t *testing.T) {
	negativeTests := []struct {
		name     string
		input    interface{}
		expected interface{}
		params   map[string]interface{}
	}{
		{"input val is nil", nil, nil, nil},
		{"input val is not string", 1, nil, nil}}

	for _, test := range negativeTests {
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return html.NewMockDocument()
		}}
		actual := htmlCtx.removeScripts(test.input, test.params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	functionalTests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{{"remove html scripts and return doc html", "foo", "qux"}}

	for _, test := range functionalTests {
		mockDocument := html.NewMockDocument()
		mockDocument.On("Remove", []string{"script"}).Return()
		mockDocument.On("RemoveAttrsWhen", mock.Anything).Return()
		mockDocument.On("HTML").Return(test.expected)
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		actual := htmlCtx.removeScripts(test.input, nil)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	attrMatcherTests := []struct {
		name     string
		key      string
		val      string
		expected bool
	}{{"remove data- attrs", "data-foo", "", true},
		{"remove when value contains 'javascript:' (starts with)", "foo", "javascript:bar", true},
		{"remove when value contains 'javascript:' (contains)", "foo", "foo:javascript:bar", true},
		{"remove when value contains 'javascript:' (ends with)", "foo", "foo:javascript:", true},
		{"do not remove - neither data- nor value contains 'javascript:'", "foo", "foo-bar", false}}

	for _, test := range attrMatcherTests {
		mockDocument := html.NewMockDocument()
		mockDocument.On("Remove", []string{"script"}).Return()
		var actual bool
		mockDocument.On("RemoveAttrsWhen", mock.AnythingOfTypeArgument("func(*html.Attribute) bool")).Return().Run(func(args mock.Arguments) {
			when := args.Get(0).(func(*gohtml.Attribute) bool)
			actual = when(&gohtml.Attribute{Key: test.key, Val: test.val})
		})

		mockDocument.On("HTML").Return("bar")
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		htmlCtx.removeScripts("foo", nil)

		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}
}

func TestRemoveNonContentElements(t *testing.T) {
	negativeTests := []struct {
		name     string
		input    interface{}
		expected interface{}
		params   map[string]interface{}
	}{
		{"input val is nil", nil, nil, nil},
		{"input val is not string", 1, nil, nil},
	}

	for _, test := range negativeTests {
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return html.NewMockDocument()
		}}
		actual := htmlCtx.removeNonContentElements(test.input, test.params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	functionalTests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{{"remove non content and return doc html", "foo", "qux"}}

	for _, test := range functionalTests {
		mockDocument := html.NewMockDocument()
		mockDocument.On("RemoveNonContent").Return()
		mockDocument.On("HTML").Return(test.expected)
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		actual := htmlCtx.removeNonContentElements(test.input, nil)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}
}

func TestRemoveElementsMatchingText(t *testing.T) {
	negativeTests := []struct {
		name     string
		input    interface{}
		expected interface{}
		params   map[string]interface{}
	}{
		{"input val is nil", nil, nil, nil},
		{"input val is not string", 1, nil, nil},
		{"params is nil", "foo", nil, nil},
		{"params has no matchers", "foo", nil, map[string]interface{}{}},
		{"params has empty matchers", "foo", nil, map[string]interface{}{"matchers": []interface{}{}}}}

	for _, test := range negativeTests {
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return html.NewMockDocument()
		}}
		actual := htmlCtx.removeElementsMatchingText(test.input, test.params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	functionalTests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{{"remove html elements with matching text element and return doc html", "foo", "qux"}}

	for _, test := range functionalTests {
		mockDocument := html.NewMockDocument()
		mockDocument.On("Remove", []string{"script"}).Return()
		mockDocument.On("RemoveNodesWhen", mock.Anything).Return()
		mockDocument.On("HTML").Return(test.expected)
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		actual := htmlCtx.removeElementsMatchingText(test.input, map[string]interface{}{"matchers": []interface{}{"bar"}})
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	nodeMatcherTests := []struct {
		name     string
		matchers []interface{}
		val      string
		expected bool
	}{
		{"contains match", []interface{}{"bar"}, "foobarbaz", true},
		{"starts with match", []interface{}{"^foo"}, "foobarbaz", true},
		{"ends with match", []interface{}{"baz$"}, "foobarbaz", true},
		{"at least one match", []interface{}{"qux", "baz"}, "foobarbaz", true},
		{"no match found", []interface{}{"quxx"}, "foobarbaz", false},
	}

	for _, test := range nodeMatcherTests {
		mockDocument := html.NewMockDocument()
		mockDocument.On("Remove", []string{"script"}).Return()
		var actual bool
		mockDocument.On("RemoveNodesWhen", mock.AnythingOfTypeArgument("func(*html.Node) bool")).Return().Run(func(args mock.Arguments) {
			when := args.Get(0).(func(*gohtml.Node) bool)
			actual = when(&gohtml.Node{Data: test.val})
		})

		mockDocument.On("HTML").Return("bar")
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		htmlCtx.removeElementsMatchingText("foo", map[string]interface{}{"matchers": test.matchers})

		if test.expected != actual {
			fail(t, fmt.Sprintf("%s. matchers:%v, target:%s", test.name, test.matchers, test.val), test.expected, actual)
		}
	}
}

func TestGetMetadata(t *testing.T) {
	negativeTests := []struct {
		name     string
		input    interface{}
		expected interface{}
		params   map[string]interface{}
	}{
		{"input val is nil", nil, nil, nil},
		{"input val is not string", 1, nil, nil},
		{"params is nil", "foo", nil, nil},
		{"params is empty", "foo", nil, map[string]interface{}{}},
		{"params has no keyAttr", "foo", nil, map[string]interface{}{"bar": []interface{}{}}},
		{"params has no keyVal", "foo", nil, map[string]interface{}{"keyAttr": "bar"}},
		{"params has no valAttr", "foo", nil, map[string]interface{}{"keyAttr": "bar", "keyVal": "baz"}},
		{"params has nil for keyAttr", "foo", nil, map[string]interface{}{"keyAttr": nil, "keyVal": "bar", "valAttr": "baz"}},
	}

	for _, test := range negativeTests {
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return html.NewMockDocument()
		}}
		actual := htmlCtx.getMetadata(test.input, test.params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}

	functionalTests := []struct {
		name     string
		input    interface{}
		params   map[string]interface{}
		expected interface{}
	}{{"get metdata from html", "foo", map[string]interface{}{"keyAttr": "bar", "keyVal": "baz", "valAttr": "qux"}, "quxx"}}

	for _, test := range functionalTests {
		mockDocument := html.NewMockDocument()
		mockDocument.On("GetMetadata", test.params["keyAttr"].(string), test.params["keyVal"].(string), test.params["valAttr"].(string)).Return(test.expected)
		mockDocument.On("HTML").Return(test.expected)
		htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
			return mockDocument
		}}

		actual := htmlCtx.getMetadata(test.input, test.params)
		if test.expected != actual {
			fail(t, test.name, test.expected, actual)
		}
	}
}

func addSelectorsToParam(params map[string]interface{}, strs []string) {
	selectors := make([]interface{}, len(strs))
	for i, v := range strs {
		selectors[i] = v
	}
	params["selectors"] = selectors
}
