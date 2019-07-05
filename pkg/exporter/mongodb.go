package exporter

import (
	"context"
	"time"

	"github.com/kodebot/databot/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ExportToMongoDB exports the records into mongo DB
func ExportToMongoDB(recs []map[string]interface{}, conStr string) {

	dbClient, err := mongo.NewClient(options.Client().ApplyURI(conStr))
	if err != nil {
		logger.Fatalf("error when creating new mongo client %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = dbClient.Connect(ctx)
	defer func() { dbClient.Disconnect(context.TODO()) }()
	if err != nil {
		logger.Fatalf("error when connecting to mongo database %s", err.Error())
	}
	// todo: take the database name from the config
	collection := dbClient.Database("databot").Collection("export")

	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	for _, rec := range recs {
		collection.InsertOne(ctx, rec)
	}
}
