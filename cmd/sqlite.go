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
			usernames := argMap["usernames"]
			tokens := strings.Split(usernames, ",")
			pub58s := []string{}
			for _, username := range tokens {
				pub58 := cloutcli.UsernameToPub58(username)
				pub58s = append(pub58s, pub58)
			}

			database.PostsByAuthor(db, tokens, pub58s)
			return
		}

		if argMap["testing"] != "" {
			cloutcli.Testing = true
		}
		cloutcli.Tables = argMap["tables"]
		cloutcli.ImportFromBadgerToSqlite(dir)
	} else if command == "graph" {
		ProduceCloutGV()
	} else if command == "query" {
		HandleSimpleQueries()
	}
}

func HandleSimpleQueries() {
	if CheckForCloutDbFirst() == false {
		return
	}
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

func DirCheck() string {
	dir := argMap["dir"]
	if dir == "" {
		fmt.Println("run with --dir=/path/to/badgerdb")
		return ""
	}
	if strings.HasPrefix(dir, "~") {
		home := files.UserHomeDir()
		return home + dir[1:]
	}
	return dir
}

func CheckForCloutDbFirst() bool {
	_, err := os.Stat("clout.db")
	if err != nil {
		fmt.Println("run 'clout demo sqlite' first to create your sqlite db.")
		return false
	}
	return true
}
func ProduceCloutGV() {
	if CheckForCloutDbFirst() == false {
		return
	}

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
