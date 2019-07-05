package config

const (
	defaultUseDBCache  = true
	defaultCacheDBType = "mongo"
	defaultCacheConStr = "mongodb://localhost:27017"
)

// Options is the container for all configuration options
type Options struct {
	useDBCache    bool
	cacheDBType   string
	cacheDBConStr string
}

// UseDBCache returns true when database cache should be used; false for in memory cache
func (o *Options) UseDBCache() bool {
	return o.useDBCache
}

// CacheDBType returns the type of the database that should be used for storing cached data
func (o *Options) CacheDBType() string {
	return o.cacheDBType
}

// CacheDBConStr returns connection string to access the cache database
func (o *Options) CacheDBConStr() string {
	return o.cacheDBConStr
}
