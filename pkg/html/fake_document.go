package html

import "golang.org/x/net/html"

type fakeDocument struct{}

// NewFakeDocument returns a model that represents HTML document
func NewFakeDocument() Document {
	return &fakeDocument{}
}

// Remove removes the elements from the document that matches given selectors
func (d *fakeDocument) Remove(selectors ...string) {
}

// Select only keeps the matching elements in the document
func (d *fakeDocument) Select(selectors ...string) {
}

// Body returns the document body as string
func (d *fakeDocument) Body() string {
	return ""
}

// HTML returns the document body as string
func (d *fakeDocument) HTML() string {
	return ""
}

// RemoveAttrs removes the specified attribute
func (d *fakeDocument) RemoveAttrs(attrs ...string) {
}

// RemoveAttrsWhen removes the specified attribute when the given condition is met
func (d *fakeDocument) RemoveAttrsWhen(when func(attr string, val string) bool) {
}

// RemoveNonContent removes all empty elements including comment elements
func (d *fakeDocument) RemoveNonContent() {
}

// RemoveNodeWhen removes all the nodes matching the given condition
func (d *fakeDocument) RemoveNodeWhen(when func(node *html.Node) bool) {
}

// GetMetadata retrives the value of valAttr attribute from meta element that that has keyAttr attribute with valAttr value
func (d *fakeDocument) GetMetadata(keyAttr string, keyVal string, valAttr string) string {
	return ""
}
