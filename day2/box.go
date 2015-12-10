package main

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

// RequiredPaper implements the algorith for calculating the required square
// footage of wrapping paper to wrap a Box.
func (b Box) RequiredPaper() int {
	var (
		top  = b.Length * b.Width
		side = b.Length * b.Height
		end  = b.Width * b.Height
	)

	return 2*top + 2*side + 2*end + smallest(top, side, end)
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
