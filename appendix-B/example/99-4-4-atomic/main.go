// Windows 등의 환경에서는 제대로 작동하지 않을 수 있다

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	// 10초 후에 종료되는 context.Context를 생성
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// SIGUSR1을 보내는 명령어 예시
	fmt.Println("Reset counter:")
	fmt.Printf("  $ kill -SIGUSR1 %d\n", os.Getpid())

	// int64 형식으로 초기화하고 값을 0으로 설정
	i := int64(0)
	// 첫 번째 값으로 10을 기입
	atomic.StoreInt64(&i, 10)

	// 대기를 위해 sync.WaitGroup을 생성
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			timer := time.After(1 * time.Second)

			// 1초마다 현재 값을 가져와 표준 출력에 표시
			n := atomic.LoadInt64(&i)
			fmt.Printf("load now: %d\n", n)

			// 1초마다 반복하기 위한 select문을 사용
			select {
			case <-ctx.Done():
				return
			case <-timer:
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			timer := time.After(100 * time.Millisecond)

			// 0.1초마다 값을 증가시킨다
			// 증가한 결과 값을 사용하려면 반환값을 활용
			atomic.AddInt64(&i, 1)

			// 0.1초마다 반복하기 위한 select문을 사용
			select {
			case <-ctx.Done():
				return
			case <-timer:
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			timer := time.After(10 * time.Millisecond)

			// 0.01초마다 현재 값을 가져와서 50이면 0으로 설정
			swapped := atomic.CompareAndSwapInt64(&i, 50, 0)
			if swapped {
				// 값을 변경하는 데 성공했을 때에만 표준 출력에 로그를 표시
				fmt.Printf("CAS now: reset zero\n")
			}

			// 0.01초마다 반복하기 위한 select문을 사용
			select {
			case <-ctx.Done():
				return
			case <-timer:
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		// SIGUSR1을 수신하는 채널을 생성
		sig := make(chan os.Signal, 1)
		// SIGUSR1이 도착하면 sig 채널에 기록
		signal.Notify(sig, syscall.SIGUSR1)

		for {
			// context.Context가 종료된 경우 루프를 종료
			// sig 채널에서 수신하면 처리를 실행
			select {
			case <-ctx.Done():
				return
			case <-sig:
			}

			// 이 샘플 코드를 실행 중에 언제든지 프로세스에 SIGUSR1을 보내면 
			// 현재 값을 0으로 설정하고 해당 시점의 값을 표준 출력에 표시
			old := atomic.SwapInt64(&i, 0)
			fmt.Printf("SIGUSR1: reset zero: old: %d\n", old)
		}
	}()

	// 모든 처리가 완료될 때까지 대기
	wg.Wait()
}