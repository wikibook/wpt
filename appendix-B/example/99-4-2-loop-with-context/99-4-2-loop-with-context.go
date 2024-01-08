package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  // 5초 후에 타임아웃되는 context.Context를 생성
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

L: // 루프 탈출을 위한 라벨
  for {
    // 각 루프에 대한 출력
    fmt.Println("loop")

    select {
      // ctx가 종료되었다면, L 라벨로 이동해 for 루프를 종료
      // 단순히 break 이라고 쓰면 select 의 break 가 되어 버려 무한 루프가 계속되므로 주의
      case <-ctx.Done():
        break L
      // ctx가 종료되지 않았다면, 1초간 기다린다
      default:
        time.Sleep(1 * time.Second)
    }
  }
}
