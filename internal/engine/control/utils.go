package control

type (
	state = uint8
)

const (
	stateUnknown state = iota
	statePressed
	stateReleased
	stateDown
	stateUp
)
