package main

import (
	"clout/files"
	"fmt"
	"os"
	"strings"

	"github.com/andrewarrow/cloutcli"
)

func PrintSqliteHelp() {
	fmt.Println("")
	fmt.Println("  clout sqlite fill           # --dir=/path/to/badgerdb")
	fmt.Println("  clout sqlite query          # --term=foo")
	fmt.Println("")
}
func HandleSqlite() {
	if len(os.Args) < 3 {
		PrintSqliteHelp()
		return
	}
	command := os.Args[2]
	if command == "fill" {
		dir := DirCheck()
		cloutcli.ImportFromBadgerToSqlite(dir)
	} else if command == "query" {
		cloutcli.QuerySqlitePosts(argMap["term"])
	}
}

func DirCheck() string {
	dir := argMap["dir"]
	if dir == "" {
		fmt.Println("run with --dir=/home/name/path/to/badgerdb")
		return ""
	}
	if strings.HasPrefix(dir, "~") {
		home := files.UserHomeDir()
		return home + dir[1:]
	}
	return dir
}
