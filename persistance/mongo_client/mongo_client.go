package mongo_client

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo related suff here
var collection *mongo.Collection
var ctxMongo = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctxMongo, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctxMongo, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("proxy").Collection("client_stats")
}

func InsertToCollection(jsonString []byte) {
	var v interface{}

	if err := json.Unmarshal(jsonString, &v); err != nil {
		// handle error
		log.Fatal(err)
	}
	collection.InsertOne(ctxMongo, v)
}
