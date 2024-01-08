package main

import (
	"log"
	"sync"
	"time"
)

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

	v = HeavyGet(key)

	c.Set(key, v)

	return v
}

// 실제로 데이터베이스에 대한 액세스 등이 발생한다
// 이번에는 1초 동안 sleep한 다음 key의 2배를 반환한다.
func HeavyGet(key int) int {
	time.Sleep(time.Second)
	return key * 2
}

func main() {
	mCache := NewCache()
	log.Println(mCache.Get(3))
}
