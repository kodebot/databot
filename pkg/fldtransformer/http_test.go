package fldtransformer

import (
	"errors"
	"testing"

	"github.com/kodebot/databot/pkg/cache"
	"github.com/kodebot/databot/pkg/html"
)

type fakeDocReader struct {
	success bool
}

func (r *fakeDocReader) ReadAsString() (string, error) {
	if r.success {
		return "foo", nil
	}
	return "", errors.New("test error bar")
}

func TestHttpGet(t *testing.T) {
	var tests = []struct {
		name          string
		input         interface{}
		expected      interface{}
		params        map[string]interface{}
		readerSuccess bool
	}{
		{"all nil value test", nil, nil, nil, true},
		{"non string input value test", 1, nil, nil, true},
		{"positive test", "", "foo", nil, true},
		{"reader failure test", "", nil, nil, false}}

	for _, test := range tests {
		fakeCtx := httpContext{docReaderFn: func(u string, cache *cache.Manager) html.DocumentReader {
			return &fakeDocReader{success: test.readerSuccess}
		}}

		actual := fakeCtx.httpGet(test.input, test.params)
		if actual != test.expected {
			fail(t, test.name, test.expected, actual)
		}
	}
}
