package main

import (
	"fmt"
	"os"
)

func PrintMessageHelp() {
	fmt.Println("")
	fmt.Println("  clout message bulk           # --to=allfollowers [--text=foo]")
	fmt.Println("  clout message inbox          # --filter=myhodlers")
	fmt.Println("  clout message new            # --to=username [--text=foo]")
	fmt.Println("  clout message reply          # --id=foo [--text=foo]")
	fmt.Println("  clout message show           # --id=foo")
	fmt.Println("")
}
func HandleMessage() {
	if len(os.Args) < 3 {
		PrintMessageHelp()
		return
	}
	command := os.Args[2]
	if command == "inbox" {
		MessageInbox()
	}
}

func MessageInbox() {
	username := os.Getenv("CLOUTCLI_USERNAME")
	if username == "" {
		fmt.Println("set CLOUTCLI_USERNAME")
		return
	}
}
