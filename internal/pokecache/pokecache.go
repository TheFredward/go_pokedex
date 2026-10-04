package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data         map[string]CacheEntry
	mu           sync.Mutex
	currInterval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.data[key] = CacheEntry{time.Now(), val}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	foundValue, ok := c.data[key]
	if !ok {
		return nil, ok
	}
	return foundValue.val, ok

}

func (c *Cache) reaploop() {
	ticker := time.NewTicker(c.currInterval)
	for range ticker.C {
		c.mu.Lock()
		for key, dataEntry := range c.data {
			if time.Since(dataEntry.createdAt) > c.currInterval {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data:         make(map[string]CacheEntry),
		currInterval: interval,
	}
	go cache.reaploop()
	return &cache
}
