package main

// Action is what to do to a board.
type Action string

// Possible values for Action.
const (
	TurnOn  Action = "turn on"
	TurnOff        = "turn off"
	Toggle         = "toggle"
)
