package simplecache

import (
	"sync"
	"time"
)

const (
	indefinite time.Duration = -1
	NeverCleanup time.Duration = -1
)

// cacheEntry is an individual record within the HashCache.
type cacheEntry struct {
	expiresAt int64
	value     interface{}
}

// cacheData is the data store type which backs the HashCache.
type cacheData map[string]cacheEntry

// HashCache is a concurrency safe in-memory cache, featuring time duration based entry lifetime and expiry.
type HashCache struct {
	mu   sync.RWMutex
	data cacheData
	done chan struct{}
}

// NewHashCache creates a new hashCache and returns a pointer to it.  If a 'cleanupFrequency' > 0 is provided, the cleaner goroutine
// is launched (which removes expired values from the cache).
func NewHashCache(cleanupFrequency time.Duration) *HashCache {
	c := &HashCache{
		data: make(cacheData),
		done: make(chan struct{}),
	}
	if cleanupFrequency > 0 {
		go cleaner(c, cleanupFrequency)
	}
	return c
}

// cleaner removes expired values from the cache.
func cleaner(c *HashCache, frequency time.Duration) {
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
func (c *HashCache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var expiresAt int64
	if duration != indefinite {
		expiresAt = time.Now().Add(duration).UnixNano()
	}

	c.data[key] = cacheEntry{
		expiresAt: expiresAt,
		value:     value,
	}
}

// SetIndefinite sets a value in the cache which will not expire.
func (c *HashCache) SetIndefinite(key string, value interface{}) {
	c.Set(key, value, indefinite)
}

// Delete removes a value from the cache.
func (c *HashCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Get retrieves a value from the cache, using the boolean value to indicate the presence or non-presence of the value.
func (c *HashCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.data[key]
	return entry.value, ok
}

// Close terminates the cleaner goroutine, resolving the potential for a goroutine leak.
func (c *HashCache) Close() {
	close(c.done)
}

// Len provides the number of items stored in the cache.
func (c *HashCache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}
