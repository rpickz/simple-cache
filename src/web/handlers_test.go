package web

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"simplecache"
	"testing"
	"time"
)

func TestTransactionalHandlerFunc(t *testing.T) {
	tests := []struct {
		name                string
		initialCacheData    map[string]interface{}
		expectedCacheData   map[string]interface{}
		expectedCacheLength int
		request             *http.Request
		responseCode        int
	}{
		{
			name:         "GetOnEmptyCacheProvidesNothing",
			request:      httptest.NewRequest(http.MethodGet, "/abc123", nil),
			responseCode: http.StatusNotFound,
		},
		{
			name:                "GetReturnsResult",
			initialCacheData:    map[string]interface{}{"/abc123": []byte("something")},
			request:             httptest.NewRequest(http.MethodGet, "/abc123", nil),
			responseCode:        http.StatusOK,
			expectedCacheLength: 1,
		},
		{
			name:                "PutSetsValue",
			expectedCacheData:   map[string]interface{}{"/abc123": []byte("something")},
			request:             httptest.NewRequest(http.MethodPut, "/abc123", bytes.NewBufferString("something")),
			responseCode:        http.StatusOK,
			expectedCacheLength: 1,
		},
		{
			name:                "DeleteRemovesValue",
			initialCacheData:    map[string]interface{}{"/abc123": []byte("something")},
			request:             httptest.NewRequest(http.MethodDelete, "/abc123", nil),
			responseCode:        http.StatusOK,
			expectedCacheLength: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			c := simplecache.NewHashCache(time.Minute)
			for k, v := range test.initialCacheData {
				c.SetIndefinite(k, v)
			}

			logger := log.New(os.Stdout, "", 0)

			handlerFunc := NewTransactionalHandlerFunc(logger, c)

			// Perform a transaction...
			response := httptest.NewRecorder()
			handlerFunc(response, test.request)

			// Assert on the result...
			if response.Code != test.responseCode {
				t.Errorf(`wanted %v - got %v`, test.responseCode, response.Code)
			}

			// Cache is in desired state...
			for k, expectedVal := range test.expectedCacheData {
				result, ok := c.Get(k)
				if !ok {
					t.Errorf(`wanted %v - could not find value in cache`, result)
					continue
				}
				if !reflect.DeepEqual(result, expectedVal) {
					t.Errorf(`wanted %q - got %q`, expectedVal, result)
				}
			}

			if cLen := c.Len(); cLen != test.expectedCacheLength {
				t.Errorf(`wanted len %d - got %d`, test.expectedCacheLength, cLen)
			}
		})
	}
}

func BenchmarkTransactionalHandlerFunc(b *testing.B) {
	benchmarks := []struct {
		name    string
		request *http.Request
	}{
		{
			name:    "Get",
			request: httptest.NewRequest(http.MethodGet, "/abc123", nil),
		},
		{
			name:    "Put",
			request: httptest.NewRequest(http.MethodPut, "/abc123", bytes.NewBufferString("something")),
		},
		{
			name:    "Delete",
			request: httptest.NewRequest(http.MethodDelete, "/abc123", nil),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {

			c := simplecache.NewHashCache(simplecache.NeverCleanup)
			c.SetIndefinite("/abc123", "something")

			logger := log.New(os.Stdout, "", 0)
			handlerFunc := NewTransactionalHandlerFunc(logger, c)
			response := httptest.NewRecorder()

			for i := 0; i < b.N; i++ {
				handlerFunc(response, bm.request)
			}
		})
	}
}
