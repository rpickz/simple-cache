package cache

import (
	"sync"
	"time"
)

const (
	indefinite time.Duration = -1
)

type cacheEntry struct {
	expiresAt int64
	value interface{}
}

type cacheData map[string]cacheEntry

type Cache struct {
	lock sync.RWMutex
	data cacheData
	done chan struct{}
}

func New(cleanupFrequency time.Duration) *Cache {
	c := &Cache{
		data: make(cacheData),
		done: make(chan struct{}),
	}
	if cleanupFrequency > 0 {
		go cleaner(c, cleanupFrequency)
	}
	return c
}

func cleaner(c *Cache, frequency time.Duration) {
	ticker := time.NewTicker(frequency)
	for {
		select {
		// Terminate the goroutine if the cache has been closed.
		case <-c.done:
			return
		case <-ticker.C:
			c.lock.Lock()
			for k, v := range c.data {
				if v.expiresAt == 0 {
					continue
				}
				if time.Now().UnixNano() > v.expiresAt {
					delete(c.data, k)
				}
			}
			c.lock.Unlock()
		}
	}
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	var expiresAt int64
	if duration != indefinite {
		expiresAt = time.Now().Add(duration).UnixNano()
	}

	c.data[key] = cacheEntry{
		expiresAt: expiresAt,
		value:    value,
	}
}

func (c *Cache) SetIndefinite(key string, value interface{}) {
	c.Set(key, value, indefinite)
}

func (c *Cache) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, key)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	entry, ok := c.data[key]
	return entry.value, ok
}

func (c *Cache) Close() {
	close(c.done)
}