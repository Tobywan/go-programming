// Mandelbrot emits a PNG of the mandlebrot set

package mandlebrotex35

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

type colour color.RGBA

func (c colour) rgba() color.RGBA {
	return color.RGBA{R: c.R, G: c.G, B: c.B, A: c.A}
}

type column []colour
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
	col colour
}

func (c *Canvas) pixel(x, y int) colour {
	return c.pixels[x][y]
}

func (c *Canvas) setColor(x, y int, col colour) {

	c.pixels[x][y] = col
}

func (c *Canvas) setPixel(p pixel) {

	c.pixels[p.x][p.y] = p.col
}

var tokens chan struct{}

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
			img.SetRGBA(x, y, c.pixel(x, y).rgba())
		}
	}
	png.Encode(w, img)
}

// NewCanvas creates  new canvas for drawing on
func NewCanvas(height, width int) *Canvas {
	c := Canvas{height, width, nil}
	c.pixels = make(columns, width)
	for i := 0; i < width; i++ {
		c.pixels[i] = make([]colour, height)
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

func mandlebrot(z complex128) colour {
	const iterations = 255

	var v complex128
	for n := uint8(0); n <= iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colourize(n)
		}

	}
	return colour{R: 0, G: 0, B: 0, A: 255}
}

// colourize generates a colour for a given uint8
// the higher the number, the darker the colour
// 0 -> white, 255 -> black
func colourize(n uint8) colour {
	//GRB
	//
	var r, g, b, a uint8
	x := ^n                             // invert bits
	g = ((x & (3 << 6)) >> 6) / 3 * 255 // 2 highest bits
	r = ((x & (7 << 3)) >> 3) / 7 * 255 // 3 next highest bits
	b = (x & 7) / 7 * 255               // 3 lowest bits
	a = 255                             // opaque
	return colour{R: r, G: g, B: b, A: a}
}
