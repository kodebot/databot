package collectors

import (
	"testing"

	"github.com/mmcdole/gofeed"
)

func TestRssatom(t *testing.T) {
	tests := []struct {
		input    interface{}
		params   map[string]interface{}
		expected interface{}
	}{
		{nil, nil, nil},
		{"input", nil, nil},
		{&gofeed.Item{}, nil, nil},
		{&gofeed.Item{Title: "foo"}, map[string]interface{}{}, nil},
		{&gofeed.Item{Title: "foo"}, map[string]interface{}{"source": "Title"}, "foo"}}

	for _, test := range tests {
		actual := rssatom(test.input, test.params)
		if test.expected != actual {
			fail(t, "rssatom collector doesn't work as it should", test.expected, actual)
		}
	}
}
