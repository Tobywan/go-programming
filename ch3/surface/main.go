// Surface computes an SVG rendering of a 3-D surface function
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis reanges ( -xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y
	zscale        = height * 0.1
	angle         = math.Pi / 5
)

var sinAng, cosAng = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find the point at corner of cell
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := curvy(x, y, 3)

	// isometric projection onto 2d canvas
	sx := width/2 + (x-y)*cosAng*xyscale
	sy := height/2 + (x+y)*sinAng*xyscale - z*zscale
	return sx, sy
}

func curvy(x, y, scale float64) float64 {
	r := scale * math.Hypot(x, y)
	return math.Sin(r) / r
}

func zero(x, y float64) float64 {
	return 0
}
