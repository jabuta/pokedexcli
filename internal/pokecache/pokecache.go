package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(t time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}
	go cache.reapLoop(t)
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c Cache) reapLoop(t time.Duration) {
	clearTicker := time.NewTicker(t)
	for {
		<-clearTicker.C
		c.mu.Lock()
		for k, v := range c.cache {
			if time.Since(v.createdAt) <= t {
				continue
			}
			delete(c.cache, k)
		}
		c.mu.Unlock()
	}
}
