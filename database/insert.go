package database

import (
	"database/sql"
	"fmt"
	"strings"
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
	pic := base58.Encode(profile.ProfilePic)

	s := `insert into users (pic, bio, username, pub58, created_at) values (?, ?, ?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(pic, bio, username, pub58, time.Now())
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
		if !strings.Contains(e.Error(), "UNIQUE constraint failed") {
			fmt.Println(e)
		}
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
func InsertDiamondSqlite(sdb *sql.DB, de *lib.DiamondEntry) {
	tx, _ := sdb.Begin()

	s := `insert into diamonds (hash, sender, receiver, level) values (?, ?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(base58.Encode(de.DiamondPostHash.Bytes()),
		base58.Encode(de.SenderPKID[:]),
		base58.Encode(de.ReceiverPKID[:]),
		de.DiamondLevel)
	if e != nil {
		fmt.Println(e)
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
func InsertRecloutSqlite(sdb *sql.DB, re *lib.RecloutEntry) {
	tx, _ := sdb.Begin()

	s := `insert into reclouts (hash, other_hash, reclouter) values (?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(base58.Encode(re.RecloutedPostHash.Bytes()),
		base58.Encode(re.RecloutPostHash.Bytes()),
		base58.Encode(re.ReclouterPubKey))
	if e != nil {
		fmt.Println(e)
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
func InsertLikeSqlite(sdb *sql.DB, le *lib.LikeEntry) {
	tx, _ := sdb.Begin()

	s := `insert into likes (hash, liker) values (?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(base58.Encode(le.LikedPostHash),
		base58.Encode(le.LikerPubKey))
	if e != nil {
		fmt.Println(e)
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
