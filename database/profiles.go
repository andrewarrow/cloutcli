package database

import (
	"bytes"
	"encoding/gob"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/dgraph-io/badger/v3"
)

func EnumerateProfiles(db *badger.DB, c *chan *lib.ProfileEntry) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := []byte{23}

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			val, _ := nodeIterator.Item().ValueCopy(nil)

			profile := &lib.ProfileEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(profile)
			*c <- profile
		}
		return nil
	})

}
func UsernameToPub(db *badger.DB, username string) []byte {
	pub := []byte{}

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := []byte{25}
		key := append(prefix, []byte(username)...)
		profileEntryItem, _ := txn.Get(key)
		if profileEntryItem != nil {
			pub, _ = profileEntryItem.ValueCopy(nil)
		}

		return nil
	})

	return pub

}
func LookupUsername(db *badger.DB, pkid []byte) string {

	username := ""
	err := db.View(func(txn *badger.Txn) error {

		key := append([]byte{23}, pkid...)
		profile, err := txn.Get(key)

		if err != nil {
			return err
		}
		profile.Value(func(valBytes []byte) error {
			profile := &lib.ProfileEntry{}
			gob.NewDecoder(bytes.NewReader(valBytes)).Decode(profile)
			username = string(profile.Username)
			return nil
		})

		return nil
	})

	if err == nil {
		return username
	}

	return "404"
}
