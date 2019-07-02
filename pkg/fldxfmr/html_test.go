package fldxfmr

import (
	"testing"

	"github.com/kodebot/databot/pkg/html"
)

func TestRemoveHTMLElements(t *testing.T) {

	htmlCtx := htmlContext{newDocFn: func(s string) html.Document {
		return html.NewFakeDocument()
	}}

	tests := []struct{ name string }{{
		name: "dummy test"}}

	for _, test := range tests {
		actual := htmlCtx.removeHTMLElements(nil, nil)
		fail(t, test.name, nil, actual)
	}

}
