package cache

// todo: update this to use MongoDB for persistence
// todo: think about concurrency

type dbCache struct {
	cache map[string]interface{}
}

// NewDbCache returns new in memory cache
func NewDbCache() Manager {
	return &dbCache{cache: make(map[string]interface{})}
}

func (c *dbCache) Get(key string) interface{} {
	return c.cache[key]
}

func (c *dbCache) Add(key string, val interface{}) {
	c.cache[key] = val
}
func (c *dbCache) Reset() {
	c.cache = make(map[string]interface{})
}
func (c *dbCache) Prune() {
	// todo: for LRU cache this need implementing
}
