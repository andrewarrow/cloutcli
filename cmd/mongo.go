package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PrintMongoHelp() {
	fmt.Println("")
	fmt.Println("  clout mongo ls           # list what is there")
	fmt.Println("")
}
func HandleMongo() {
	if len(os.Args) < 3 {
		PrintMongoHelp()
		return
	}
	command := os.Args[2]
	if command == "ls" {
		MongoList()
	}
}

func MongoConnect() *mongo.Client {
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb://192.168.1.50:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed establishing a connection with MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB:  %v", err)
	}

	return client

}

type Thing struct {
	//MongoMeta string
	PublicKey     string
	LikedPostHash interface{}
}

type Profile struct {
	PublicKey string
	Username  string
}

func MongoList() {
	client := MongoConnect()
	client2 := MongoConnect()
	collection := client.Database("bitclout").Collection("data")
	collection2 := client2.Database("bitclout").Collection("data")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	/*
		var result Thing
		filter := bson.D{{"BadgerKeyPrefix", "_PrefixLikedPostHashToLikerPubKey:31"}}
		collection.FindOne(ctx, filter).Decode(&result)
		fmt.Println(result.PublicKey)

		var profile Profile
		//filter = bson.D{{"BadgerKeyPrefix", "_PrefixProfilePubKeyToProfileEntry:23"}}
		filter = bson.D{{"BadgerKeyPrefix", "_PrefixProfilePubKeyToProfileEntry:23"}, {"PublicKey", result.PublicKey}}
		collection.FindOne(ctx, filter).Decode(&profile)
		fmt.Println(profile)
	*/

	//_PrefixProfilePubKeyToProfileEntry:23
	cur, _ := collection.Find(ctx, bson.D{{"BadgerKeyPrefix", "_PrefixLikedPostHashToLikerPubKey:31"}})
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Thing
		cur.Decode(&result)
		fmt.Printf("%s\n", result.PublicKey)
		filter := bson.D{{"BadgerKeyPrefix", "_PrefixProfilePubKeyToProfileEntry:23"}, {"PublicKey", result.PublicKey}}

		var profile Profile
		ctx2, _ := context.WithTimeout(context.Background(), 1*time.Second)
		collection2.FindOne(ctx2, filter).Decode(&profile)
		fmt.Println(profile)
	}
}
