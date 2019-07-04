package cache

// Manager is abstract cache lookup service
type Manager interface {
	Get(key string) interface{}
	Add(key string, val interface{})
	Reset()
	Prune()
}

var currentCache Manager

func init() {
	// todo: change the current cache type via config
	currentCache = NewMemCache()
}

// Current returns the currently configured cache
func Current() Manager {
	return currentCache
}
