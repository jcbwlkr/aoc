package main

import (
	"image"
	"image/color"
	"image/color/palette"
)

// r is used by Board for creating images
// Lights have some padding and take 4 pixels each plus some extra padding on
// the right and bottom. So 2 rows of 2 lights would look like:
//
// yx01234
// 0 .....
// 1 .*.*.
// 2 .....
// 3 .*.*.
// 4 .....
//
// Therefore 1000 lights across will be 2000 pixels wide plus 1 pixel for the
// last column's right padding. Similarly 1000 lights tall will be 2000 pixels
// high plus 1 pixel for the last row's bottom padding.
var r = image.Rect(0, 0, 2001, 2001)

// Board is a collection of lights
type Board map[Point]*Light

// NewBoard preps a 1000x1000 board of lights
func NewBoard() Board {
	b := Board{}
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			b[Point{X: x, Y: y}] = NewLight()
		}
	}

	return b
}

// ImageBW expresses the Board's current state as black and white image.
func (b Board) ImageBW() *image.Paletted {
	img := image.NewPaletted(r, palette.Plan9)
	for p, l := range b {
		if l.On {
			img.Set(p.X*2+1, p.Y*2+1, color.White)
		}
	}

	return img
}

// ImageColor expresses the Board's current state as colored image.
func (b Board) ImageColor() *image.Paletted {
	img := image.NewPaletted(r, palette.Plan9)
	for p, l := range b {
		img.Set(p.X*2+1, p.Y*2+1, l.Color())
	}

	return img
}

// LitCount reports how many lights on the board are lit according to Part 1
// rules.
func (b Board) LitCount() int {
	count := 0
	for _, l := range b {
		if l.On {
			count++
		}
	}

	return count
}

// Brightness gives the total brightness of all lights on the board.
func (b Board) Brightness() int {
	brightness := 0
	for _, l := range b {
		brightness += l.Brightness
	}

	return brightness
}
