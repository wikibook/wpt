package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// 10초 후에 타임아웃되는 context.Context를 생성
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// int 슬라이스를 생성
	// 비어있을 경우 고루틴 실행 타이밍에 따라 제로 나누기 오류가 발생할 수 있으므로 미리 하나의 값을 넣어 둔다
	responseTimes := []int{200}
	// sync.RWMutex의 생성
	responseTimeMutex := &sync.RWMutex{}

	// 처리 대기를 위해 sync.WaitGroup을 생성
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			// responseTimes에 대한 쓰기 락을 취득
			// 이 쓰기 락을 취득하는 동안에는 읽기 락도 차단
			responseTimeMutex.Lock()
			// 100에서 199 사이의 무작위 숫자를 생성하고 슬라이스에 추가
			responseTimes = append(responseTimes, rand.Intn(100)+100)
			// responseTimes에 대한 쓰기 락을 해제
			responseTimeMutex.Unlock()

			// context.Context의 종료에 따라 루프를 종료하거나 100 밀리초를 대기
			select {
			case <-ctx.Done():
				return
			case <-time.After(100 * time.Millisecond):
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			// responseTimes에 대한 읽기 락을 취득
			// 이 락은 다른 읽기 락과는 충돌하지 않지만, 쓰기 락은 차단
			responseTimeMutex.RLock()

			// responseTimes の個数を標準出力へ表示
			fmt.Printf("response times count: %d\n", len(responseTimes))

			// responseTimes에 대한 읽기 락을 해제
			responseTimeMutex.RUnlock()

			// context.Context의 종료에 따라 루프를 종료하거나 1초를 대기
			select {
			case <-ctx.Done():
				return
			case <-time.After(1 * time.Second):
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			// responseTimes에 대한 읽기 락을 취득
			// 이 락은 다른 읽기 락과는 충돌하지 않지만, 쓰기 락은 차단
			responseTimeMutex.RLock()

			// responseTimes의 합계를 구한 후, 개수로 나누어 평균값을 계산
			responseTimeSum := 0
			responseTimeCount := len(responseTimes)
			for _, responseTime := range responseTimes {
				responseTimeSum += responseTime
			}
			responseTimeAverage := responseTimeSum / responseTimeCount

			// responseTimes에 대한 읽기 락을 해제
			responseTimeMutex.RUnlock()

			// 계산한 평균값과 개수를 표준 출력에 표시
			fmt.Printf("response times average: %d / %d\n", responseTimeAverage, responseTimeCount)

			// context.Context의 종료에 따라 루프를 종료하거나 1초를 대기
			select {
			case <-ctx.Done():
				return
			case <-time.After(1 * time.Second):
			}
		}
	}()

	// 모든 처리가 완료될 때까지 대기
	wg.Wait()
}