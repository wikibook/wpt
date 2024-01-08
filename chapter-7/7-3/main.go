package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

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

	// singleflight를 사용하면 동시에 여러 번 호출되더라도 두 번째 호출부터는 첫 번째 실행이 완료될 때까지 대기한다
	vv, err, _ := group.Do(fmt.Sprintf("cacheGet_%d", key), func() (interface{}, error) {
		value := HeavyGet(key)
		c.Set(key, value)
		return value, nil
	})

	if err != nil {
		panic(err)
	}

	// interface {} 형이므로 int 형으로 캐스팅
	return vv.(int)
}

// 실제로 데이터베이스에 대한 액세스 등이 발생한다
// 이번에는 1초 동안 sleep한 다음 key의 2배를 반환한다
func HeavyGet(key int) int {
	log.Printf("call HeavyGet %d\n", key)
	time.Sleep(time.Second)
	return key * 2
}

func main() {
	mCache := NewCache()

	for i := 0; i < 100; i++ {
		go func(i int) {
			// 0부터 9까지의 각 키를 거의 동시에 10회 취득하지만 각각 한 번밖에 HeavyGet는 실행되지 않는다
			mCache.Get(i % 10)
		}(i)
	}

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		log.Println(mCache.Get(i))
	}
}
