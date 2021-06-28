package main

import (
	"fmt"
	"os"

	"github.com/andrewarrow/cloutcli"
)

func PrintDemoHelp() {
	fmt.Println("")
	fmt.Println("  clout demo visualizegraph  # make clout.gv graph file")
	fmt.Println("  clout demo printall        # print all clouts")
	fmt.Println("  clout demo search          # search sqlite database")
	fmt.Println("  clout demo sqlite          # place data into local sqlite database")
	fmt.Println("")
	fmt.Println("  search examples:")
	fmt.Println("")
	fmt.Println("  ./clout demo search --term=hi --table=users")
	fmt.Println("  ./clout demo search --term=hi --table=posts")
	fmt.Println("  ./clout demo search --term=username --table=follow --degrees=2")
	fmt.Println("")
}
func HandleDemo() {
	if len(os.Args) < 3 {
		PrintDemoHelp()
		return
	}
	command := os.Args[2]
	if command == "visualizegraph" {
		ProduceCloutGV()
	} else if command == "printall" {
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
