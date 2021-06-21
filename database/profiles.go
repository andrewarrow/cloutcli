package database

import (
	"bytes"
	"encoding/gob"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/dgraph-io/badger/v3"
)

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
