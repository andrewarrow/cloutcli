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

	okList := map[string]bool{"diamond": true,
		"like":    true,
		"reclout": true,
		"follow":  true,
		"post":    true,
		"profile": true}
	if Tables != "" {
		okList = map[string]bool{}
		for _, item := range strings.Split(Tables, ",") {
			okList[item] = true
		}
	}
	okList["done"] = true
	fmt.Println(okList)

	i := 0
	for entry := range entryChan {
		i++
		if i%1000 == 0 {
			fmt.Println("iteration", i, entry.Flavor)
		}
		if okList[entry.Flavor] == false {
			continue
		}
		if entry.Flavor == "diamond" {
			database.InsertDiamondSqlite(sdb, entry.Thing.(*lib.DiamondEntry))
		} else if entry.Flavor == "follow" {
			database.InsertFollowee(sdb, base58.Encode(entry.Followed),
				base58.Encode(entry.Follower))
		} else if entry.Flavor == "like" {
			database.InsertLikeSqlite(sdb, entry.Thing.(*lib.LikeEntry))
		} else if entry.Flavor == "post" {
			database.InsertPostSqlite(sdb, entry.Thing.(*lib.PostEntry))
		} else if entry.Flavor == "profile" {
			database.InsertProfileSqlite(sdb, entry.Thing.(*lib.ProfileEntry))
		} else if entry.Flavor == "reclout" {
			database.InsertRecloutSqlite(sdb, entry.Thing.(*lib.RecloutEntry))
		} else if entry.Flavor == "done" {
			break
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
