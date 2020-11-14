package main

import (
	"hash/fnv"
	"log"
	"net/http"
	"os"
	"simplecache"
	"time"

	_ "net/http/pprof"

	"github.com/valyala/fasthttp"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	cacheCount := 10
	caches := make([]simplecache.Cache, 0)

	for i := 0; i < cacheCount; i++ {
		c := simplecache.NewHashCache(time.Minute)
		caches = append(caches, c)
	}

	go func() {
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			log.Printf("Error from HTTP listenAndServe: %v", err)
		}
	}()

	err := fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		HandleFastHTTP(logger, caches, ctx)
	})
	if err != nil {
		log.Printf("Error from fasthttp listenAndServe: %v", err)
	}
}

func hash(s []byte) uint32 {
	h := fnv.New32a()
	h.Write(s)
	return h.Sum32()
}

func HandleFastHTTP(logger *log.Logger, caches []simplecache.Cache, ctx *fasthttp.RequestCtx) {

	key := string(ctx.Path())
	partitions := len(caches)
	partition := int(hash(ctx.Path())) % partitions
	c := caches[partition]

	switch string(ctx.Method()) {
	case http.MethodGet:
		// Retrieve the value.
		value, ok := c.Get(key)
		if !ok {
			ctx.Error("Value not found", http.StatusNotFound)
			return
		}

		// Convert to serialisable byte slice.
		val, ok := value.([]byte)
		if !ok {
			ctx.Error("Data format resolution error", http.StatusInternalServerError)
			return
		}

		// Write value back to client.
		_, err := ctx.Write(val)
		if err != nil {
			logger.Printf("Could not send retrieved data back to client - connection error: %v", err)
		}

	case http.MethodPut:

		data := ctx.PostBody()

		expiryStr := string(ctx.Request.Header.Peek("cache-expiry"))
		if expiryStr == "" {
			expiryStr = "0"
		}

		duration, err := time.ParseDuration(expiryStr)
		if err != nil {
			log.Printf("Could not convert cache-expiry header to Duration - error: %v", err)
			ctx.Error("Bad request", http.StatusBadRequest)
			return
		}

		c.Set(key, data, duration)

	case http.MethodDelete:
		c.Delete(key)

	default:
		ctx.Error("Unsupported method", http.StatusMethodNotAllowed)
	}
}
