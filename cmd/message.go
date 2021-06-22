package main

import "fmt"

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
	PrintMessageHelp()
}
