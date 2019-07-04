package rssatom

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mmcdole/gofeed"
)

func TestPluck(t *testing.T) {
	tests := []struct {
		name     string
		source   *gofeed.Item
		params   map[string]interface{}
		expected interface{}
	}{
		{"source is nil", nil, nil, nil},
		{"params is nil", &gofeed.Item{}, nil, nil},
		{"params is empty", &gofeed.Item{}, map[string]interface{}{}, nil},
		{"params has no source", &gofeed.Item{}, map[string]interface{}{"foo": "bar"}, nil},
		{"params has invalid source", &gofeed.Item{Content: "foo"}, map[string]interface{}{"source": "bar"}, nil},
		{"params has invalid source", &gofeed.Item{Content: "foo"}, map[string]interface{}{"source": "bar"}, nil},
		{"params has incorrect case of source field", &gofeed.Item{Content: "foo"}, map[string]interface{}{"source": "content"}, nil},
		{"source and params valid", &gofeed.Item{Content: "foo"}, map[string]interface{}{"source": "Content"}, "foo"},
	}

	for _, test := range tests {
		actual := pluck(test.source, test.params)
		assert.Equal(t, test.expected, actual, test.name)
	}
}
