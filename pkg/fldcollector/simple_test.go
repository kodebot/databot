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
