package main

import (
  "context"
)

func main() {
  // context.Context를 지원하지 않는 함수를 호출
  // 사실 이 메인 함수 내에서 context.Context를 생성하여 전달하고 싶음
  ContextNotSupportedFunc()
}

// context.Context를 인자로 받지 않는 함수
// 나중에 context.Context를 인자로 받도록 변경될 예정
func ContextNotSupportedFunc() {
  // 전달할 context.Context가 없으므로, 임시적으로 context.TODO를 생성하여 전달
  RequiredContextFunc(context.TODO())
}

// 인자로 context.Context가 필요한 함수
func RequiredContextFunc(ctx context.Context) {
}
