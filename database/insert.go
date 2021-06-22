package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/btcsuite/btcutil/base58"
)

func InsertPostSqlite(sdb *sql.DB, post *lib.PostEntry) {
	tx, _ := sdb.Begin()

	body := string(post.Body)
	hash := base58.Encode(post.PostHash.Bytes())
	author := base58.Encode(post.PosterPublicKey)

	s := `insert into posts (author, hash, body, created_at) values (?, ?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
		return
	}
	_, e = thing.Exec(author, hash, body, time.Now())
	if e != nil {
		fmt.Println(e)
		return
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
func InsertProfileSqlite(db *sql.DB, profile *lib.ProfileEntry) {
	tx, _ := db.Begin()

	pub58 := base58.Encode(profile.PublicKey)
	username := string(profile.Username)
	bio := string(profile.Description)

	s := `insert into users (bio, username, pub58, created_at) values (?, ?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(bio, username, pub58, time.Now())
	if e != nil {
		fmt.Println(e)
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
func InsertFollowee(sdb *sql.DB, followee, follower string) {
	tx, _ := sdb.Begin()

	s := `insert into user_follower (followee, follower) values (?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(followee, follower)
	if e != nil {
		fmt.Println(e)
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
