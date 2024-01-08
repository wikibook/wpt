package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  // sync.WaitGroup의 생성
  wg := &sync.WaitGroup{}

  // 이 코드 예제에서는 대기해야 하는 처리가 2개로 확실하기 때문에 wg.Add의 인수로 2를 전달하고 있다
  // wg.Add로 추가한 수 이상으로 wg.Done을 호출하면 패닉이 발생할 수 있으니 주의
  wg.Add(2)

  // 루프 중에서 고루틴을 생성하는 경우에는 각 생성 시에 wg.Add(1)을 호출하는 것이 좋ㅏ
  // wg.Add(1)
  go func() {
    // 자주 발생하는 오류 중 하나로는 고루틴 내에서 wg.Add(1)을 하는 경우가 있다
    // 그러나 이런 경우에는 고루틴이 시작되기 전에 wg.Wait()에 도달하여 대기가 끝나버릴 수 있으므로
    // 고루틴 내에서 wg.Add(1)을 하지 않도록 주의

    // 처리가 완료되고 함수를 빠져나갈 때 wg.Done이 확실하게 호출되도록 하기 위해 처음에 defer를 사용하여 wg.Done을 호출한다
    defer wg.Done()

    // 표준 출력으로 5초간 매초 표시
    for i := 0; i < 5; i++ {
      fmt.Printf("wg 1: %d / 5\n", i+1)
      time.Sleep(1 * time.Second)
    }
  }()

  // wg.Add(1)
  go func() {
    // 처리가 완료되고 함수를 빠져나갈 때 wg.Done이 확실하게 호출되도록 하기 위해 처음에 defer를 사용하여 wg.Done을 호출한다
    defer wg.Done()

    // 표준 출력으로 5초간 매초 표시
    for i := 0; i < 5; i++ {
      fmt.Printf("wg 2: %d / 5\n", i+1)
      time.Sleep(1 * time.Second)
    }
  }()

  // 이곳에서는 두 개의 고루틴이 종료되고 wg.Done이 호출되기를 기린다
  // wg.Wait의 반환 값이 없으며, 채널을 통한 종료 알림 수신 등이 불가능하므로 주의해야 한ㅏ
  wg.Wait()

  fmt.Println("wg: done")
}
