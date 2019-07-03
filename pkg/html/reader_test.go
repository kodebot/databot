package html

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAsString(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		panic(errors.New("foo"))
	}))
	defer func() { httpServer.Close() }()

	reader := NewDocumentReader("bar")
	actual, _ := reader.ReadAsString()
	assert.Equal(t, "", actual)

}
