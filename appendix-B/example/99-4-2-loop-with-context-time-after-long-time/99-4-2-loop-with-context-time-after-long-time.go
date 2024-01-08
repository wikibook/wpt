package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  // 10초 후에 타임아웃되는 context.Context를 생성
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  go LoopWithBefore(ctx)
  go LoopWithAfter(ctx)

  <-ctx.Done()
}

// 루프 시작 부분에서 time.After를 생성하고 기다리는 패턴
// Heavy Process에 1.5초 걸리지만, 1루프의 시간은 3초 이내
func LoopWithBefore(ctx context.Context) {
  // 루프 이전 시간 가져오기
  beforeLoop := time.Now()
  for {
    // 한 루프의 지속 시간을 처음에 설정
    loopTimer := time.After(3 * time.Second)

    // 1.5초가 걸리는 처리
    HeavyProcess(ctx, "BEFORE")

    select {
      case <-ctx.Done():
        return
      // 처음에 생성한 time.After를 사용해 대기
      case <-loopTimer:
        // 한 루프에 걸린 시간을 표준 출력에 표시하고, beforeLoop에 현재 시간을 설정
        fmt.Printf("[BEFORE] loop duration: %.2fs\n", time.Now().Sub(beforeLoop).Seconds())
        beforeLoop = time.Now()
    }
  }
}

// 루프의 끝에서 time.After를 생성하고 대기하는 패턴
// HeavyProcess가 1.5초 걸리고, 이후 3초 대기하므로, 한 루프의 총 시간은 4.5초가 된다
func LoopWithAfter(ctx context.Context) {
  beforeLoop := time.Now()
  for {
    // 1.5초가 걸리는 처리
    HeavyProcess(ctx, "AFTER")

    select {
      case <-ctx.Done():
        return
      // 이 장소에서 생성한 time.After를 사용하여 대기
      case <-time.After(3 * time.Second):
        // 한 루프에 걸린 시간을 표준 출력에 표시하고, beforeLoop에 현재 시간을 설정
        fmt.Printf("[AFTER] loop duration: %.2fs\n", time.Now().Sub(beforeLoop).Seconds())
        beforeLoop = time.Now()
    }
  }
}

// 각 루프가 호출된 위치를 표시하면서 1.5초 대기
func HeavyProcess(ctx context.Context, pattern string) {
  fmt.Printf("[%s] Heavy Process\n", pattern)
  time.Sleep(1 * time.Second + 500 * time.Millisecond)
}
