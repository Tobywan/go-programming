// Surface computes an SVG rendering of a 3-D surface function
package main

import (
	"fmt"
	"github.com/gerow/go-color"
	"math"
	"os"
)

const (
	width, height = 1200, 640           // canvas size
	cells         = 200                 // number of grid cells
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
	nw, ne, se, sw point
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
	// debug(fmt.Sprintf("Meanz:%g, %g, %g, %g\n", p.nw.z, p.ne.z, p.se.z, p.sw.z))
	ret := (p.nw.z + p.ne.z + p.se.z + p.sw.z) / 4
	//debug(fmt.Sprintf("meanZ: %g\n", ret))
	if math.IsNaN(ret) {
		// debug(fmt.Sprintf("Invalid poly: %v", p))
		ret = 0
	}
	return ret
}

// toSVG generates a plotted SVG ploygon where 0<= relheight <=1 is the
// relative hieght of the poligon compared to all the others
func (p *poly) toSVG(relHeight float64) string {
	//debug(fmt.Sprintf("relHeight: %g", relHeight))

	pp := p.plot()
	c := color.HSL{H: (1 - relHeight) * 0.7, S: 1, L: 0.5}.ToHTML()
	return fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
		"style='stroke: grey; fill: #%s; stroke-width: 0.3' "+
		"/>\n",
		pp.nw.x, pp.nw.y, pp.ne.x, pp.ne.y, pp.se.x, pp.se.y, pp.sw.x, pp.sw.y, c)
}

// loop over points
//		fmt.Printf
//			ax, ay, bx, by, cx, cy, dx, dy)
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%d' height='%d'>", width, height)

	polys, minZ, maxZ := getCorners()

	// find height range
	delta := maxZ - minZ

	for _, p := range polys {
		//		debug(fmt.Sprintf("delta: %g", delta))
		relHeight := (p.meanZ() - minZ) / delta
		fmt.Println(p.toSVG(relHeight))
	}

	fmt.Println("</svg>")
}

func debug(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
}

func getCorners() (polys []poly, minZ float64, maxZ float64) {

	debug(fmt.Sprintf("min: %g, max: %g", minZ, maxZ))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			//				ax, ay := corner(i+1, j)
			//				bx, by := corner(i, j)
			//				cx, cy := corner(i, j+1)
			//				dx, dy := corner(i+1, j+1)
			nw := corner(i, j)
			ne := corner(i+1, j)
			se := corner(i+1, j+1)
			sw := corner(i, j+1)
			p := poly{nw, ne, se, sw}
			mean := p.meanZ()
			minZ = math.Min(mean, minZ)
			maxZ = math.Max(mean, maxZ)
			polys = append(polys, p)
		}
	}
	// debug(fmt.Sprintf("min: %g, max: %g", minZ, maxZ))

	return polys, minZ, maxZ
}

func corner(i, j int) point {
	// Find the point at corner of cell
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//z := curvy(x, y, 1.5)
	z := eggbox(x, y, 1.5)
	//z := saddle(x, y, 0.11)
	return point{x, y, z}
}

func eggbox(x, y, scale float64) float64 {
	return (math.Sin(scale*x) + math.Sin(scale*y)) / 4
}

func saddle(x, y, scale float64) float64 {
	return (scale * scale * x * x) - (scale * scale * y * y)
}

func curvy(x, y, scale float64) float64 {
	r := scale * math.Hypot(x, y)
	if r == 0 {
		return 1
	}

	return math.Sin(r) / r
}

func zero(x, y float64) float64 {
	return 0
}
