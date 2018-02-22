// Server1 is a minimal echo server

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Args[1]
	if port == "" {
		port = "8000"
	}

	hostport := fmt.Sprintf("localhost:%s", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(hostport, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
