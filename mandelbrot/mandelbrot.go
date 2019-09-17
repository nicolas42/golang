package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {

	img := drawMandelbrot(1000, 1000)

	var buffer bytes.Buffer
	png.Encode(&buffer, img)

	f, err := os.Create("image.png")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	f.Write(buffer.Bytes())

}

func testOsCreate() {
	img := drawMandelbrot(1000, 1000)

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func drawMandelbrot(width int, height int) image.Image {
	var x, y int
	var xmin, xmax, ymin, ymax float64
	var xscale, yscale, zx, zy, cx, cy, zxtemp float64
	var i uint32
	var isInSet bool

	// const width = 1000 // pixels
	// const height = 1000
	// const stride = 4
	const IMAX = 40

	xmin = -2
	xmax = 2
	ymin = -2
	ymax = 2

	xscale = float64(width) / (xmax - xmin)
	yscale = float64(height) / (ymax - ymin)

	black := [4]uint8{0, 0, 0, 255}
	white := [4]uint8{255, 255, 255, 255}
	// color := [4]uint8{255, 255, 255, 255}

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	// Work out the other pixels
	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {

			// Work out whether pixel is in set
			// image is flipped vertically. numbers go up, pixels go down
			cx = float64(x)/xscale - xmax
			cy = float64(y)/yscale - ymax
			zx = 0
			zy = 0
			isInSet = true

			for i = 1; i <= IMAX; i++ {
				zxtemp = zx*zx - zy*zy + cx
				zy = 2*zx*zy + cy
				zx = zxtemp

				if zx*zx+zy*zy > 4 {
					isInSet = false
					break
				}
			}

			if isInSet {

				img.Set(x, y, color.NRGBA{
					R: black[0],
					G: black[1],
					B: black[2],
					A: 255,
				})
				// print("*")

			} else {

				img.Set(x, y, color.NRGBA{
					R: uint8(uint32(white[0]) * i / IMAX),
					G: uint8(uint32(white[1]) * i / IMAX),
					B: uint8(uint32(white[2]) * i / IMAX),
					A: 255,
				})
				// print(" ")
			}
		}
		// print("\n")
	}
	return img
}
