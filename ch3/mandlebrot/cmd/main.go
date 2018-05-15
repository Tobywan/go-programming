package main

import (
	"os"
	"strconv"

	"github.com/tobywan/go-programming/ch3/mandlebrot"
)

func main() {
	var reZ, imZ, h, w float64

	reZ, _ = strconv.ParseFloat(os.Args[1], 64)
	imZ, _ = strconv.ParseFloat(os.Args[2], 64)
	h, _ = strconv.ParseFloat(os.Args[3], 64)
	w, _ = strconv.ParseFloat(os.Args[4], 64)

	c := mandlebrot.NewCanvas(1024, 1024)
	a := mandlebrot.NewArgand(complex(reZ, imZ), h, w)

	c.PlotMandelbrotChan(a, 8)
	c.PNG(os.Stdout)

}
