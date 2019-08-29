package cache

import (
	"sync"

	"github.com/kodebot/databot/pkg/config"
)

var lock sync.Mutex

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
	lock.Lock()
	defer lock.Unlock()
	if currentCache == nil {
		if config.Current().UseDBCache() {
			currentCache = NewDBCache()
		} else {
			currentCache = NewMemCache()
		}
	}

	return currentCache
}
