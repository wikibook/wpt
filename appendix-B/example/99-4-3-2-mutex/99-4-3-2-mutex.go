package main

import (
  "fmt"
  "sync"
)

func main() {
  // int 슬라이스를 생성
  userIDs := []int{}
  // sync.Mutex 를 생성
  userIDsLock := &sync.Mutex{}

  // 처리 대기에 사용할 sync.WaitGroup을 생성
  wg := &sync.WaitGroup{}

  for i := 0; i < 20; i++ {
    wg.Add(1)
    go func(id int) {
      defer wg.Done()

      // userIDs에 대한 쓰기 경합을 방지하기 위해 락을 사용
      // 다른 고루틴에서 이미 락이 설정되어있는 경우, 해당 락이 해제될 때까지 여기서 처리가 차단
      userIDsLock.Lock()
      // 데이터를 슬라이스에 추가
      userIDs = append(userIDs, id)
      // 락의 해제
      userIDsLock.Unlock()
    }(i)
  }

  // 모든 추가 처리를 대기
  wg.Wait()

  // 추가된 모든 값들을 표시
  // 고루틴은 시작 순서대로 실행되지 않으므로 실행할 때마다 추가 순서가 다를 수 있다
  fmt.Printf("userIDs: %v\n", userIDs)
}
