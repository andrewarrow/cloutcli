package cloutcli

import (
	"fmt"

	"github.com/andrewarrow/cloutcli/database"
)

func QuerySqlitePosts(term string) {
	sdb := database.OpenSqliteDB()
	defer sdb.Close()

	rows, err := sdb.Query("select body from posts where body like '%" + term + "%'")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var body string
		rows.Scan(&body)
		fmt.Println(body)
	}
}
func QuerySqliteUsers(s string) {
	db := database.OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select username,bio from users where (username like '%" + s + "%') or (bio like '%" + s + "%')")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		var bio string
		rows.Scan(&username, &bio)
		fmt.Println(username, bio)
	}
}
