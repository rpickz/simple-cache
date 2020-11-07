package cache

import (
	"sync"
	"time"
)

const (
	indefinite time.Duration = -1
)

// cacheEntry is an individual record within the Cache.
type cacheEntry struct {
	expiresAt int64
	value interface{}
}

// cacheData is the data store type which backs the Cache.
type cacheData map[string]cacheEntry

// Cache is a concurrency safe in-memory cache, featuring time duration based entry lifetime and expiry.
type Cache struct {
	mu   sync.RWMutex
	data cacheData
	done chan struct{}
}

// New creates a new Cache and returns a pointer to it.  If a 'cleanupFrequency' > 0 is provided, the cleaner goroutine
// is launched (which removes expired values from the cache).
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

// cleaner removes expired values from the cache.
func cleaner(c *Cache, frequency time.Duration) {
	ticker := time.NewTicker(frequency)
	for {
		select {
		// Terminate the goroutine if the cache has been closed.
		case <-c.done:
			return
		case <-ticker.C:
			c.mu.Lock()
			for k, v := range c.data {
				if v.expiresAt == 0 {
					continue
				}
				if time.Now().UnixNano() > v.expiresAt {
					delete(c.data, k)
				}
			}
			c.mu.Unlock()
		}
	}
}

// Set sets a value in the cache, which is removed when the 'duration' is exceeded.
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var expiresAt int64
	if duration != indefinite {
		expiresAt = time.Now().Add(duration).UnixNano()
	}

	c.data[key] = cacheEntry{
		expiresAt: expiresAt,
		value:    value,
	}
}

// SetIndefinite sets a value in the cache which will not expire.
func (c *Cache) SetIndefinite(key string, value interface{}) {
	c.Set(key, value, indefinite)
}

// Delete removes a value from the cache.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Get retrieves a value from the cache, using the boolean value to indicate the presence or non-presence of the value.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.data[key]
	return entry.value, ok
}

// Close terminates the cleaner goroutine, resolving the potential for a goroutine leak.
func (c *Cache) Close() {
	close(c.done)
}
