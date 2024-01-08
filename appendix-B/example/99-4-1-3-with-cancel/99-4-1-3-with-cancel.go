package main

import (
  "context"
  "fmt"
)

func main() {
  // 일반적으로 함수 내에서 생성된 context.CancelFunc는 defer 등을 사용하여 context.Context가 종료되도록 보장
  ctxParent, cancelParent := context.WithCancel(context.Background())
  defer cancelParent()

  ctxChild, cancelChild := context.WithCancel(ctxParent)
  defer cancelChild()

  // 부모 context.Context의 중단은 자식 context.Context에도 전파
  // context.CancelFunc는 원하는 만큼 호출 가능 (두 번째 이후의 호출은 아무 작업도 수행하지 않는다).
  cancelParent()
  // 부모 context.Context로 중단이 전파되지 않도록 하려면, 바로 위의 줄을 주석 처리하고, 바로 아래 줄의 주석을 해제
  // cancelChild()

  // context.Canceled가 반환, 자식 context.Context 의 context.CancelFunc 만 실행하면 nil
  fmt.Printf("parent.Err is %v\n", ctxParent.Err())
  // => parent.Err is context canceled

  // 부모 context.Context의 취소가 전파되어 자식 context.Context도 context.Canceled 상태가 된다
  fmt.Printf("child.Err is %v\n", ctxChild.Err())
  // => child.Err is context canceled
}
