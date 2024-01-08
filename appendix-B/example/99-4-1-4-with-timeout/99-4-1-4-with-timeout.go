package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  ctxMain := context.Background()

  go func() {
    // 5초 후에 타임아웃되는 context.Context를 생성
    ctxTimeout, cancelTimeout := context.WithTimeout(ctxMain, 5 * time.Second)
    // context.CancelFunc를 호출하여 해제하는 것을 잊지 않는다.
    defer cancelTimeout()

    // context.Context가 끝나기를 기다린다
    <-ctxTimeout.Done()

    // 정확히 5초 후에 출력
    fmt.Println("timeout!")
  }()

  go func() {
    // 3초 후에 타임아웃되는 context.Context를 생성
    ctxDeadline, cancelDeadline := context.WithDeadline(
      ctxMain,
      // 현재 시간에 3초를 더한다
      time.Now().Add(3 * time.Second),
    )
    // context.CancelFunc를 호출하여 해제하는 것을 잊지 않는다.
    defer cancelDeadline()

    // context.Context가 끝나기를 기다린다
    <-ctxDeadline.Done()

    // 정확히 3초 후에 출력
    fmt.Println("deadline!")
  }()

  // 표준 출력으로 10초 동안 매 초마다 n초...를 출력하는 코드
  for i := 0; i < 10; i++ {
    fmt.Printf("%d sec...\n", i)
    time.Sleep(1 * time.Second)
  }
}
