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
	prefix := []byte{35}
	prefix = append(prefix, UsernameToPub(db, author)...)
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		timestampSizeBytes := 8
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			rawKey := it.Item().Key()
			keyWithoutPrefix := rawKey[1:]
			publicKeySizeBytes := lib.HashSizeBytes + 1
			hash := keyWithoutPrefix[(publicKeySizeBytes + timestampSizeBytes):]
			posts = append(posts, base58.Encode(hash))
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
