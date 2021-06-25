package database

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/dgraph-io/badger/v3"
)

func PostsByAuthor(sdb *sql.DB, db *badger.DB, author string) {

	postMap := map[string]bool{}
	prefix := []byte{17}
	goal := UsernameToPub(db, author)
	//prefix = append(prefix, UsernameToPub(db, author)...)
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		i := 0
		for itr.Seek(nil); itr.Valid(); itr.Next() {
			if i%1000 == 0 {
				fmt.Println("iteration", i)
			}
			i++
		}

		/*
			for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
				if i%1000 == 0 {
					fmt.Println("iteration", i)
				}
				val, _ := it.Item().ValueCopy(nil)

				post := &lib.PostEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(post)

				if bytes.Compare(post.PosterPublicKey, goal) == 0 {
					postMap[base58.Encode(post.PostHash.Bytes())] = true
				}

				i++
			}
		*/
		return nil
	})
	fmt.Println("postMap", len(postMap))

	/*
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
		fmt.Println("likeMap", len(likeMap))

		recloutMap := map[string]bool{}
		prefix = []byte{39}
		db.View(func(txn *badger.Txn) error {
			opts := badger.DefaultIteratorOptions
			it := txn.NewIterator(opts)
			defer it.Close()

			for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
				val, _ := it.Item().ValueCopy(nil)
				re := &lib.RecloutEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(re)

				if postMap[base58.Encode(re.RecloutedPostHash.Bytes())] {
					recloutMap[base58.Encode(re.ReclouterPubKey)] = true
					InsertRecloutSqlite(sdb, re)
				}
			}
			return nil
		})
		fmt.Println("recloutMap", len(recloutMap))

		diamondMap := map[string]bool{}
		prefix = []byte{41}
		db.View(func(txn *badger.Txn) error {
			opts := badger.DefaultIteratorOptions
			it := txn.NewIterator(opts)
			defer it.Close()

			for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
				val, _ := it.Item().ValueCopy(nil)
				de := &lib.DiamondEntry{}
				gob.NewDecoder(bytes.NewReader(val)).Decode(de)

				if postMap[base58.Encode(de.DiamondPostHash.Bytes())] {
					diamondMap[base58.Encode(de.SenderPKID[:])] = true
					InsertDiamondSqlite(sdb, de)
				}
			}
			return nil
		})
		fmt.Println("diamondMap", len(diamondMap))

		prefix = []byte{23}
		db.View(func(txn *badger.Txn) error {
			opts := badger.DefaultIteratorOptions
			it := txn.NewIterator(opts)
			defer it.Close()

			i := 0
			for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
				key := it.Item().Key()
				pub58 := base58.Encode(key[1:])
				if likeMap[pub58] || recloutMap[pub58] || diamondMap[pub58] {
					val, _ := it.Item().ValueCopy(nil)
					profile := &lib.ProfileEntry{}
					gob.NewDecoder(bytes.NewReader(val)).Decode(profile)
					InsertProfileSqlite(sdb, profile)
					i++
				}
			}
			fmt.Println("i", i)
			return nil
		})
	*/
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
