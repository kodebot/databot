package data

import (
	"context"
	"time"

	"github.com/golang/glog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func init() {
	println("initialising data package")
	// todo: make dataaccess as reusable
	// todo: take connection string from config
	var err error
	dbClient, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		glog.Fatalf("error when creating new mongo client %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = dbClient.Connect(ctx)
	if err != nil {
		glog.Fatalf("error when connecting to mongo database %s", err.Error())
	}
}

// FindResultDecoder type
type FindResultDecoder func(*mongo.Cursor) interface{}

// GetCollection returns the provided collection
func GetCollection(collectionName string) (*mongo.Collection, error) {
	// todo: take the database name from the config
	collection := dbClient.Database("newsfeed").Collection(collectionName)
	return collection, nil
}

// FindOne gets the first document matches the filter in the collection
func FindOne(collection *mongo.Collection, filter interface{}, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return collection.FindOne(ctx, filter).Decode(result)
}

// Find gets the first document matches the filter in the collection
func Find(collection *mongo.Collection, filter interface{}, decode FindResultDecoder, opts ...*options.FindOptions) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, filter, opts...)
	if err != nil {
		glog.Errorf("find one failed with error %s", err.Error())
		return nil, err
	}

	var result []interface{}

	for cursor.Next(ctx) {
		result = append(result, decode(cursor))
	}
	return result, nil
}

// InsertOne inserts one document to the given collection
func InsertOne(collection *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return collection.InsertOne(ctx, &document)
}

// Delete deletes the documents matching the filter
func Delete(collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return collection.DeleteMany(ctx, filter)
}
