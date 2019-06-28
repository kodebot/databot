package cache

import (
	"github.com/kodebot/databot/pkg/databot"
)

// todo: think about concurrency
type memCache struct {
	cache map[interface{}]interface{}
}

// NewMemCache returns new in memory cache
func NewMemCache() databot.CacheService {
	return memCache{cache: make(map[interface{}]interface{})}
}

func (c memCache) Get(key interface{}) interface{} {
	return c.cache[key]
}

func (c memCache) Add(key interface{}, val interface{}) {
	c.cache[key] = val
}
func (c memCache) Reset() {
	c.cache = make(map[interface{}]interface{})
}
func (c memCache) Prune() {
	// noop
	// todo: for LRU cache this need implementing
}
