package database

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/dgraph-io/badger/v3"
)

type EntryHolder struct {
	Flavor string
	Thing  interface{}
}

func EnumerateAll(testing bool, db *badger.DB, c *chan EntryHolder) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		prefix := []byte{}

		flavorMap := map[string]int{}
		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			key := nodeIterator.Item().Key()
			keyPrefix := fmt.Sprintf("%d", key[0])

			if keyPrefix != "17" && keyPrefix != "23" {
				continue
			}

			val, _ := nodeIterator.Item().ValueCopy(nil)
			holder := EntryHolder{}
			if keyPrefix == "17" {
				if testing && flavorMap["post"] > 1000 {
					continue
				}
				post := &lib.PostEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(post)
				holder.Flavor = "post"
				holder.Thing = post
			} else if keyPrefix == "23" {
				if testing && flavorMap["profile"] > 1000 {
					continue
				}
				profile := &lib.ProfileEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(profile)
				holder.Flavor = "profile"
				holder.Thing = profile
			}
			*c <- holder
			flavorMap[holder.Flavor]++
		}
		nodeIterator.Close()
		holder := EntryHolder{}
		holder.Flavor = "done"
		*c <- holder
		return nil
	})

}
