package database

import (
	"bytes"
	"encoding/gob"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/btcsuite/btcutil/base58"
	"github.com/dgraph-io/badger/v3"
)

func PostsByAuthor(db *badger.DB, author string) []string {

	posts := []string{}
	author58 := base58.Encode(UsernameToPub(db, author))
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := []byte{18}

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			key := nodeIterator.Item().Key()
			if base58.Encode(key[1:34]) == author58 {
				posts = append(posts, base58.Encode(key[34:]))
			}
		}
		return nil
	})

	return posts

}
func EnumeratePosts(db *badger.DB, c *chan *lib.PostEntry) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := []byte{17}

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			val, _ := nodeIterator.Item().ValueCopy(nil)

			post := &lib.PostEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(post)
			*c <- post
		}
		return nil
	})

}
