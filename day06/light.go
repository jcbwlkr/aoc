package main

// Light is a single christmas light
type Light struct {
	On         bool
	Brightness int
	Color      int
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
