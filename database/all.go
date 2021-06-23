package database

import (
	"bytes"
	"encoding/gob"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/dgraph-io/badger/v3"
)

type EntryHolder struct {
	Flavor   string
	Thing    interface{}
	Follower []byte
	Followed []byte
}

func EnumerateAll(okList map[string]bool, testing bool, db *badger.DB, c *chan EntryHolder) {
	desiredPrefixes := map[byte]string{17: "post",
		23: "profile",
		29: "follow",
		30: "like",
		39: "reclout",
		41: "diamond"}
	for b, flavor := range desiredPrefixes {
		if okList[flavor] == false {
			continue
		}
		EnumeratePrefix(flavor, []byte{b}, testing, db, c)
	}
	holder := EntryHolder{}
	holder.Flavor = "done"
	*c <- holder
}

func EnumeratePrefix(flavor string, prefix []byte, testing bool, db *badger.DB, c *chan EntryHolder) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)

		i := 0
		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			i++
			holder := EntryHolder{}
			holder.Flavor = flavor
			if flavor == "post" {
				val, _ := nodeIterator.Item().ValueCopy(nil)
				post := &lib.PostEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(post)
				holder.Thing = post
			} else if flavor == "profile" {
				val, _ := nodeIterator.Item().ValueCopy(nil)
				profile := &lib.ProfileEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(profile)
				holder.Thing = profile
			} else if flavor == "follow" {
				key := nodeIterator.Item().Key()
				follower := key[1:34]
				followed := key[34:]
				holder.Follower = follower
				holder.Followed = followed
			} else if flavor == "like" {
				key := nodeIterator.Item().Key()
				le := &lib.LikeEntry{}
				le.LikerPubKey = key[1:34]
				le.LikedPostHash = key[34:]
				holder.Thing = le
			} else if flavor == "reclout" {
				val, _ := nodeIterator.Item().ValueCopy(nil)
				re := &lib.RecloutEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(re)
				holder.Thing = re
			} else if flavor == "diamond" {
				val, _ := nodeIterator.Item().ValueCopy(nil)
				de := &lib.DiamondEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(de)
				holder.Thing = de
			}
			if testing && i > 1000 {
				continue
			}
			*c <- holder
		}
		nodeIterator.Close()
		return nil
	})

}
