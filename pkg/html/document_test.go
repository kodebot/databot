package html

import (
	"testing"

	"golang.org/x/net/html"
)

func TestNewDocument(t *testing.T) {
	tests := []string{"<html></html>", "this is test string"}

	for _, test := range tests {
		actual := NewDocument(test)

		if actual == nil {
			t.Fatalf("document creation failed")
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		input     string
		selectors []string
		expected  string
	}{
		{
			"<html><head></head><body><div><span>removeme</span></div></body></html>",
			[]string{"span"},
			"<html><head></head><body><div></div></body></html>"},
		{
			"<html><head></head><body><div><span class='remove0'>removeme</span><span>dontremoveme</span><span class='remove1'>removeme</span></div></body></html>",
			[]string{".remove0", ".remove1"},
			"<html><head></head><body><div><span>dontremoveme</span></div></body></html>"},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		doc.Remove(test.selectors...)
		actual := doc.HTML()
		if actual != test.expected {
			t.Fatalf("remove failed for selector %v. EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.selectors, test.expected, actual)
		}
	}
}

func TestSelect(t *testing.T) {
	tests := []struct {
		input     string
		selectors []string
		expected  string
	}{
		{
			`<html><head></head><body><div><span>keepme</span></div></body></html>`,
			[]string{`span`},
			`<span>keepme</span>`},
		{
			`<html><head></head><body><div><span class="keep0">keepme</span><span>dontkeepme</span><span class="keep1">keepme</span></div></body></html>`,
			[]string{`.keep0`, `.keep1`},
			`<span class="keep0">keepme</span><span class="keep1">keepme</span>`},
		{
			`<html><head></head><body><div><span>keepme</span></div></body></html>`,
			[]string{`.keepme`},
			``},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		doc.Select(test.selectors...)
		actual := doc.HTML()

		if actual != test.expected {
			t.Fatalf("keep failed for selector %v. EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.selectors, test.expected, actual)
		}
	}
}

func TestBody(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`<html><head></head><body><div><span>keepme</span></div></body></html>`,
			`<div><span>keepme</span></div>`},
		{
			`<html><head></head><body></body></html>`,
			``},
		{
			`<div><span></span></div>`,
			`<div><span></span></div>`},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		actual := doc.Body()
		if actual != test.expected {
			t.Fatalf("retrieving body failed. EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.expected, actual)
		}
	}
}

func TestHTML(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`<html><head></head><body><div><span>keepme</span></div></body></html>`,
			`<html><head></head><body><div><span>keepme</span></div></body></html>`},
		{
			`<html><head></head><body></body></html>`,
			`<html><head></head><body></body></html>`},
		{
			`<div><span></span></div>`,
			`<html><head></head><body><div><span></span></div></body></html>`},
		{
			``,
			`<html><head></head><body></body></html>`},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		actual := doc.HTML()
		if actual != test.expected {
			t.Fatalf("retrieving HTML of the document failed. EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.expected, actual)
		}
	}
}

func TestRemoveAttrs(t *testing.T) {
	tests := []struct {
		input    string
		attrs    []string
		expected string
	}{
		{
			`<html><head></head><body><div foo="bar" baz="qux"></div></body></html>`,
			[]string{"foo"},
			`<html><head></head><body><div baz="qux"></div></body></html>`},
		{
			`<html><head></head><body><div foo="bar" baz="qux"></div></body></html>`,
			[]string{"foo", "baz"},
			`<html><head></head><body><div></div></body></html>`},
		{
			`<html><head></head><body><div foo="bar" ></div></body></html>`,
			[]string{"foo", "baz"},
			`<html><head></head><body><div></div></body></html>`},
		{
			`<html><head></head><body><div ></div></body></html>`,
			[]string{"foo", "baz"},
			`<html><head></head><body><div></div></body></html>`},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		doc.RemoveAttrs(test.attrs...)
		actual := doc.HTML()
		if actual != test.expected {
			t.Fatalf("remove attributes failed. EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.expected, actual)
		}
	}
}

func TestRemoveAttrsWhen(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		whenFn   func(*html.Attribute) bool
		expected string
	}{
		{
			"remove attr foo conditionally (foo appears as first attr)",
			`<html><head></head><body><div foo="bar" baz="qux"></div></body></html>`,
			func(attr *html.Attribute) bool { return attr.Key == "foo" },
			`<html><head></head><body><div baz="qux"></div></body></html>`},
		{
			"remove attr foo conditionally (foo appears as last attrs)",
			`<html><head></head><body><div baz="qux" foo="bar" ></div></body></html>`,
			func(attr *html.Attribute) bool { return attr.Key == "foo" },
			`<html><head></head><body><div baz="qux"></div></body></html>`},
		{
			"remove all attrs (more than one present)",
			`<html><head></head><body><div baz="qux" foo="bar" ></div></body></html>`,
			func(attr *html.Attribute) bool { return true },
			`<html><head></head><body><div></div></body></html>`},
		{
			"remove all attrs (just one present)",
			`<html><head></head><body><div foo="bar" ></div></body></html>`,
			func(attr *html.Attribute) bool { return true },
			`<html><head></head><body><div></div></body></html>`},
		{
			"remove attrs (no attrs present)",
			`<html><head></head><body><div ></div></body></html>`,
			func(attr *html.Attribute) bool { return true },
			`<html><head></head><body><div></div></body></html>`},
		{
			"don't remove attrs (multiple attrs present)",
			`<html><head></head><body><div foo="bar" baz="qux"></div></body></html>`,
			func(attr *html.Attribute) bool { return false },
			`<html><head></head><body><div foo="bar" baz="qux"></div></body></html>`},
		{
			"don't remove attrs (no attrs present)",
			`<html><head></head><body><div ></div></body></html>`,
			func(attr *html.Attribute) bool { return false },
			`<html><head></head><body><div></div></body></html>`},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		doc.RemoveAttrsWhen(test.whenFn)
		actual := doc.HTML()
		if actual != test.expected {
			t.Fatalf("remove attributes when failed (%s). EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.name, test.expected, actual)
		}
	}
}

func TestRemoveNonContent(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`<html><head></head><body><div></div></body></html>`,
			`<html></html>`,
		},
		{
			`<html><head></head><body><div>something</div></body></html>`,
			`<html><body><div>something</div></body></html>`,
		},
		{
			`<html><head></head><body><img></img></body></html>`,
			`<html><body><img/></body></html>`,
		},
		{
			`<html><head></head><body><br/></body></html>`,
			`<html><body><br/></body></html>`,
		},
		{
			`<html><head></head><body><div>something</div><!-- comment --></body></html>`,
			`<html><body><div>something</div></body></html>`,
		},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		doc.RemoveNonContent()
		actual := doc.HTML()
		if actual != test.expected {
			t.Fatalf("remove non content failed. EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.expected, actual)
		}
	}
}

func TestRemoveNodesWhen(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		whenFn   func(*html.Node) bool
		expected string
	}{
		{
			"remove node foo conditionally",
			`<html><head></head><body><foo></foo><bar></bar></body></html>`,
			func(node *html.Node) bool { return node.Data == "foo" },
			`<html><head></head><body><bar></bar></body></html>`},
		{
			"don't remove node when condition not met",
			`<html><head></head><body><foo></foo><bar></bar></body></html>`,
			func(node *html.Node) bool { return false },
			`<html><head></head><body><foo></foo><bar></bar></body></html>`},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		doc.RemoveNodesWhen(test.whenFn)
		actual := doc.HTML()
		if actual != test.expected {
			t.Fatalf("remove attributes when failed (%s). EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.name, test.expected, actual)
		}
	}
}

func TestGetMetadata(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		keyAttr  string
		keyVal   string
		valAttr  string
		expected string
	}{
		{
			"key matched and valAttr present with value",
			`<meta foo="bar" baz="qux"/>`,
			"foo",
			"bar",
			"baz",
			"qux"},
		{
			"key matched and valAttr present with no value",
			`<meta foo="bar" baz=""/>`,
			"foo",
			"bar",
			"baz",
			""},
		{
			"key matched and valAttr not present",
			`<meta foo="bar"/>`,
			"foo",
			"bar",
			"baz",
			""},
		{
			"keyAttr matched but keyVal not matched ",
			`<meta foo="bar" baz="qux"/>`,
			"foo",
			"quxx",
			"baz",
			""},
		{
			"keyAttr not matched",
			`<meta foo="bar" baz="qux"/>`,
			"quxx",
			"bar",
			"baz",
			""},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		actual := doc.GetMetadata(test.keyAttr, test.keyVal, test.valAttr)
		if actual != test.expected {
			t.Fatalf("get metadata failed (%s). EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.name, test.expected, actual)
		}
	}
}
