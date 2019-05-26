package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells         = 60
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		plot(w)
	})
	log.Print("Listening on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func plot(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, colA := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			if math.IsInf(ax+ay+bx+by+cx+cy+dx+dy, 0) {
				continue // skip infinity coordinate
			}
			fmt.Fprintf(out, "<polygon fill='%s' points='%.4g,%.4g %.4g,%.4g %.4g,%.4g %.4g,%.4g'/>\n",
				colA, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, string) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	color := getColor(z)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color
}

func getColor(z float64) string {
	clamp := 0.1
	if z > clamp {
		z = clamp
	} else if z < -clamp {
		z = -clamp
	}
	sat := math.Abs(z / clamp)
	b := uint8(sat * 255)
	if z > 0 {
		return fmt.Sprintf("#%02x%02x%02x", b, 0x00, 0x00)
	} else {
		return fmt.Sprintf("#%02x%02x%02x", 0x00F, 0x00, b)
	}

}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggShape(x, y float64) float64 {
	r2 := x*x + 2.5*y*y
	if r2 > 100.0 {
		return 0.0
	}
	return math.Sqrt(100-r2) / 20
}

func sincInf(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r > 15 {
		return math.Inf(+1)
	}
	return math.Sin(r) / r
}
