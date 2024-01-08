package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int       `db:"id" json:"id"`
	AccountName string    `db:"account_name" json:"account_name"`
	Passhash    string    `db:"passhash" json:"passhash"`
	Authority   int       `db:"authority" json:"authority"`
	DelFlg      int       `db:"del_flg" json:"del_flg"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type Post struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	Body      string    `db:"body" json:"body"`
	Mime      string    `db:"mime" json:"mime"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	User      User      `json:"users"`
}

var db *sqlx.DB

func main() {
	var err error
	// 데이터베이스에 연결
	db, err = sqlx.Open("mysql", "isuconp:@tcp(localhost:3306)/isuconp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	results := []Post{}
	// 게시물 목록 얻기 ❶
	err = db.Select(&results, "SELECT `id`, `user_id`, `body`, `mime`, `created_at` FROM `posts` ORDER BY `created_at` DESC LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range results {
		// 게시물 목록에 사용자 정보 부여 ❷
		p.User = getUser(p.UserID)
	}
	out, _ := json.Marshal(results)
	fmt.Fprint(os.Stdout, string(out))
}

func getUser(id int) User {
	user := User{}
	// 사용자 정보 얻기
	err := db.Get(&user, "SELECT * FROM `users` WHERE `id` = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
