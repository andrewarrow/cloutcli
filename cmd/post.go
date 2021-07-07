package main

import (
	"fmt"
	"os"
)

func PrintPostHelp() {
	fmt.Println("")
	fmt.Println("  clout post inspect [payload]      # full url or hex hash")
	fmt.Println("")
}
func HandlePosts() {
	if len(os.Args) < 3 {
		PrintPostHelp()
		return
	}
	command := os.Args[2]
	if command == "inspect" {
		payload := os.Args[3]
		fmt.Println(payload)
	}
}

func HandleOnePost() {
}
