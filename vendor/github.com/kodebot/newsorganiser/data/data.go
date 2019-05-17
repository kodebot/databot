package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetCollection returns the provided collection
func GetCollection(collectionName string) (*mongo.Collection, error) {
	// todo: make dataaccess as reusable
	// todo: take connection string from config
	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = dbClient.Connect(ctx)
	if err != nil {
		return nil, err
	}
	// todo: take the database name from the config
	collection := dbClient.Database("newsorganiser").Collection(collectionName)
	return collection, nil
}

// FindOne gets the first document matches the filter in the collection
func FindOne(collection *mongo.Collection, filter interface{}, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return collection.FindOne(ctx, filter).Decode(result)
}

// Find gets the first document matches the filter in the collection
func Find(collection *mongo.Collection, filter interface{}) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var result []interface{}

	for cursor.Next(ctx) {
		result = append(result, cursor.Current)
	}

	return result, nil
}

// InsertOne inserts one document to the given collection
func InsertOne(collection *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return collection.InsertOne(ctx, &document)
}
