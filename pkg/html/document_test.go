package html

import (
	"testing"
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
		actual, _ := doc.document.Html()
		if actual != test.expected {
			t.Fatalf("remove failed for selector %v. EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.selectors, test.expected, actual)
		}
	}
}

func TestKeep(t *testing.T) {
	tests := []struct {
		input     string
		selectors []string
		expected  string
	}{
		{
			"<html><head></head><body><div><span>keepme</span></div></body></html>",
			[]string{"span"},
			"<span>keepme</span>"},
		{
			`<html><head></head><body><div><span class="keep0">keepme</span><span>dontkeepme</span><span class="keep1">keepme</span></div></body></html>`,
			[]string{".keep0", ".keep1"},
			`<span class="keep0">keepme</span><span class="keep1">keepme</span>`},
	}

	for _, test := range tests {
		doc := NewDocument(test.input)
		doc.Keep(test.selectors...)
		actual, err := doc.document.Html()
		if err != nil {
			t.Fatal(err.Error())
		}
		if actual != test.expected {
			t.Fatalf("keep failed for selector %v. EXPECTED: <<%s>>, ACTUAL: <<%s>>", test.selectors, test.expected, actual)
		}
	}
}
