package main

import (
	"os"
	"strconv"

	"github.com/tobywan/go-programming/ch3/mandlebrotex35"
)

// main takes 4 arguments, the first two are the centre of the
// argnd, and the last two are height and width
// e.g.
// $ ./cmd 0 0 4 4 > ~/junk/wendy.png
func main() {
	var reZ, imZ, h, w float64

	reZ, _ = strconv.ParseFloat(os.Args[1], 64)
	imZ, _ = strconv.ParseFloat(os.Args[2], 64)
	h, _ = strconv.ParseFloat(os.Args[3], 64)
	w, _ = strconv.ParseFloat(os.Args[4], 64)

	c := mandlebrotex35.NewCanvas(1024, 1024)
	a := mandlebrotex35.NewArgand(complex(reZ, imZ), h, w)

	c.PlotMandelbrot(a)
	c.PNG(os.Stdout)

}
