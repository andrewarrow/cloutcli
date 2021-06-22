package main

import (
	"clout/files"
	"fmt"
	"os"
	"time"

	"github.com/andrewarrow/cloutcli"
	"github.com/andrewarrow/cloutcli/keys"
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
	} else if command == "new" {
		MessageNew()
	} else if command == "bulk" {
		MessageBulk()
	}
}

func MessageBulk() {
	words := os.Getenv("CLOUTCLI_SEED_WORDS")
	if words == "" {
		fmt.Println("set CLOUTCLI_SEED_WORDS")
		return
	}
	pub58, _ := keys.ComputeKeysFromSeed(words)
	to := argMap["to"]
	text := argMap["text"]

	if text == "" {
		text = files.ReadFromIn()
	}

	bulkList := []string{}
	if to == "allfollowers" {
		me := cloutcli.Pub58ToUser(pub58)
		items := cloutcli.LoopThruAllFollowing(pub58, me.ProfileEntryResponse.Username, 0)
		for _, item := range items {
			bulkList = append(bulkList, item.Username)
		}
	}

	for _, username := range bulkList {
		if argMap["dryrun"] != "" {
			fmt.Println(username)
			continue
		}
		argMap["to"] = username
		argMap["text"] = text
		MessageNew()
	}
}

func MessageNew() {
	words := os.Getenv("CLOUTCLI_SEED_WORDS")
	if words == "" {
		fmt.Println("set CLOUTCLI_SEED_WORDS")
		return
	}
	to := argMap["to"]
	text := argMap["text"]

	if text == "" {
		text = files.ReadFromIn()
	}

	ok := cloutcli.SendMessage(words, to, text)
	fmt.Println(to, ok)
}

func MessageInbox() {
	username := os.Getenv("CLOUTCLI_USERNAME")
	if username == "" {
		fmt.Println("set CLOUTCLI_USERNAME")
		return
	}

	list := cloutcli.MessageInbox(username)
	for j, oc := range list.OrderedContactsWithMessages {
		from := list.PublicKeyToProfileEntry[oc.PublicKeyBase58Check].Username
		fmt.Println("  ", from)
		for i, m := range oc.Messages {
			ts := time.Unix(m.TstampNanos/1000000000, 0)
			ago := timeago.FromDuration(time.Since(ts))
			fmt.Println("    ", ago)
			if i > 3 {
				break
			}
		}
		if j > 3 {
			break
		}
	}
}
