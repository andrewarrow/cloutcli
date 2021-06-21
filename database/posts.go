package database

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/dgraph-io/badger/v3"
)

func EnumeratePosts(db *badger.DB, prefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			val, _ := nodeIterator.Item().ValueCopy(nil)

			post := &lib.PostEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(post)
			fmt.Println(string(post.Body))
		}
		return nil
	})

}
