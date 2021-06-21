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
