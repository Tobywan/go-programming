// Server2 is a minimal echo server

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex
var count int

func main() {

	port := os.Args[1]
	if port == "" {
		port = "8000"
	}

	hostport := fmt.Sprintf("localhost:%s", port)
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(hostport, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Request count:%d\n", count)
	mu.Unlock()
}
