package database

import (
	"bytes"
	"database/sql"
	"encoding/gob"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/btcsuite/btcutil/base58"
	"github.com/dgraph-io/badger/v3"
)

func PostsByAuthor(sdb *sql.DB, db *badger.DB, author string) {

	postMap := map[string]bool{}
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
			postMap[base58.Encode(hash)] = true
		}
		return nil
	})

	likeMap := map[string]bool{}
	prefix = []byte{30}
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			key := it.Item().Key()
			le := &lib.LikeEntry{}
			le.LikerPubKey = key[1:34]
			le.LikedPostHash = key[34:]
			if postMap[base58.Encode(le.LikedPostHash)] {
				likeMap[base58.Encode(le.LikerPubKey)] = true
				InsertLikeSqlite(sdb, le)
			}
		}
		return nil
	})

	prefix = []byte{23}
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			key := it.Item().Key()
			if likeMap[base58.Encode(key[1:])] {
				val, _ := it.Item().ValueCopy(nil)
				profile := &lib.ProfileEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(profile)
				InsertProfileSqlite(sdb, profile)
			}
		}
		return nil
	})
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
