package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenSqliteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "clout.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func CreateSchema(sdb *sql.DB) {
	sqlStmt := `
create table posts (author text, hash text, body text, created_at datetime);

CREATE UNIQUE INDEX posts_hash_idx
  ON posts (hash);

CREATE INDEX posts_username_idx
  ON posts (author);

create table users (bio text, username text, pub58 text, created_at datetime);

CREATE UNIQUE INDEX users_idx
  ON users (pub58);

CREATE INDEX users_username_idx
  ON users (username);

create table user_follower (followee text, follower text);

CREATE INDEX uf_followee_idx
  ON user_follower (followee);

CREATE INDEX uf_follower_idx
  ON user_follower (follower);

create table diamonds (hash, sender, receiver text, level integer);
CREATE INDEX diamonds_hash ON diamonds (hash);
CREATE INDEX diamonds_receiver ON diamonds (receiver);
`
	_, err := sdb.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
