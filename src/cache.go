// Package simplecache contains a concurrency safe, in-memory cache for the low latency retrieval of values which may be
// computationally intensive to generate or find.
package simplecache

import (
	"time"
)

type Cache interface {
	Set(key string, value interface{}, duration time.Duration)
	SetIndefinite(key string, value interface{})
	Delete(key string)
	Get(key string) (interface{}, bool)
	Close()
	Len() int
}
