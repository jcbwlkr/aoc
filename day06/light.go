package main

import "image/color"

// Light is a single christmas light
type Light struct {
	On         bool
	Brightness int
	color      int
}

// Mapping ints to general colors
const (
	Red = iota + 1
	Green
	Blue
	Max
)

var (
	lightCount int
	currColor  = Red
)

// NewLight makes a new light with a color based on creation order.
func NewLight() *Light {
	lightCount++
	if lightCount%10000 == 0 {
		currColor++
	}
	if currColor == Max {
		currColor = Red
	}

	return &Light{
		color: currColor,
	}
}

// TakeAction adjusts a light's On and Brightness values based on an Action.
func (l *Light) TakeAction(a Action) {
	switch a {
	case TurnOn:
		l.On = true
		l.Brightness++
	case TurnOff:
		l.On = false
		if l.Brightness > 0 {
			l.Brightness--
		}
	case Toggle:
		l.On = !l.On
		l.Brightness += 2
	}
}

// Color returns this light's color with the appropriate brightness.
func (l *Light) Color() color.RGBA {
	if l.Brightness == 0 {
		// Black
		return color.RGBA{0x00, 0x00, 0x00, 0xff} // Maybe a turned off light should be kind of grey?
	}

	var val byte
	switch l.Brightness {
	case 1:
		val = 0x66
	case 2:
		val = 0xaa
	default:
		val = 0xff
	}

	switch l.color {
	case Red:
		return color.RGBA{val, 0x00, 0x00, 0xff}
	case Green:
		return color.RGBA{0x00, val, 0x00, 0xff}
	case Blue:
		return color.RGBA{0x00, 0x00, val, 0xff}
	}

	return color.RGBA{0xff, 0xff, 0xff, 0xff}
}
