package cache

import (
	"errors"
	"fmt"

	"github.com/kodebot/databot/pkg/config"

	"github.com/kodebot/databot/pkg/cache/dbcache"
	"github.com/kodebot/databot/pkg/cache/mongodb"
	"github.com/kodebot/databot/pkg/logger"
)

// todo: update this to use MongoDB for persistence
// todo: think about concurrency

type dbCache struct {
	adapter dbcache.Adapter
}

var current *dbCache

// NewDBCache returns new in memory cache
func NewDBCache() Manager {
	if current != nil {
		logger.Fatalf("multiple cache are not supported")
	}

	var adapter dbcache.Adapter

	cacheDBType := config.Current().CacheDBType()
	switch cacheDBType {
	case string(dbcache.Mongo):
		adapter = newDBAdapter(dbcache.Mongo)
	default:
		panic(fmt.Errorf("cache database type %s is not supported", cacheDBType))
	}

	adapter.Connect(config.Current().CacheDBConStr())

	current = &dbCache{adapter: adapter}
	return current
}

func (c *dbCache) Get(key string) interface{} {
	return c.adapter.Get(key)
}

func (c *dbCache) Add(key string, val interface{}) {
	c.adapter.Add(key, val)
}

func (c *dbCache) Reset() {
	c.adapter.Reset()
}

func (c *dbCache) Prune() {
	// todo: for LRU cache this need implementing
	c.adapter.Prune()
}

func newDBAdapter(dbType dbcache.DBType) dbcache.Adapter {
	switch dbType {
	case dbcache.Mongo:
		return mongodb.NewAdapter()
	default:
		panic(errors.New("not supported database type"))
	}
}
