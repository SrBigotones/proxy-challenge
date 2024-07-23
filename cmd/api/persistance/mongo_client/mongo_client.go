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

type MongoClient struct {
	collection *mongo.Collection
	ctx        context.Context

	Addr           string
	Port           string
	Db             string
	CollectionName string
}

func NewMognoClient(addr string, port string, dbName string, collectionName string) *MongoClient {
	mongoClient := MongoClient{
		Addr:           addr,
		Port:           port,
		Db:             dbName,
		CollectionName: collectionName,
		ctx:            context.TODO(),
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + mongoClient.Addr + ":" + mongoClient.Port + "/")
	client, err := mongo.Connect(mongoClient.ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(mongoClient.ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient.collection = client.Database("proxy").Collection("client_stats")

	return &mongoClient
}

func (mongoSession *MongoClient) InsertToCollection(jsonString []byte) {
	var v interface{}

	if err := json.Unmarshal(jsonString, &v); err != nil {
		// handle error
		log.Fatal(err)
	}
	mongoSession.collection.InsertOne(mongoSession.ctx, v)
}

func (mongoSession *MongoClient) FindByIP(ip string) ([]*user_stats.UserStats, error) {

	exampleFilter := bson.D{primitive.E{Key: "ip", Value: ip}}
	cur, err := mongoSession.collection.Find(mongoSession.ctx, exampleFilter)

	var results []*user_stats.UserStats

	if err != nil {
		return results, err
	}

	for cur.Next(mongoSession.ctx) {
		var client user_stats.UserStats
		err := cur.Decode(&client)
		if err != nil {
			return results, err
		}

		results = append(results, &client)
	}

	cur.Close(mongoSession.ctx)

	return results, nil
}

func (mongoSession *MongoClient) FindAll() ([]*user_stats.UserStats, error) {

	// filter := bson.D{{}}

	findOptions := options.Find()

	cur, err := mongoSession.collection.Find(mongoSession.ctx, bson.D{{}}, findOptions)

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

	for cur.Next(mongoSession.ctx) {
		var client user_stats.UserStats
		err := cur.Decode(&client)
		if err != nil {
			return results, err
		}

		results = append(results, &client)
	}

	cur.Close(mongoSession.ctx)

	return results, nil
}
