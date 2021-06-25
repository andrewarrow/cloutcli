package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

func MongoList() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed establishing a connection with MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB:  %v", err)
	}

	fmt.Println("Successfully Connected to MongoDB.")
}
