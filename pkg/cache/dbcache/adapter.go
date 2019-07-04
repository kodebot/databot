package dbcache

// Adapter acts as a mediator between database based cache manager and the database
type Adapter interface {
	Connect(conStr string)
	Disconnect()
	Get(key string) interface{}
	Add(key string, val interface{})
	Reset()
	Prune()
}

// DBType is the type that represents the identifiers of the database
type DBType string

const (
	// Mongo DB type
	Mongo DBType = "mongo"
)
