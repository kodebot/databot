package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/kodebot/databot/pkg/cache/dbcache"
	"github.com/kodebot/databot/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

type mongoDBAdapter struct{}

type document struct {
	Key        string
	Val        interface{}
	LastWrite  *time.Time
	WriteCount int
	LastRead   *time.Time
	ReadCount  int
}

// NewAdapter returns mongo DB cache adapter
func NewAdapter() dbcache.Adapter {
	return &mongoDBAdapter{}
}

// Connect establishes connection to mongodb
func (a *mongoDBAdapter) Connect(conStr string) {

	if dbClient != nil {
		return // connection is already created
	}

	var err error
	dbClient, err = mongo.NewClient(options.Client().ApplyURI(conStr))
	if err != nil {
		logger.Fatalf("error when creating new mongo client %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = dbClient.Connect(ctx)
	if err != nil {
		logger.Fatalf("error when connecting to mongo database %s", err.Error())
	}
}

// Disconnect closes the connection with mongodb
func (a *mongoDBAdapter) Disconnect() {
	err := dbClient.Disconnect(context.TODO())
	if err != nil {
		logger.Errorf("error when disconneting mongodb connection. error: %s", err.Error())
	}

	dbClient = nil
}

// Get returns an item matching the key from the cache
func (a *mongoDBAdapter) Get(key string) interface{} {
	var doc document
	err := findOne(bson.M{"key": key}, &doc)
	if err != nil {
		logger.Tracef("error when reading from mongodb backed cache. error: %s", err.Error())
		return nil
	}
	// todo: this hits the database twice, find a way to improve this situation
	now := time.Now().UTC()
	doc.LastRead = &now
	doc.ReadCount++
	_, err = update(&doc)
	if err != nil {
		logger.Errorf("error when updating cache read audit data")
	}
	return doc.Val
}

// Add inserts new item to the cache
func (a *mongoDBAdapter) Add(key string, val interface{}) {
	now := time.Now().UTC()
	_, err := insertOne(&document{key, val, &now, 1, nil, 0})
	if err != nil {
		logger.Errorf("adding item to mongodb backed cache failed with error: %s", err.Error())
	}
}

// Reset clears all the items from the cache
func (a *mongoDBAdapter) Reset() {
	panic(errors.New("not implemented"))
}

// Prune removes rarely used items from the cache
func (a *mongoDBAdapter) Prune() {
	// todo: for LRU cache this need implementing
	panic(errors.New("not implemented"))
}

func getCacheCollection() *mongo.Collection {
	// todo: take the database name from the config
	collection := dbClient.Database("databot").Collection("cache")
	return collection
}

func findOne(filter interface{}, result *document) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return getCacheCollection().FindOne(ctx, filter).Decode(result)
}

func insertOne(doc *document) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return getCacheCollection().InsertOne(ctx, doc)
}

func update(doc *document) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	update := bson.M{
		"$set": bson.M{
			"lastread":  doc.LastRead,
			"readcount": doc.ReadCount,
		},
	}
	return getCacheCollection().UpdateOne(ctx, bson.M{"key": doc.Key}, update)
}

func delete(filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return getCacheCollection().DeleteMany(ctx, filter)
}
