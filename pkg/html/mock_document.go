package html

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/html"
)

type MockDocument struct {
	mock.Mock
}

// NewMockDocument returns a model that represents HTML document
func NewMockDocument() *MockDocument {
	return new(MockDocument)
}

func (d *MockDocument) Remove(selectors ...string) {
	d.Called(selectors)
}

func (d *MockDocument) Select(selectors ...string) {
	d.Called(selectors)
}

func (d *MockDocument) Body() string {
	args := d.Called()
	return args.String(0)
}

func (d *MockDocument) HTML() string {
	args := d.Called()
	return args.String(0)
}

func (d *MockDocument) RemoveAttrs(attrs ...string) {
	d.Called(attrs)
}

func (d *MockDocument) RemoveAttrsWhen(when func(attr string, val string) bool) {
	d.Called(when)
}

func (d *MockDocument) RemoveNonContent() {
	d.Called()
}

func (d *MockDocument) RemoveNodeWhen(when func(node *html.Node) bool) {
	d.Called(when)
}

func (d *MockDocument) GetMetadata(keyAttr string, keyVal string, valAttr string) string {
	args := d.Called(keyAttr, keyVal, valAttr)
	return args.String(0)
}
