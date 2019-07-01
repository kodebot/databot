package cache

// todo: think about concurrency
type memCache struct {
	cache map[string]interface{}
}

// NewMemCache returns new in memory cache
func NewMemCache() Manager {
	return &memCache{cache: make(map[string]interface{})}
}

func (c *memCache) Get(key string) interface{} {
	return c.cache[key]
}

func (c *memCache) Add(key string, val interface{}) {
	c.cache[key] = val
}
func (c *memCache) Reset() {
	c.cache = make(map[string]interface{})
}
func (c *memCache) Prune() {
	// noop
	// todo: for LRU cache this need implementing
}
