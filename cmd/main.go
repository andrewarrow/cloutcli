package main

import (
	"clout/args"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clout account               # list your various accounts")
	fmt.Println("  clout message               # send, send bulk, read")
	fmt.Println("  clout sell                  # sell coins")
	fmt.Println("  clout sqlite                # import from badger, query sqlite")
	fmt.Println("  clout ls                    # list global posts")
	fmt.Println("")
}

var argMap map[string]string

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]
	argMap = args.ToMap()

	if command == "account" || command == "accounts" {
	} else if command == "message" || command == "messages" {
		HandleMessage()
	} else if command == "ls" {
		HandleLs()
	} else if command == "sell" {
		HandleSell()
	} else if command == "sqlite" {
		HandleSqlite()
	}

}
