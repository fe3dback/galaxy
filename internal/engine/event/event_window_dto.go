package event

import "github.com/veandco/go-sdl2/sdl"

const (
	WindowEventTypeResize      WindowEventType = sdl.WINDOWEVENT_RESIZED
	WindowEventTypeSizeChanged WindowEventType = sdl.WINDOWEVENT_SIZE_CHANGED
)

type (
	WindowEventType = uint8

	WindowEvent struct {
		WindowID  uint32
		EventType WindowEventType
		Data1     int32
		Data2     int32
	}
)
