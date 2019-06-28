package cache

import (
	"github.com/kodebot/databot/pkg/databot"
)

// todo: update this to use MongoDB for persistence
// todo: think about concurrency

type dbCache struct {
	cache map[interface{}]interface{}
}

// NewDbCache returns new in memory cache
func NewDbCache() databot.CacheService {
	return dbCache{cache: make(map[interface{}]interface{})}
}

func (c dbCache) Get(key interface{}) interface{} {
	return c.cache[key]
}

func (c dbCache) Add(key interface{}, val interface{}) {
	c.cache[key] = val
}
func (c dbCache) Reset() {
	c.cache = make(map[interface{}]interface{})
}
func (c dbCache) Prune() {
	// todo: for LRU cache this need implementing
}
