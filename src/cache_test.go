package cache

import (
	"strconv"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := New(time.Minute)

	// Get should fail as no item is set
	val, ok := cache.Get("abc123")
	if ok {
		t.Errorf(`wanted %v, %v - got %v, %v`, nil, false, val, ok)
	}

	// Set the item...
	cache.Set("abc123", "something", time.Minute * 20)

	// Get should succeed as there is an item now
	val, ok = cache.Get("abc123")
	if !ok {
		t.Errorf(`wanted %v, %v - got %v, %v`, "something", true, val, ok)
	}

	// Delete the item...
	cache.Delete("abc123")

	// Get should fail as no item is set
	val, ok = cache.Get("abc123")
	if ok {
		t.Errorf(`wanted %v, %v - got %v, %v`, nil, false, val, ok)
	}
}

var cacheResponse interface{}

func BenchmarkGet(b *testing.B) {
	benchmarks := []struct {
		name string
		dataSize int
		values int
	}{
		{"Byte16x10", 16, 10},
		{"Byte16x100", 16, 100},
		{"Byte16x1000", 16, 1000},
		{"Byte16x10000", 16, 10000},

		{"Byte128x10", 128, 10},
		{"Byte128x100", 128, 100},
		{"Byte128x1000", 128, 1000},
		{"Byte128x10000", 128, 10000},

		{"Byte1024x10", 1024, 10},
		{"Byte1024x100", 1024, 100},
		{"Byte1024x1000", 1024, 1000},
		{"Byte1024x10000", 1024, 10000},

		{"Byte8192x10", 8192, 10},
		{"Byte8192x100", 8192, 100},
		{"Byte8192x1000", 8192, 1000},
		{"Byte8192x10000", 8192, 10000},

		{"Byte65536x10", 65536, 10},
		{"Byte65536x100", 65536, 100},
		{"Byte65536x1000", 65536, 1000},
		{"Byte65536x10000", 65536, 10000},

		{"Byte524288x10", 524288, 10},
		{"Byte524288x100", 524288, 100},
		{"Byte524288x1000", 524288, 1000},
		{"Byte524288x10000", 524288, 10000},

		{"Byte4194304x10", 4194304, 10},
		{"Byte4194304x100", 4194304, 100},
		{"Byte4194304x1000", 4194304, 1000},
		{"Byte4194304x10000", 4194304, 10000},

		{"Byte33554432x10", 33554432, 10},
		{"Byte33554432x100", 33554432, 100},
		{"Byte33554432x1000", 33554432, 1000},
		{"Byte33554432x10000", 33554432, 10000},

		{"Byte268435456x10", 268435456, 10},
		{"Byte268435456x100", 268435456, 100},
		{"Byte268435456x1000", 268435456, 1000},
		{"Byte268435456x10000", 268435456, 10000},

		{"Byte2147483648x10", 2147483648, 10},
		{"Byte2147483648x100", 2147483648, 100},
		{"Byte2147483648x1000", 2147483648, 1000},
		{"Byte2147483648x10000", 2147483648, 10000},
	}


	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cache := New(time.Second * 30)
			defer cache.Close()
			value := make([]byte, bm.dataSize)
			for i := 0; i < bm.values; i++ {
				cache.Set(strconv.Itoa(i), value, time.Minute * 20)
			}
			cache.Set("abc123", value, time.Minute * 20)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				cache.Get("abc123")
			}
		})
	}
}

func BenchmarkSet(b *testing.B) {
	benchmarks := []struct {
		name string
		dataSize int
		values int
	}{
		{"Byte16x10", 16, 10},
		{"Byte16x100", 16, 100},
		{"Byte16x1000", 16, 1000},
		{"Byte16x10000", 16, 10000},

		{"Byte128x10", 128, 10},
		{"Byte128x100", 128, 100},
		{"Byte128x1000", 128, 1000},
		{"Byte128x10000", 128, 10000},

		{"Byte1024x10", 1024, 10},
		{"Byte1024x100", 1024, 100},
		{"Byte1024x1000", 1024, 1000},
		{"Byte1024x10000", 1024, 10000},

		{"Byte8192x10", 8192, 10},
		{"Byte8192x100", 8192, 100},
		{"Byte8192x1000", 8192, 1000},
		{"Byte8192x10000", 8192, 10000},

		{"Byte65536x10", 65536, 10},
		{"Byte65536x100", 65536, 100},
		{"Byte65536x1000", 65536, 1000},
		{"Byte65536x10000", 65536, 10000},

		{"Byte524288x10", 524288, 10},
		{"Byte524288x100", 524288, 100},
		{"Byte524288x1000", 524288, 1000},
		{"Byte524288x10000", 524288, 10000},

		{"Byte4194304x10", 4194304, 10},
		{"Byte4194304x100", 4194304, 100},
		{"Byte4194304x1000", 4194304, 1000},
		{"Byte4194304x10000", 4194304, 10000},

		{"Byte33554432x10", 33554432, 10},
		{"Byte33554432x100", 33554432, 100},
		{"Byte33554432x1000", 33554432, 1000},
		{"Byte33554432x10000", 33554432, 10000},

		{"Byte268435456x10", 268435456, 10},
		{"Byte268435456x100", 268435456, 100},
		{"Byte268435456x1000", 268435456, 1000},
		{"Byte268435456x10000", 268435456, 10000},

		{"Byte2147483648x10", 2147483648, 10},
		{"Byte2147483648x100", 2147483648, 100},
		{"Byte2147483648x1000", 2147483648, 1000},
		{"Byte2147483648x10000", 2147483648, 10000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cache := New(time.Second * 30)
			defer cache.Close()
			value := make([]byte, bm.dataSize)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				cache.Set(strconv.Itoa(i), value, time.Minute * 20)
			}
		})
	}
}

func BenchmarkDelete(b *testing.B) {
	cache := New(time.Second * 30)
	defer cache.Close()

	for i := 0; i < b.N; i++ {
		cache.Delete("abc123")
	}
}
