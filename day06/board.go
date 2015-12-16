package main

import (
	"image"
	"image/color"
	"image/color/palette"
)

// r is used by Board for creating images
var r = image.Rect(0, 0, 1000, 1000)

// Board is a collection of lights
type Board map[Point]*Light

// NewBoard preps a 1000x1000 board of lights
func NewBoard() Board {
	b := Board{}
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			b[Point{X: x, Y: y}] = &Light{}
		}
	}

	return b
}

// Image expresses the Board's current state as an image.
func (b Board) Image() *image.Paletted {
	img := image.NewPaletted(r, palette.Plan9)
	for p, l := range b {
		if l.On {
			img.Set(p.X, p.Y, color.White)
		}
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
