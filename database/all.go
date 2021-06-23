package database

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/dgraph-io/badger/v3"
)

type EntryHolder struct {
	Flavor   string
	Thing    interface{}
	Follower []byte
	Followed []byte
}

func EnumerateAll(testing bool, db *badger.DB, c *chan EntryHolder) {

	desiredPrefixes := map[string]bool{"17": true,
		"23": true,
		"29": true,
		"39": true,
		"41": true}
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		prefix := []byte{}

		flavorMap := map[string]int{}
		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			key := nodeIterator.Item().Key()
			keyPrefix := fmt.Sprintf("%d", key[0])

			if desiredPrefixes[keyPrefix] == false {
				continue
			}

			val, _ := nodeIterator.Item().ValueCopy(nil)
			holder := EntryHolder{}
			if keyPrefix == "17" {
				post := &lib.PostEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(post)
				holder.Flavor = "post"
				holder.Thing = post
			} else if keyPrefix == "23" {
				profile := &lib.ProfileEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(profile)
				holder.Flavor = "profile"
				holder.Thing = profile
			} else if keyPrefix == "29" {
				follower := key[1:34]
				followed := key[34:]
				holder.Flavor = "follow"
				holder.Follower = follower
				holder.Followed = followed
			} else if keyPrefix == "30" {
				le := &lib.LikeEntry{}
				le.LikerPubKey = key[1:34]
				le.LikedPostHash = key[34:]
				holder.Flavor = "like"
				holder.Thing = le
			} else if keyPrefix == "39" {
				holder.Flavor = "reclout"
				re := &lib.RecloutEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(re)
				holder.Thing = re
			} else if keyPrefix == "41" {
				de := &lib.DiamondEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(de)
				holder.Flavor = "diamond"
				holder.Thing = de
			}
			if testing && flavorMap[holder.Flavor] > 1000 {
				continue
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
