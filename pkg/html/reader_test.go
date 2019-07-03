package html

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAsString(t *testing.T) {

	tests := []struct {
		httpError  bool
		statusCode int
		body       string
		expected   string
	}{
		{true, 200, "bar", ""},
		{false, 200, "bar", "bar"},
		{false, 401, "bar", ""},
	}

	for _, test := range tests {
		httpServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			if test.httpError {
				panic(errors.New("foo"))
			}
			res.WriteHeader(test.statusCode)
			res.Write([]byte(test.body))
		}))
		defer func() { httpServer.Close() }()

		reader := NewDocumentReader("http://bar")
		actual, _ := reader.ReadAsString()
		assert.Equal(t, test.expected, actual)
	}
}
