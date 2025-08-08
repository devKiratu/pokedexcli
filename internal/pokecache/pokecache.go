package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	data map[string]cacheEntry
	// interval time.Duration
	mu *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.data[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.data[key]
	if !ok {
		return nil, ok
	}
	return item.val, ok

}

func (c *Cache) reapLoop(interval time.Duration){
	go func(){
		ticker:= time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			// remove the stale items
			c.mu.Lock()
			for k, v := range c.data {
				dur := time.Since(v.createdAt)
				if dur >=interval {
					delete(c.data, k)
				}
			}
			c.mu.Unlock()
		}
	}()
}
