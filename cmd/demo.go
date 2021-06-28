package main

import (
	"fmt"
	"os"

	"github.com/andrewarrow/cloutcli"
)

func PrintDemoHelp() {
	fmt.Println("")
	fmt.Println("  clout demo graph       # make clout.gv graph file")
	fmt.Println("  clout demo posts       # print all clouts")
	fmt.Println("  clout demo search      # search sqlite database")
	fmt.Println("  clout demo sqlite      # place clouts into local sqlite database")
	fmt.Println("")
}
func HandleDemo() {
	if len(os.Args) < 3 {
		PrintDemoHelp()
		return
	}
	command := os.Args[2]
	if command == "graph" {
		ProduceCloutGV()
	} else if command == "posts" {
		dir := DirCheck()
		if dir == "" {
			return
		}
		cloutcli.PrintAllPostsFromBadger(dir)
	} else if command == "search" {
		HandleSimpleQueries()
	} else if command == "sqlite" {
		dir := DirCheck()
		if dir == "" {
			return
		}
		cloutcli.Tables = "post,profile,follow"
		cloutcli.ImportFromBadgerToSqlite(dir)
	}
}
