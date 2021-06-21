package cloutcli

import (
	"github.com/andrewarrow/cloutcli/database"
	"github.com/dgraph-io/badger/v3"
)

func ImportFromBadgerToSqlite(dir string) error {
	return nil
}

func PrintAllPostsFromBadger(dir string) error {
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		return err
	}
	defer db.Close()
	// magic numbers are defined in
	// https://github.com/bitclout/core/blob/main/lib/db_utils.go
	database.EnumeratePosts(db, []byte{17})
	return nil
}
