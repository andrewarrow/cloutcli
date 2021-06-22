package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clout account               # list your various accounts")
	fmt.Println("  clout message               # send, send bulk, read")
	fmt.Println("  clout ls                    # list global posts")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "account" || command == "accounts" {
	} else if command == "message" || command == "messages" {
		HandleMessage()
	} else if command == "ls" {
		HandleLs()
	}

}
