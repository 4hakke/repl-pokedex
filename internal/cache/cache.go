package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{entries: make(map[string]cacheEntry)}
	cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for t := range ticker.C {
			cache.flushOutdatedEntries(interval, t)
		}
	}()
}

func (c *Cache) flushOutdatedEntries(interval time.Duration, time time.Time) {
	then := time.Add(-interval)
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.entries {
		if value.createdAt.Before(then) {
			delete(c.entries, key)
		}
	}
}
