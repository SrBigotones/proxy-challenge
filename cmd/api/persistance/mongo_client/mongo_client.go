package mongo_client

import (
	"context"
	"encoding/json"
	"log"

	"github.com/SrBigotones/proxy-challenge/cmd/api/model/user_stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func FindByIP(ip string) ([]*user_stats.UserStats, error) {

	exampleFilter := bson.D{primitive.E{Key: "ip", Value: ip}}
	cur, err := collection.Find(ctxMongo, exampleFilter)

	var results []*user_stats.UserStats

	if err != nil {
		return results, err
	}

	for cur.Next(ctxMongo) {
		var client user_stats.UserStats
		err := cur.Decode(&client)
		if err != nil {
			return results, err
		}

		results = append(results, &client)
	}

	cur.Close(ctxMongo)

	return results, nil
}

func FindAll() ([]*user_stats.UserStats, error) {

	// filter := bson.D{{}}

	findOptions := options.Find()

	cur, err := collection.Find(ctxMongo, bson.D{{}}, findOptions)

	var results []*user_stats.UserStats

	if err != nil {
		return results, err
	}

	// var foo []*user_stats.UserStats
	// if err = cur.All(ctxMongo, &foo); err != nil {
	// 	panic(err)
	// }

	// cur.Close(ctxMongo)
	// return foo, err

	for cur.Next(ctxMongo) {
		var client user_stats.UserStats
		err := cur.Decode(&client)
		if err != nil {
			return results, err
		}

		results = append(results, &client)
	}

	cur.Close(ctxMongo)

	return results, nil
}
