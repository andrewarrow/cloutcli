package main

import (
	"clout/files"
	"fmt"
	"os"
	"strings"

	"github.com/andrewarrow/cloutcli"
	"github.com/andrewarrow/cloutcli/database"
	"github.com/dgraph-io/badger/v3"
)

func PrintSqliteHelp() {
	fmt.Println("")
	fmt.Println("  clout sqlite fill           # --dir=/path/to/badgerdb")
	fmt.Println("  clout sqlite graph          # produce clout.gv file")
	fmt.Println("  clout sqlite query          # --term=foo")
	fmt.Println("  clout sqlite likes          # --username=foo")
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

		if argMap["stats"] != "" {

			db, _ := badger.Open(badger.DefaultOptions(dir))
			defer db.Close()
			sdb := database.OpenSqliteDB()
			database.CreateSchema(sdb)
			defer sdb.Close()
			database.PostsByAuthor(sdb, db, argMap["username"])
			return
		}

		if argMap["testing"] != "" {
			cloutcli.Testing = true
		}
		cloutcli.Tables = argMap["tables"]
		cloutcli.ImportFromBadgerToSqlite(dir)
	} else if command == "graph" {
		ProduceCloutGV()
	} else if command == "likes" {
		SortUsersByNumberLikesToYourPosts()
	} else if command == "query" {
		term := argMap["term"]
		table := argMap["table"]

		degrees := argMap["degrees"]
		if degrees == "" {
			degrees = "2"
		}

		if table == "" || table == "posts" || table == "post" {
			cloutcli.QuerySqlitePosts(term)
		} else if table == "users" || table == "user" {
			cloutcli.QuerySqliteUsers(term)
		} else if table == "follow" {
			cloutcli.QuerySqliteFollow("", term, degrees)
		}
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

func ProduceCloutGV() {
	fmt.Println("creating clout.gv...")
	f, _ := os.Create("clout.gv")
	f.Write([]byte("digraph regexp {\n"))
	fmt.Println("writing nodes...")
	cloutcli.QuerySqliteNodesInOrder(f)
	fmt.Println("writing connections...")
	cloutcli.QuerySqliteNodeConnections(f)
	f.Write([]byte("}\n"))
	fmt.Println("done")
	f.Close()
}

func SortUsersByNumberLikesToYourPosts() {
	username := argMap["username"]
	if username == "" {
		fmt.Println("run with --username=foo")
		return
	}
	cloutcli.QuerySqliteLikesForAuthor(username)
}
