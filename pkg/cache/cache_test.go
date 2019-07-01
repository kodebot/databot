package cache

import "testing"

type TestT struct {
	cache *Manager
}

func TestCurrent(t *testing.T) {
	t1 := TestT{cache: Current()}

	(*(t1.cache)).Add("1", "one")
	result := (*(t1.cache)).Get("1")
	t.Errorf(result.(string))
}
