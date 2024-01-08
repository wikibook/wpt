package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  // sync.Mutex를 값으로 생성
  var mu sync.Mutex

  // mu 을 락
  mu.Lock()

  // 1초 후에 mu를 언락
  go func() {
    <-time.After(1 * time.Second)
    mu.Unlock()
  }()

  // 값을 전달하면 복사되어서 1초 후의 Unlock이 함수 호출 위치의 mu에 전달되지 않으므로
  // 함수 내부의 Lock이 계속 해제되지 않고 deadlock이 발생하여 Go 런타임이 강제로 종료
  LockWithValue(mu)

  // 참조 전달인 경우 1초 후의 Unlock이 함수 내부의 mu에 정상적으로 전달되므로
  // deadlock이 발생하지 않고 "mutex unlocked"가 표준 출력에 표시되며 프로그램이 정상적으로 종료
  // LockWithReference(&mu)

  fmt.Println("mutex unlocked")
}

// 값을 전달받아 sync.Mutex를 사용하는 함수
func LockWithValue(mu sync.Mutex) {
  mu.Lock()
  mu.Unlock()
}

// 참조 전달로 sync.Mutex를 사용하는 함수
func LockWithReference(mu *sync.Mutex) {
  mu.Lock()
  mu.Unlock()
}
