package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
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
var mc *memcache.Client

func main() {
	var err error
	// 데이터베이스에 연결
	db, err = sqlx.Open("mysql", "isuconp:@tcp(localhost:3306)/isuconp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	// memcached에 연결
	mc = memcache.New("127.0.0.1:11211")

	results := []Post{}
	// 게시물 목록 얻기
	err = db.Select(&results, "SELECT `id`, `user_id`, `body`, `mime`, `created_at` FROM `posts` ORDER BY `created_at` DESC LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range results {
		// 게시물 목록에 사용자 정보 부여
		p.User = getUser(p.UserID)
	}
	out, _ := json.Marshal(results)
	fmt.Fprint(os.Stdout, string(out))
}

func getUser(id int) User {
	user := User{}
	// memcached에서 사용자 정보 얻기 ❶
	it, err := mc.Get(fmt.Sprintf("user_id:%d", id))
	if err == nil {
		// 사용자 정보가 있으면 JSON을 디코딩하고 반환
		err := json.Unmarshal(it.Value, &user)
		if err == nil {
			return user
		}
	}
	// 데이터베이스에서 사용자 정보 검색 ❷
	err = db.Get(&user, "SELECT * FROM `users` WHERE `id` = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	// JSON으로 인코딩
	j, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	// memcached에 저장 ❸
	mc.Set(&memcache.Item{
		Key:        fmt.Sprintf("user_id:%d", id),
		Value:      j,
		Expiration: 3600,
	})
	return user
}
