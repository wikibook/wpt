package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// UNIX 시간을 사용하여 의사 난수의 시드 값을 초기화
	rand.Seed(time.Now().Unix())

	// Users를 새로 생성
	users := NewUsers(
		// 먼저 연속적인 ID를 가진 5 명을 생성
		WithSequentialIDUsers(5),
		// 그런 다음 무작위 ID를 가진 5 명을 생성
		WithRandomIDUsers(5),
		// 마지막으로 다시 연속적인 ID를 가진 5 명을 생성
		WithSequentialIDUsers(5),
	)

	// 등록된 Users 수를 표준 출력에 표시
	fmt.Printf("users count: %d\n", users.Len())

	// 등록된 사용자의 ID를 등록된 순서대로 차례로 표준 출력에 표시
	users.ForEach(func(i int, u *User) {
		fmt.Printf("%02d: user id: %d\n", i+1, u.ID)
	})
}

// User는 ID만을 가지고 있는 간단한 구조체
type User struct {
	ID int
}

// User 목록과 ID 맵을 가지며 읽기 및 쓰기 잠금을 가진 구조체
type Users struct {
	mu   sync.RWMutex
	list []*User
	dict map[int]*User
}

// Users를 생성할 때 지정할 수 있는 UsersOption의 형식을 정의
// 이렇게 함수를 통해 옵션을 지정할 수 있도록 하는 것이 Functional Option 패턴
type UsersOption func(users *Users)

// Users를 생성하는 함수
// 가변 인수로 UsersOption을 받아 여러 개의 UsersOption을 받을 수 있도록 만든다
func NewUsers(opts ...UsersOption) *Users {
	// Users の生成
	users := &Users{
		// sync.RWMutex는 값 형식의 필드이므로 생성 코드를 작성하지 않아도 문제없다
		// mu:   sync.RWMutex{},

		// list는 슬라이스, dict는 맵이므로 초기화할 때 생성해야 한다
		// nil이 되어 버린다
		list: make([]*User, 0),
		dict: make(map[int]*User, 0),
	}

	// 인수에 전달된 UsersOption을 순차적으로 실행
	for _, opt := range opts {
		opt(users)
	}

	// 모든 UsersOption이 적용된 Users 반환
	return users
}

// Users에 User를 추가하는 메서드
// ID가 이미 등록되어 있으면 false를 반환하고 등록하지 않는다
func (u *Users) Add(user *User) (ok bool) {
	// 쓰기를 잠그고 함수를 종료할 때 잠금을 해제
	u.mu.Lock()
	defer u.mu.Unlock()

	// 추가하려는 User의 ID가 0 이하인 경우 추가하지 않는다
	if user.ID <= 0 {
		return false
	}

	// 추가하려는 User의 ID가 이미 등록되어 있는 경우 추가하지 않는다
	if _, found := u.dict[user.ID]; found {
		return false
	}

	u.list = append(u.list, user)
	u.dict[user.ID] = user
	return true
}

// Users에 등록된 User의 수를 반환하는 메소드
func (u *Users) Len() int {
	// 쓰기를 잠그고 함수를 종료할 때 잠금을 해제
	u.mu.RLock()
	defer u.mu.RUnlock()

	// Users.list の len は登録済みの User の数
	return len(u.list)
}

// Users에 등록된 모든 User에 대해 추가된 순서대로 함수를 실행하는 메소드
func (u *Users) ForEach(f func(i int, u *User)) {
	// 쓰기를 잠그고 함수를 종료할 때 잠금을 해제
	u.mu.RLock()
	defer u.mu.RUnlock()

	// 리스트를 루프로 반복하고 순서대로 함수를 실행
	for i, u := range u.list {
		f(i, u)
	}
}

// 등록된 User가 가지고 있는 ID 중에서 최대 ID를 반환하는 메소드
func (u *Users) MaxID() int {
	// 쓰기를 잠그고 함수를 종료할 때 잠금을 해제
	u.mu.RLock()
	defer u.mu.RUnlock()

	// User가 한 명도 등록되어 있지 않으면 0을 반환
	maxID := 0
	for _, user := range u.list {
		// 현재의 maxID보다 User의 ID가 크다면 덮어쓴다
		if user.ID > maxID {
			maxID = user.ID
		}
	}

	return maxID
}

// count 매개변수의 수만큼 연속된 ID를 가진 사용자를 Users에 추가하는 옵션
func WithSequentialIDUsers(count int) UsersOption {
	return func(u *Users) {
		for i := 0; i < count; i++ {
			// 등록된 최대 ID를 가져와서 그 값에 1을 더한 ID로 순서를 매긴다
			id := u.MaxID() + 1
			user := &User{
				ID: id,
			}

			// 최대 ID보다 크기 때문에 항상 추가가 성공할 것으로 간주하고 추가의 성공을 확인하지 않는다
			u.Add(user)
		}
	}
}

// Users에 count 매개변수로 지정된 수만큼 무작위 ID를 가진 사용자(User)를 추가하는 옵션
func WithRandomIDUsers(count int) UsersOption {
	return func(u *Users) {
		for i := 1; i <= count; i++ {
			// 등록된 최대 ID에서 +50에서 -50 사이의 무작위 ID를 생성하여 번호를 부여
			id := u.MaxID() + rand.Intn(101) - 50
			user := &User{
				ID: id,
			}

			// ID가 중복되거나 음수가 되어 등록에 실패할 수 있으므로
			// 실패한 경우 요청된 수만큼 생성할 수 있도록 루프 카운터를 하나 감소
			ok := u.Add(user)
			if !ok {
				i--
			}
		}
	}
}