package main

import (
	"clout/args"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/cloutcli"
	"github.com/btcsuite/btcutil/base58"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clout account               # list your various accounts")
	fmt.Println("  clout demo                  # demo some core concepts")
	fmt.Println("  clout ls                    # list global posts")
	fmt.Println("  clout post                  # things you can do with posts")
	fmt.Println("  clout profile               # things you can do with profiles")
	fmt.Println("  clout message               # send, send bulk, read")
	fmt.Println("  clout sell                  # sell coins")
	fmt.Println("  clout sqlite                # import from badger, query sqlite")
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
	} else if command == "demo" {
		HandleDemo()
	} else if command == "ls" {
		HandleLs()
	} else if command == "message" || command == "messages" {
		HandleMessage()
	} else if command == "post" || command == "posts" {
		HandlePosts()
	} else if command == "profile" || command == "profiles" {
		HandleProfiles()
	} else if command == "sell" {
		HandleSell()
	} else if command == "sqlite" {
		HandleSqlite()
	} else if command == "username" {
		pub58 := cloutcli.UsernameToPub58(argMap["username"])
		fmt.Println(pub58)
		decoded := base58.Decode(pub58)
		fmt.Println(decoded[3 : len(decoded)-4])
	}

}
