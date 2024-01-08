package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  // 5.5초 후에 타임아웃되는 context.Context를 생성
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second + 500 * time.Millisecond)
  defer cancel()

  i := 0

L: // 루프 탈출을 위한 라ㄹ
  for {
    // 각 루프에 대한 출력
    // time.After를 사용하면 루프 5까지만 출력되지만, time.Sleep를 사용하면 루프 6까지 한 번 더 출력
    fmt.Printf("loop %d\n", i)
    i++

    select {
      // ctx가 종료되었다면, L 레이블로 이동하여 for 루프를 종료
      // 단순히 break 이라고 쓰면 select 의 break 가 되어 버려 무한 루프가 계속되므로 주의
      case <-ctx.Done():
        break L
      // ctx가 종료되지 않았다면 1초를 기다리지만, 채널 수신을 사용하고 있기 때문에, ctx가 먼저 종료되면 실행된ㅏ
      case <-time.After(1 * time.Second):

      // time.Sleep을 사용해 기다리는 예시
      // default:
      //   time.Sleep(1 * time.Second)
    }
  }
}
