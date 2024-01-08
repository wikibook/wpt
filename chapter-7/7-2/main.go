package main

import (
	"log"
	"sync"
	"time"
)

const defaultValue = 100

type Cache struct {
	mu    sync.Mutex
	items map[int]int
}

func NewCache() *Cache {
	m := make(map[int]int)
	c := &Cache{
		items: m,
	}
	return c
}

func (c *Cache) Set(key int, value int) {
	c.mu.Lock()
	c.items[key] = value
	c.mu.Unlock()
}

func (c *Cache) Get(key int) int {
	c.mu.Lock()
	v, ok := c.items[key]
	c.mu.Unlock()

	if ok {
		return v
	}

	go func() {
		// 비동기적으로 캐시 업데이트 처리 실시
		v := HeavyGet(key)

		c.Set(key, v)
	}()

	return defaultValue
}

// 실제로 데이터베이스에 대한 액세스 등이 발생한다
// 이번에는 1초 동안 sleep한 다음 key의 2배를 반환한다
func HeavyGet(key int) int {
	time.Sleep(time.Second)
	return key * 2
}

func main() {
	mCache := NewCache()
	// 처음에는 기본값이 반환된다
	log.Println(mCache.Get(3))
	time.Sleep(time.Second)
	// 다음은 갱신된다
	log.Println(mCache.Get(3))
}
