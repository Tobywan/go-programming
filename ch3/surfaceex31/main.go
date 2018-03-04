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

type point struct {
	x, y, z float64
}

// plot returns a canvas point. This is a 2d plot so ignore the z component
func (p *point) plot() (plot point) {
	// isometric projection onto 2d canvas
	plot.x = width/2 + (p.x-p.y)*cosAng*xyscale
	plot.y = height/2 + (p.x+p.y)*sinAng*xyscale - p.z*zscale
	return plot
}

type poly struct {
	nw, ne, sw, se point
}

// plot returns a canvas poly. This is a 2d plot so ignore the z component
func (p *poly) plot() (plot poly) {
	plot.nw = p.nw.plot()
	plot.ne = p.ne.plot()
	plot.se = p.se.plot()
	plot.sw = p.sw.plot()
	return plot
}

// lowest finds the lowest spot in a 3d polygon
func (p *poly) lowest() float64 {
	return math.Min(math.Min(p.nw.z, p.ne.z),
		math.Min(p.sw.z, p.se.z))
}

// highest finds the highest spot in a 3d polygon
func (p *poly) highest() float64 {
	return math.Max(math.Max(p.nw.z, p.ne.z),
		math.Max(p.sw.z, p.se.z))
}

func (p *poly) meanZ() float64 {
	return (p.nw.z + p.ne.z + p.se.z + p.sw.z) / 4
}

func (p *poly) toSVG(zRange float64) string {
	pp := p.plot()
	return fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
		pp.nw.x, pp.nw.y, pp.ne.x, pp.ne.y, pp.se.x, pp.se.y, pp.sw.x, pp.sw.y)
}

// loop over points
//		fmt.Printf
//			ax, ay, bx, by, cx, cy, dx, dy)
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	polys, minZ, maxZ := getCorners()

	// find height range
	delta := maxZ - minZ

	for _, p := range polys {
		fmt.Println(p.toSVG(delta))
	}

	fmt.Println("</svg>")
}

func getCorners() (polys []poly, minZ float64, maxZ float64) {

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			nw := corner(i, j)
			ne := corner(i+1, j)
			se := corner(i+1, j+1)
			sw := corner(i, j+1)
			p := poly{nw, ne, se, sw}
			minZ = math.Min(p.meanZ(), minZ)
			maxZ = math.Max(p.meanZ(), maxZ)
			polys = append(polys, p)
		}
	}
	return polys, minZ, maxZ
}

func corner(i, j int) point {
	// Find the point at corner of cell
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := curvy(x, y, 3)
	return point{x, y, z}
}

func curvy(x, y, scale float64) float64 {
	r := scale * math.Hypot(x, y)
	return math.Sin(r) / r
}

func zero(x, y float64) float64 {
	return 0
}
