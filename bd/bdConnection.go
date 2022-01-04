package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is the connection object
var MongoCN = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://chrisloarryn:ltbGMTpNQuugHBH3@cluster0.wvkde.mongodb.net/test?authSource=admin&replicaSet=atlas-ngcg07-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

// ConnectBD is the method that returns the client
func ConnectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connected to MongoDB")

	return client
}

// CheckConnection is the pint to the db
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
