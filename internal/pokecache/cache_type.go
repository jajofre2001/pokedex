package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	data     map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}
	// Se lanza el reaper
	go c.reapLoop()
	return c
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reapLoop() {

	for {
		time.Sleep(c.interval)

		c.mu.Lock()

		now := time.Now()

		for key, entry := range c.data {
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.data, key)
			}
		}

		c.mu.Unlock()
	}

}
