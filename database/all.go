package database

import (
	"bytes"
	"encoding/gob"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/dgraph-io/badger/v3"
)

type EntryHolder struct {
	Flavor string
	Thing  interface{}
}

func EnumerateAll(db *badger.DB, c *chan EntryHolder) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		prefix := []byte{17}

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			val, _ := nodeIterator.Item().ValueCopy(nil)

			post := &lib.PostEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(post)
			holder := EntryHolder{}
			holder.Flavor = "post"
			holder.Thing = post
			*c <- holder
		}
		nodeIterator.Close()
		nodeIterator = txn.NewIterator(opts)

		prefix = []byte{23}
		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			val, _ := nodeIterator.Item().ValueCopy(nil)

			profile := &lib.ProfileEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(profile)
			holder := EntryHolder{}
			holder.Flavor = "profile"
			holder.Thing = profile
			*c <- holder
		}
		nodeIterator.Close()
		return nil
	})

}
