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
	db, err = sqlx.Open("mysql", "isuconp:@tcp(127.0.0.1:3306)/isuconp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	// memcached에 연결
	mc = memcache.New("127.0.0.1:11211")

	results := []Post{}
	err = db.Select(&results, "SELECT `id`, `user_id`, `body`, `mime`, `created_at` FROM `posts` ORDER BY `created_at` DESC LIMIT 30")
	if err != nil {
		log.Fatal(err)
	}
	// 캐시에서 검색할 사용자 ID 목록 만들기 ❶
	userIDs := make([]int, 0)
	for _, p := range results {
		userIDs = append(userIDs, p.UserID)
	}
	// 캐시에서 사용자 정보를 일괄 얻기 ❷
	users := getUsers(userIDs)
	for _, p := range results {
		if u, ok := users[p.UserID]; ok {
			p.User = u
		} else {
			// 캐시에서 검색할 수 없는 경우 데이터베이스에서 검색 ❹
			p.User = getUser(p.UserID)
		}
	}
	out, _ := json.Marshal(results)
	fmt.Fprint(os.Stdout, string(out))
}

// 캐시에서 사용자 정보를 대량으로 검색하는 함수
func getUsers(ids []int) map[int]User {
	// キャッシュのキーのリストを作成
	keys := make([]string, 0)
	for _, id := range ids {
		keys = append(keys, fmt.Sprintf("user_id:%d", id))
	}
	// 결과를 넣을 map(연상 배열)을 작성. 키는 사용자 ID
	users := map[int]User{}
	// 캐시에서 여러 캐시 얻기 ❸
	items, err := mc.GetMulti(keys)
	if err != nil {
		return users
	}
	for _, it := range items {
		u := User{}
		// JSON을 디코딩하고 map에 사용자 ID를 키로 저장
		err := json.Unmarshal(it.Value, &u)
		if err != nil {
			log.Fatal(err)
		}
		users[u.ID] = u
	}
	return users
}

func getUser(id int) User {
	user := User{}
	// 데이터베이스에서 사용자 정보 검색
	err := db.Get(&user, "SELECT * FROM `users` WHERE `id` = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	// JSON으로 인코딩
	j, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	// 캐시에 저장 ❺
	mc.Set(&memcache.Item{
		Key:        fmt.Sprintf("user_id:%d", id),
		Value:      j,
		Expiration: 3600,
	})
	return user
}
