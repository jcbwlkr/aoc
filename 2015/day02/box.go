package main

import (
	"errors"
	"strconv"
	"strings"
)

// Box represents a box that needs wrapped
type Box struct {
	Length int
	Width  int
	Height int
}

// NewBox constructs a Box
func NewBox(l, w, h int) Box {
	return Box{
		Length: l,
		Width:  w,
		Height: h,
	}
}

// These errors may be returned when calling NewBoxFromDimensions with invalid
// input
var (
	ErrDimensionsMalformed = errors.New("dimension string is not of the form NxNxN")
	ErrDimensionsNonInt    = errors.New("dimensions are not all ints")
)

// NewBoxFromDimensions creates a box from a string of the form "1x2x3" meaning
// length 1, width 2, and height 3.
func NewBoxFromDimensions(dimensions string) (Box, error) {
	parts := strings.Split(dimensions, "x")
	if len(parts) != 3 {
		return Box{}, ErrDimensionsMalformed
	}
	dims := make([]int, 3)
	for i, v := range parts {
		d, err := strconv.Atoi(v)
		if err != nil {
			return Box{}, ErrDimensionsNonInt
		}
		dims[i] = d
	}

	return NewBox(dims[0], dims[1], dims[2]), nil
}

// RequiredPaper calculates the required square footage of wrapping paper to
// wrap a Box.
func (b Box) RequiredPaper() int {
	// areas of the faces
	var (
		top  = b.Length * b.Width
		side = b.Length * b.Height
		end  = b.Width * b.Height
	)

	return 2*top + 2*side + 2*end + smallest(top, side, end)
}

// RequiredRibbon calculates the required length of ribbon to dress a Box
func (b Box) RequiredRibbon() int {
	// perimeters of the faces
	var (
		top  = 2*b.Length + 2*b.Width
		side = 2*b.Length + 2*b.Height
		end  = 2*b.Width + 2*b.Height
	)

	volume := b.Length * b.Width * b.Height

	return smallest(top, side, end) + volume
}

func smallest(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	x := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < x {
			x = nums[i]
		}
	}
	return x
}
