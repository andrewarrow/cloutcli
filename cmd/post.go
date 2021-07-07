package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/andrewarrow/cloutcli"
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
		HandleOnePost(payload)
	}
}

func HandleOnePost(payload string) {
	// https://bitclout.com/posts/57d0a7d5640bdc2676df0e40f16f70393c4c484731a7867974c9c27789662ad8
	hash := payload
	if strings.HasPrefix(payload, "https") {
		hash = payload[27:]
	}
	fmt.Println(hash)
	p := cloutcli.SinglePost(hash)
	fmt.Printf("%+v\n", p)
}
