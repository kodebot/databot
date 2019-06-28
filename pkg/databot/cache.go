package databot

// CacheService is abstract cache lookup service
type CacheService interface {
	Get(key interface{}) interface{}
	Add(key interface{}, val interface{})
	Reset()
	Prune()
}
