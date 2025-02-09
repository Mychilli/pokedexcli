package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	data     map[string]CacheEntry
	interval time.Duration
}
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(time time.Duration) *Cache {

	return &Cache{
		data:     make(map[string]CacheEntry),
		interval: time,
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	if _, ok := c.data[key]; ok {
		result := c.data[key].val
		c.mu.Unlock()
		return result, true
	} else {
		c.mu.Unlock()
		fmt.Println("Result Not Found")
		return nil, false
	}
}
func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, value := range c.data {
				fmt.Println(key, value)
				if time.Since(c.data[key].createdAt) >= c.interval {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
