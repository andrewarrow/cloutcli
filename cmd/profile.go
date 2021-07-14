package main

import (
	"fmt"
	"os"

	"github.com/andrewarrow/cloutcli"
)

func PrintProfileHelp() {
	fmt.Println("")
	fmt.Println("  clout profile lookup [payload]      # pub58 string")
	fmt.Println("")
}
func HandleProfiles() {
	if len(os.Args) < 3 {
		PrintProfileHelp()
		return
	}
	command := os.Args[2]
	if command == "lookup" {
		payload := os.Args[3]
		LookupProfile(payload)
	} else if command == "" {
	}
}

func LookupProfile(pub58 string) {
	user := cloutcli.Pub58ToUser(pub58)
	username := user.ProfileEntryResponse.Username
	fmt.Println("")
	fmt.Println(username)
	fmt.Println("")
}
