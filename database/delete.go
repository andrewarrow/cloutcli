package database

import (
	"database/sql"
	"fmt"
)

func DeleteUserNodeSqlite(sdb *sql.DB) {
	tx, _ := sdb.Begin()

	s := `delete from user_nodes`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
		return
	}
	_, e = thing.Exec()
	if e != nil {
		fmt.Println(e)
		return
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
