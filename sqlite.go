package cloutcli

import (
	"fmt"
	"os"
	"strconv"

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
func QuerySqliteFollow(tab, s, degrees string) {
	pub58 := SearchSqliteUsername(s)
	db := database.OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select follower from user_follower where followee = '" + pub58 + "'")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	limit, _ := strconv.Atoi(degrees)
	tabSize := len(tab) / 2

	for rows.Next() {
		var follower string
		rows.Scan(&follower)
		username := SearchSqlitePub58(follower)
		fmt.Printf("%s%s\n", tab, username)

		if tabSize+1 < limit {
			QuerySqliteFollow(tab+"  ", username, degrees)
		}
	}
}
func QuerySqliteNodesInOrder(f *os.File) {
	db := database.OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select username from users order by user_id")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	i := int64(0)
	for rows.Next() {
		var username string
		rows.Scan(&username)
		line := fmt.Sprintf(" n%d [label=\"%s\"];\n", i, username)
		f.Write([]byte(line))
		i++
	}
}
func QuerySqliteNodeConnections(f *os.File) {
	db := database.OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select followee, follower from user_follower")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		var followee string
		var follower string
		rows.Scan(&followee, &follower)
		line := fmt.Sprintf(" n%d -> n%d;", 0, 0)
		f.Write([]byte(line))
		i++
	}
}
func SearchSqliteUsername(s string) string {
	db := database.OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select pub58 from users where username='" + s + "'")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		var pub58 string
		rows.Scan(&pub58)
		return pub58
	}

	return ""
}
func SearchSqlitePub58(s string) string {
	db := database.OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select username from users where pub58='" + s + "'")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		rows.Scan(&username)
		return username
	}

	return ""
}
