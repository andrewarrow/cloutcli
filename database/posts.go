package database

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/btcsuite/btcutil/base58"
	"github.com/dgraph-io/badger/v3"
)

func PostsByAuthor(db *badger.DB, authors, pub58s []string) {

	authorToSdb := map[string]*sql.DB{}
	goalToUsername := map[string]string{}

	for i, author := range authors {
		sdb := OpenSqliteDB("user_sqlites/" + author + ".db")
		CreateSchema(sdb)
		authorToSdb[author] = sdb
		decoded := base58.Decode(pub58s[i])
		goal := decoded[3 : len(decoded)-4]
		goalToUsername[base58.Encode(goal)] = author
	}
	HandleGatherLikesRecloutsDiamonds(authorToSdb, db, goalToUsername)
}

func HandleGatherLikesRecloutsDiamonds(sdbMap map[string]*sql.DB, db *badger.DB, goalMap map[string]string) {

	prefix := []byte{17}
	postMap := map[string]map[string]bool{}
	likeMap := map[string]map[string]bool{}
	recloutMap := map[string]map[string]bool{}
	diamondMap := map[string]map[string]bool{}
	for k, _ := range sdbMap {
		postMap[k] = map[string]bool{}
		likeMap[k] = map[string]bool{}
		recloutMap[k] = map[string]bool{}
		diamondMap[k] = map[string]bool{}
	}
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		i := 0
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			if i%1000 == 0 {
				fmt.Println("iteration", i)
			}
			val, _ := it.Item().ValueCopy(nil)

			post := &lib.PostEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(post)

			compare := base58.Encode(post.PosterPublicKey)
			if goalMap[compare] != "" {
				author := goalMap[compare]
				postMap[author][base58.Encode(post.PostHash.Bytes())] = true
			}

			i++
		}
		return nil
	})
	fmt.Println("postMap", len(postMap))

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
			for k, v := range postMap {
				if v[base58.Encode(le.LikedPostHash)] {
					likeMap[k][base58.Encode(le.LikerPubKey)] = true
					InsertLikeSqlite(sdbMap[k], le)
				}
			}
		}
		return nil
	})
	fmt.Println("likeMap", len(likeMap))

	prefix = []byte{39}
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			val, _ := it.Item().ValueCopy(nil)
			re := &lib.RecloutEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(re)

			for k, v := range postMap {
				if v[base58.Encode(re.RecloutedPostHash.Bytes())] {
					recloutMap[k][base58.Encode(re.ReclouterPubKey)] = true
					InsertRecloutSqlite(sdbMap[k], re)
				}
			}
		}
		return nil
	})
	fmt.Println("recloutMap", len(recloutMap))

	prefix = []byte{41}
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			val, _ := it.Item().ValueCopy(nil)
			de := &lib.DiamondEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(de)

			for k, v := range postMap {
				if v[base58.Encode(de.DiamondPostHash.Bytes())] {
					diamondMap[k][base58.Encode(de.SenderPKID[:])] = true
					InsertDiamondSqlite(sdbMap[k], de)
				}
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
			for k, sdb := range sdbMap {
				if likeMap[k][pub58] || recloutMap[k][pub58] || diamondMap[k][pub58] {
					val, _ := it.Item().ValueCopy(nil)
					profile := &lib.ProfileEntry{}
					gob.NewDecoder(bytes.NewReader(val)).Decode(profile)
					InsertProfileSqlite(sdb, profile)
					i++
				}
			}
		}
		fmt.Println("i", i)
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
