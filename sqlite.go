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
