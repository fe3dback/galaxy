package event

type eventType = uint8

//go:generate go run ../../../cmd/event_generator/main.go -path="$PWD/$GOFILE"

const (
	eventTypeQuit eventType = iota
	eventTypeKeyBoard
	eventTypeWindowResized
	eventTypeMouseButton
	eventTypeMouseWheel
	eventTypeMouseMove
	eventTypeFrameStart
	eventTypeFrameEnd
	eventTypeCameraUpdate
)
