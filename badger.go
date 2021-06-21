package cloutcli

import (
	"fmt"

	"github.com/andrewarrow/cloutcli/database"
	"github.com/andrewarrow/cloutcli/lib"
	"github.com/dgraph-io/badger/v3"
)

// magic numbers are defined in
// https://github.com/bitclout/core/blob/main/lib/db_utils.go

func ImportFromBadgerToSqlite(dir string) error {
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		return err
	}
	defer db.Close()
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
