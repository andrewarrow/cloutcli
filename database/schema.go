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

create table users (user_id INTEGER PRIMARY KEY AUTOINCREMENT, bio text, username text, pub58 text, created_at datetime);

CREATE UNIQUE INDEX users_idx
  ON users (pub58);

CREATE INDEX users_username_idx
  ON users (username);

create table user_follower (followee text, follower text);

CREATE UNIQUE INDEX uf_followee_follower_idx
  ON user_follower (followee, follower);

CREATE INDEX uf_followee_idx
  ON user_follower (followee);

CREATE INDEX uf_follower_idx
  ON user_follower (follower);

create table diamonds (hash, sender, receiver text, level integer);
CREATE INDEX diamonds_hash ON diamonds (hash);
CREATE INDEX diamonds_receiver ON diamonds (receiver);
CREATE INDEX diamonds_sender ON diamonds (sender);

create table reclouts (hash, other_hash, reclouter text);
CREATE INDEX reclouts_hash ON reclouts (hash);
CREATE INDEX reclouts_other_hash ON reclouts (other_hash);
CREATE INDEX reclouts_reclouter ON reclouts (reclouter);

create table likes (hash, liker text);
CREATE INDEX like_hash ON likes (hash);
CREATE INDEX like_liker ON likes (liker);
`
	_, err := sdb.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
