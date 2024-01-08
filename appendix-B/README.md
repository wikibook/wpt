# 부록 B ISUCON 벤치 마커 구현

부록B "ISUCON 벤치 마커 구현"의 샘플 코드입니다.

## 벤치 마커에 자주 발생하는 구현 패턴

### `context.Context` 생성

- [99-4-1-1-context.go](./example/99-4-1-1-context/99-4-1-1-context.go)

#### `context.TODO` 사용소

- [99-4-1-1-c-context-todo.go](./example/99-4-1-1-c-context-todo/99-4-1-1-c-context-todo.go)


#### `context.WithCancel`에 의한`context.Context`의 중단

- [99-4-1-3-with-cancel.go](./example/99-4-1-3-with-cancel/99-4-1-3-with-cancel.go)

#### `context.WithTimeout(ctx, d)`와 `context.WithDeadline(ctx, t)`에 의한 시간 제한의`context.Context`중단

- [99-4-1-4-with-timeout.go](./example/99-4-1-4-with-timeout/99-4-1-4-with-timeout.go)

### `time`과`context`로 루프 패턴

- [99-4-2-loop-with-context.go](./example/99-4-2-loop-with-context/99-4-2-loop-with-context.go)
- [99-4-2-loop-with-context-time-after.go](./example/99-4-2-loop-with-context-time-after/99-4-2-loop-with-context-time-after.go)
- [99-4-2-loop-with-context-time-after-long-time.go](./example/99-4-2-loop-with-context-time-after-long-time/99-4-2-loop-with-context-time-after-long-time.go)

### `sync` 패키지 사용

#### `sync.WaitGroup`에 의한 대기

- [99-4-3-1-waitGroup.go](./example/99-4-3-1-waitGroup/99-4-3-1-waitGroup.go)

#### `sync.Mutex`와`sync.RWMutex`에 의한 읽기 / 쓰기 잠금

- [99-4-3-2-mutex.go](./example/99-4-3-2-mutex/99-4-3-2-mutex.go)

#### `sync.WaitGroup`과`sync.Mutex`를 값으로 전달하여 발생하는 교착 상태와 panic

- [99-4-3-3-panic-with-mutex.go](./example/99-4-3-3-panic-with-mutex/99-4-3-3-panic-with-mutex.go)

## private-isu를 대상으로 벤치 마커 구현

이 소스 코드는 별도 리포지토리에 있습니다.

https://github.com/rosylilly/private-isu-benchmarker/

##### (컬럼)`fmt.Stringer`와`fmt.GoStringer`를 구현한다.

- [99-5-2-c-stringer.go](./example/99-5-2-c-stringer/99-5-2-c-stringer.go)
