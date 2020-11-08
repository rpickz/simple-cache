package web

import (
	"io/ioutil"
	"log"
	"net/http"
	"simplecache"
	"time"
)

// NewTransactionalHandlerFunc returns a HTTP handler func for the specified cache to receive and act on transactions sent via
// HTTP.  This follows a simple HTTP API format making use of PUT, GET and DELETE request methods.
func NewTransactionalHandlerFunc(logger *log.Logger, c simplecache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactionalHandlerFunc(logger, c, w, r)
	}
}

// transactionalHandler is a HTTP handler which performs transactions on the provided cache.  This can be used to expose
// HTTP server endpoints dedicated to the cache.
func transactionalHandlerFunc(logger *log.Logger, c simplecache.Cache, w http.ResponseWriter, r *http.Request) {

	key := r.URL.Path

	switch r.Method {
	case http.MethodGet:
		// Retrieve the value.
		value, ok := c.Get(key)
		if !ok {
			http.Error(w, "Value not found", http.StatusNotFound)
			return
		}

		// Convert to serialisable byte slice.
		val, ok := value.([]byte)
		if !ok {
			http.Error(w, "Data format resolution error", http.StatusInternalServerError)
			return
		}

		// Write value back to client.
		_, err := w.Write(val)
		if err != nil {
			logger.Printf("Could not send retrieved data back to client - connection error: %v", err)
		}

	case http.MethodPut:
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Could not read request body - error: %v", err)
			return
		}

		expiryStr := r.Header.Get("cache-expiry")
		if expiryStr == "" {
			expiryStr = "0"
		}

		duration, err := time.ParseDuration(expiryStr)
		if err != nil {
			log.Printf("Could not convert cache-expiry header to Duration - error: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		c.Set(key, data, duration)

	case http.MethodDelete:
		c.Delete(key)

	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
