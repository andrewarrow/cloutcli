package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	//	"github.com/andrewarrow/cloutcli"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clout accounts               # list your various accounts")
	fmt.Println("  clout ls                     # list global posts")
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
	} else if command == "ls" {
		HandleLs()
	}

}
