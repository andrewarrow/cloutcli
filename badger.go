package cloutcli

import (
	"fmt"
	"strings"

	"github.com/andrewarrow/cloutcli/database"
	"github.com/andrewarrow/cloutcli/lib"
	"github.com/btcsuite/btcutil/base58"
	"github.com/dgraph-io/badger/v3"
)

// magic numbers are defined in
// https://github.com/bitclout/core/blob/main/lib/db_utils.go

var Testing bool
var Tables string

func ImportFromBadgerToSqlite(dir string) error {
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		return err
	}
	defer db.Close()
	entryChan := make(chan database.EntryHolder, 1024)
	go database.EnumerateAll(Testing, db, &entryChan)

	sdb := database.OpenSqliteDB()
	database.CreateSchema(sdb)
	defer sdb.Close()

	skipList := map[string]bool{}
	for _, item := range strings.Split(Tables, ",") {
		skipList[item] = true
	}

	i := 0
	for entry := range entryChan {
		if entry.Flavor == "post" && !skipList["post"] {
			database.InsertPostSqlite(sdb, entry.Thing.(*lib.PostEntry))
		} else if entry.Flavor == "profile" && !skipList["profile"] {
			database.InsertProfileSqlite(sdb, entry.Thing.(*lib.ProfileEntry))
		} else if entry.Flavor == "follow" && !skipList["follow"] {
			database.InsertFollowee(sdb, base58.Encode(entry.Followed),
				base58.Encode(entry.Follower))
		} else if entry.Flavor == "done" {
			break
		}
		i++
		if i%1000 == 0 {
			fmt.Println("iteration", i, entry.Flavor)
		}
	}
	return nil
}

func PrintAllPostsFromBadger(dir string) error {
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		return err
	}
	defer db.Close()
	postEntryChan := make(chan *lib.PostEntry, 1024)
	go database.EnumeratePosts(db, &postEntryChan)
	for postEntry := range postEntryChan {
		fmt.Println(string(postEntry.Body))
	}
	return nil
}
