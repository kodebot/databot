package cache

import (
	"github.com/kodebot/databot/pkg/config"
)

// Manager is abstract cache lookup service
type Manager interface {
	Get(key string) interface{}
	Add(key string, val interface{})
	Reset()
	Prune()
}

var currentCache Manager

// Current returns the currently configured cache
func Current() Manager {
	if currentCache == nil {
		if config.Current().UseDBCache() {
			currentCache = NewDBCache()
		} else {
			currentCache = NewMemCache()
		}
	}

	return currentCache
}
