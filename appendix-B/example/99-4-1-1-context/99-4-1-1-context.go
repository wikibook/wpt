package main

import (
  "context"
)

func main() {
  // 메인 함수의 시작 부분에서 새로운 context.Context를 생성
  ctx := context.Background()

  // 컨텍스트를 사용하는 처리
  ExampleContextFunc(ctx)
}

func ExampleContextFunc(ctx context.Context) {
  // 이 함수에서는 새로운 context.Context를 생성하지 않고, 
  // 다른 함수가 context.Context를 필요로 할 경우 받은 ctx를 전달
}
