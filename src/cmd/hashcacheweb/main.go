package main

import (
	"log"
	"net/http"
	"os"
	"simplecache"
	"simplecache/web"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	c := simplecache.NewHashCache(time.Minute)
	th := web.NewTransactionalHandlerFunc(logger, c)

	http.HandleFunc("/", th)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error from HTTP listenAndServe: %v", err)
	}
}
