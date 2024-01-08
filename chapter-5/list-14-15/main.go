package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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
	db, err = sqlx.Open("mysql", "isuconp:isuconp@tcp(127.0.0.1:3306)/isuconp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	results := []Post{}
	// 게시물 목록 얻기
	err = db.Select(&results, "SELECT `id`, `user_id`, `body`, `mime`, `created_at` FROM `posts` ORDER BY `created_at` DESC LIMIT 30")
	if err != nil {
		log.Fatal(err)
	}
	// 사용자 ID 목록 만들기 ❶
	userIDs := make([]int, 0)
	for _, p := range results {
		userIDs = append(userIDs, p.UserID)
	}
	// 사용자 정보 미리 로드 ❷
	users := preloadUsers(userIDs)
	for _, p := range results {
		p.User = users[p.UserID]
	}
	out, _ := json.Marshal(results)
	fmt.Fprint(os.Stdout, string(out))
}

// 데이터베이스에서 사용자 정보를 대량으로 검색하는 함수
func preloadUsers(ids []int) map[int]User {
	// 결과를 넣을 map(연상 배열)을 작성. 키는 사용자 ID
	users := map[int]User{}
	// 사용자 목록이 비어 있는 경우
	if len(ids) == 0 {
		return users
	}
	// 사용자 ID용 목록
	params := make([]interface{}, 0)
	// 플레이스홀더용 목록
	placeholders := make([]string, 0)
	for _, id := range ids {
		params = append(params, id)
		// 플레이스홀더용 목록에는 '?'를 넣는다.
		placeholders = append(placeholders, "?")
	}
	us := []User{}
	// IN구를 사용하여 데이터베이스에서 사용자 정보 가져오기 ❸
	// 플레이스홀더 리스트는 ','로 연결하여 쿼리를 작성
	err := db.Select(
		&us,
		"SELECT * FROM `users` WHERE `id` IN ("+strings.Join(placeholders, ",")+")",
		params...,
	)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range us {
		users[u.ID] = u
	}
	return users
}

// list-15 N+1 데이터베이스 프리로드(sqlx.In을 사용하는 방법)
func preloadUsersIn(ids []int) map[int]User {
	// 결과를 넣는 map(연상 배열) 작성.키는 사용자 ID
	users := map[int]User{}
	// 사용자 목록이 비어있는 경우
	if len(ids) == 0 {
		return users
	}
	// IN구를 포함한 쿼리 구축
	// query: 플레이스홀더 전개된 쿼리
	// params: 쿼리 실행 시 전달하는 파라미터
	query, params, err := sqlx.In(
		"SELECT * FROM `users` WHERE `id` IN (?)",
		ids,
	)
	if err != nil {
		log.Fatal(err)
	}
	us := []User{}
	// 데이터베이스에서 사용자 정보 얻기
	err = db.Select(
		&us,
		query,
		params...,
	)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range us {
		users[u.ID] = u
	}
	return users
}
