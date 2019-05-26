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
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, julia(z))
		}
	}
	if err := png.Encode(os.Stdout, img); err != nil {
		log.Fatal(err)
	}
}

func julia(z complex128) color.Color {
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
	return color.Black
}
