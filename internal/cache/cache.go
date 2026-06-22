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
	Entries map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		Entries: make(map[string]cacheEntry),
	}

	go newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newCacheEntry := cacheEntry{createdAt: time.Now(), val: val}
	c.Entries[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry, exists := c.Entries[key]
	if !exists {
		return []byte{}, false
	}
	return cacheEntry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.Entries, key)
			}
		}
		c.mu.Unlock()
	}
}