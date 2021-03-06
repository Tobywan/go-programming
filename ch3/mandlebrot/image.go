// Mandelbrot emits a PNG of the mandlebrot set

package mandlebrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"sync"
)

type column []color.Color
type columns []column

// Canvas represents an image rectangle
type Canvas struct {
	height int
	width  int
	pixels columns
}

type pixel struct {
	x   int
	y   int
	col color.Color
}

func (c *Canvas) pixel(x, y int) color.Color {
	return c.pixels[x][y]
}

func (c *Canvas) setColor(x, y int, col color.Color) {

	c.pixels[x][y] = col
}

func (c *Canvas) setPixel(p pixel) {

	c.pixels[p.x][p.y] = p.col
}

var tokens chan struct{}

// PlotMandelbrotChan plots those points in the passed in angand diagram that lie within
// the mandelbrot set using nChans  concurrent channels
// NO IMPROVEMENT WHEN USING CHANNELS!
func (c *Canvas) PlotMandelbrotChan(a *Argand, nChans int) {
	pixels := make(chan pixel)
	tokens = make(chan struct{}, nChans)

	var wg sync.WaitGroup // number of go routines in progress
	for x := 0; x < c.width; x++ {
		for y := 0; y < c.height; y++ {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				tokens <- struct{}{} // acquire token
				z := mapComplex(x, y, *c, *a)
				col := mandlebrot(z)
				<-tokens // release it
				pixels <- pixel{x, y, col}
			}(x, y)

		}
	}
	// closer
	go func() {
		wg.Wait()
		close(pixels)
	}()

	for p := range pixels {
		c.setPixel(p)
	}

}

// PlotMandelbrot plots those points in the passed in angand diagram that lie within
// the mandelbrot set
func (c *Canvas) PlotMandelbrot(a *Argand) {
	for x := 0; x < c.width; x++ {
		for y := 0; y < c.height; y++ {
			z := mapComplex(x, y, *c, *a)
			col := mandlebrot(z)
			c.setColor(x, y, col)
		}
	}
}

// PNG exports the canvas to PNG
func (c *Canvas) PNG(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, c.width, c.height))
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			img.Set(x, y, c.pixel(x, y))
		}
	}
	png.Encode(w, img)
}

// NewCanvas creates  new canvas for drawing on
func NewCanvas(height, width int) *Canvas {
	c := Canvas{height, width, nil}
	c.pixels = make(columns, width)
	for i := 0; i < width; i++ {
		c.pixels[i] = make([]color.Color, height)
	}
	return &c
}

// Argand is the bounds of an Argand diagram
type Argand struct {
	minReal float64
	maxReal float64
	minImag float64
	maxImag float64
	centre  complex128
}

// NewArgand creates a new argand diagram area
func NewArgand(centre complex128, height float64, width float64) *Argand {
	a := Argand{
		minReal: real(centre) - width/2.0,
		maxReal: real(centre) + width/2.0,
		minImag: imag(centre) - height/2.0,
		maxImag: imag(centre) + height/2.0,
		centre:  centre,
	}
	return &a
}

// point is the point on a Canvas. (0,0 is top left)

// mapCompex converts a Canvas point to its corresponding
// complex number within the Argand diagram.
func mapComplex(x, y int, c Canvas, a Argand) complex128 {
	// calculate with origin at centre of Canvas
	cx := x - c.width/2
	cy := c.height/2 - y // invert y axis

	scalex := (a.maxReal - a.minReal) / float64(c.width)
	scaley := (a.maxImag - a.minImag) / float64(c.height)

	return complex((float64(cx)*scalex)+real(a.centre), (float64(cy)*scaley)+imag(a.centre))

}

func mandlebrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}

	}
	return color.Black
}
