package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/kodebot/databot/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

// Connect establishes connection to mongodb
func Connect(uri string) {

	if dbClient != nil {
		logger.Fatalf("connection is already created. only one connection to just one database is supported")
	}

	var err error
	dbClient, err = mongo.NewClient(options.Client().ApplyURI(uri))
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
func Disconnect() {
	err := dbClient.Disconnect(context.TODO())
	if err != nil {
		logger.Errorf("error when disconneting mongodb connection. error: %s", err.Error())
	}

	dbClient = nil
}

type document struct {
	key string
	val interface{}
}

// Get returns an item matching the key from the cache
func Get(key string) interface{} {
	var result document
	err := findOne(bson.M{"key": key}, &result)
	if err != nil {
		logger.Errorf("error when reading from mongodb backed cache. error: %s", err.Error())
		return nil
	}
	return result
}

// Add inserts new item to the cache
func Add(key string, val interface{}) {
	_, err := insertOne(document{key, val})
	if err != nil {
		logger.Errorf("adding item to mongodb backed cache failed with error: %s", err.Error())
	}
}

// Reset clears all the items from the cache
func Reset() {
	panic(errors.New("not implemented"))
}

// Prune removes rarely used items from the cache
func Prune() {
	// todo: for LRU cache this need implementing
	panic(errors.New("not implemented"))
}

func getCacheCollection() *mongo.Collection {
	// todo: take the database name from the config
	collection := dbClient.Database("databot").Collection("cache")
	return collection
}

func findOne(filter interface{}, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return getCacheCollection().FindOne(ctx, filter).Decode(result)
}

func insertOne(document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return getCacheCollection().InsertOne(ctx, &document)
}

func delete(filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return getCacheCollection().DeleteMany(ctx, filter)
}
