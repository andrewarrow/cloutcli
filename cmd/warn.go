package main

import (
	"fmt"
	"os"
)

func WarnAboutWords() string {

	words := os.Getenv("CLOUTCLI_SEED_WORDS")
	if words == "" {
		fmt.Println("")
		fmt.Println("cloutcli needs your private key i.e seed words")
		fmt.Println("temporarily set an ENVIRONEMENT_VARIABLE named \"CLOUTCLI_SEED_WORDS\"")
		fmt.Println("")
		fmt.Println("don't just type export CLOUTCLI_SEED_WORDS=words from a bash prompt")
		fmt.Println("that would go into your bash history")
		fmt.Println("instead edit your .bash_profile or equivalent and place the")
		fmt.Println("export command there and save the file.")
		fmt.Println("")

		fmt.Println("open a new terminal - this is a \"hot terminal window\"")
		fmt.Println("think of it like taking the safety off a gun, at some point you want to")
		fmt.Println("put the safety back on. And you can do this by editing that .bash_profile")
		fmt.Println("file again and removing the line you added.")
		fmt.Println("")
		fmt.Println("close the hot terminal window(s) and take a deep breath.")
		fmt.Println("but while your terminal is hot, run the command you need to run")
		fmt.Println("")
		return ""
	}
	return words
}
