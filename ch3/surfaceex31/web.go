package surfaceex31

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Listen on a port for surface plot requests
func Listen(port string) {

	hostport := fmt.Sprintf("localhost:%s", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(hostport, nil))
}

// expects a url like /?w=800&h=400&cells=20&xy=30&zs=0.5&ps=1&f=saddle
func handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	/*
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	*/

	w.Header().Set("Content-Type", "image/svg+xml")

	f := r.Form
	DrawPlot(w,
		getInt(f, "w"),
		getInt(f, "h"),
		getInt(f, "cells"),
		getFloat(f, "xy"),
		getFloat(f, "zs"),
		getFloat(f, "ps"),
		f["f"][0])

}

func getInt(v url.Values, s string) int {
	r, _ := strconv.ParseInt(v[s][0], 10, 0)
	return int(r)
}
func getFloat(v url.Values, s string) float64 {
	r, _ := strconv.ParseFloat(v[s][0], 64)
	return float64(r)
}
