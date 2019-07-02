package fldxfmr

import (
	"testing"
)

func TestInit(t *testing.T) {

	if len(TransformersMap) <= 0 {
		t.Fatalf("TransformersMap is expected to be initialised but found not")
	}
}
