package main

import (
	"fmt"
	"os"
	"time"

	"github.com/andrewarrow/cloutcli"
	"github.com/justincampbell/timeago"
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

	list := cloutcli.MessageInbox(username)
	for _, oc := range list.OrderedContactsWithMessages {
		from := list.PublicKeyToProfileEntry[oc.PublicKeyBase58Check].Username
		fmt.Println("  ", from)
		for _, m := range oc.Messages {
			ts := time.Unix(m.TstampNanos/1000000000, 0)
			ago := timeago.FromDuration(time.Since(ts))
			fmt.Println("    ", ago)
		}
	}
}
