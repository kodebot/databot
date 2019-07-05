package fldcollector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {

	tests := []struct {
		params   map[string]interface{}
		expected interface{}
	}{
		{nil, nil},
		{map[string]interface{}{}, nil},
		{map[string]interface{}{"foo": "bar"}, nil},
		{map[string]interface{}{"value": "foo"}, "foo"},
	}

	for _, test := range tests {
		actual := value(nil, test.params)
		assert.Equal(t, test.expected, actual)
	}
}

func TestSource(t *testing.T) {

	tests := []struct {
		source   interface{}
		params   map[string]interface{}
		expected interface{}
	}{
		{nil, nil, nil},
		{"foo", map[string]interface{}{}, "foo"},
		{1, map[string]interface{}{"foo": "bar"}, 1},
		{struct{}{}, nil, struct{}{}},
	}

	for _, test := range tests {
		actual := source(test.source, test.params)
		assert.EqualValues(t, test.expected, actual)
	}
}
