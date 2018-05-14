package main

import (
	"fmt"
	"github.com/Tobywan/go-programming/ch3/surfaceex31"
	"net/http"
	"net/url"
	// "os"
)

// will listen for web regests and then forward them to the relevant downstream handler

type Form url.Values

type FormHandler interface {
}

func main() {
	surfaceex31.Listen("8080")
	//surfaceex31.DrawPlot(os.Stdout, 1000, 600, 30, 30, 0.2, 1, "eggbox")
}

// Listen on a port for surface plot requests
func Listen(port string) {

	hostport := fmt.Sprintf("localhost:%s", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(hostport, nil))
}
