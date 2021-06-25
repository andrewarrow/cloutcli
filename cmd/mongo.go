package main

import (
	"fmt"
	"os"
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
}
