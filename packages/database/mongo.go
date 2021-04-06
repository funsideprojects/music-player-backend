package database

import (
	"context"
	"log"
	"time"

	"fsp/open-music/packages/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Client {
	mongoPort := env.GetEnv("MONGO_PORT")

	// Init mongo client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:" + mongoPort))
	if err != nil {
		log.Fatal(err)
	}

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Establish connection
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// List all databases
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	return client
}
