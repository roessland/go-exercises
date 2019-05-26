package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	width, height := 2560, 1440
	aspect := float64(width) / float64(height)
	ymin, ymax := -1.0, 1.0
	xmin, xmax := aspect*ymin, aspect*ymax

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y0 := (float64(py)-0.25)/float64(height)*(ymax-ymin) + ymin
		y1 := (float64(py)+0.25)/float64(height)*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x0 := (float64(px)-0.25)/float64(width)*(xmax-xmin) + xmin
			x1 := (float64(px)+0.25)/float64(width)*(xmax-xmin) + xmin

			j00 := julia(complex(x0, y0))
			j01 := julia(complex(x0, y1))
			j10 := julia(complex(x1, y0))
			j11 := julia(complex(x1, y1))

			img.Set(px, py, avg(j00, j01, j10, j11))
		}
	}
	if err := png.Encode(os.Stdout, img); err != nil {
		log.Fatal(err)
	}
}

func avg(a, b, c, d color.RGBA) color.RGBA {
	return color.RGBA{
		a.R/4 + b.R/4 + c.R/4 + d.R/4,
		a.G/4 + b.G/4 + c.G/4 + d.G/4,
		a.B/4 + b.B/4 + c.B/4 + d.B/4,
		0xFF,
	}
}

func julia(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	c := -0.835 - 0.2321i
	for n := uint8(0); n < iterations; n++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			idx := contrast * n
			return viridis[idx]
		}
		z = z*z + c
	}
	return viridis[0]
}
