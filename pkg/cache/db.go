package cache

import (
	"errors"

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

	adapter := newDBAdapter(dbcache.Mongo)
	adapter.Connect("mongodb://localhost:27017")

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
