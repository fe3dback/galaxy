package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
	event2 "github.com/fe3dback/galaxy/internal/engine/event"
)

type (
	Camera struct {
		position galx.Vec
		width    int
		height   int
		zoom     float64

		dispatcher *event2.Dispatcher
		queued     queued
	}

	queued struct {
		width  int
		height int
		zoom   float64
	}
)

func NewCamera(dispatcher *event2.Dispatcher) *Camera {
	cam := &Camera{
		dispatcher: dispatcher,

		position: galx.Vec{},
		width:    320,
		height:   240,
		zoom:     1,
	}

	cam.queued = queued{
		width:  cam.width,
		height: cam.height,
		zoom:   cam.zoom,
	}

	dispatcher.OnFrameEnd(cam.onFrameEnd)
	return cam
}

func (c *Camera) Position() galx.Vec {
	return c.position
}

func (c *Camera) Screen2World(screen galx.Vec) galx.Vec {
	return screen.Add(c.position)
}

func (c *Camera) World2Screen(world galx.Vec) galx.Vec {
	return world.Sub(c.position)
}

func (c *Camera) Width() int {
	return c.width
}

func (c *Camera) Height() int {
	return c.height
}

func (c *Camera) Zoom() float64 {
	return c.zoom
}

func (c *Camera) MoveTo(p galx.Vec) {
	c.position = p
}

func (c *Camera) CenterOn(p galx.Vec) {
	c.MoveTo(galx.Vec{
		X: p.X - (float64(c.width)/c.zoom)/2,
		Y: p.Y - (float64(c.height)/c.zoom)/2,
	})
}

func (c *Camera) Resize(width, height int) {
	if width < 1 || height < 1 {
		panic(fmt.Sprintf("can`t resize camera to %d x %d", width, height))
	}

	c.queued.width = width
	c.queued.height = height
}

func (c *Camera) ZoomView(scale float64) {
	c.queued.zoom = galx.RoundTo(
		galx.Clamp(scale, 0.25, 10),
	)
}

func (c *Camera) center() galx.Vec {
	return c.position.Add(galx.Vec{
		X: (float64(c.width) / c.zoom) / 2,
		Y: (float64(c.height) / c.zoom) / 2,
	})
}

func (c *Camera) dispatchUpdate() {
	c.dispatcher.PublishEventCameraUpdate(event2.CameraUpdateEvent{
		Width:  c.Width(),
		Height: c.Height(),
		Zoom:   c.Zoom(),
	})
}

func (c *Camera) onFrameEnd(_ event2.FrameEndEvent) error {
	updated := c.applyNewZoom() || c.applyNewSize()

	if updated {
		// update renderer properties
		c.dispatchUpdate()
	}

	return nil
}

func (c *Camera) applyNewSize() bool {
	if c.queued.width == c.width && c.queued.height == c.height {
		return false
	}

	c.autoCorrectCenter(func() {
		c.width = c.queued.width
		c.height = c.queued.height
	})

	return true
}

func (c *Camera) applyNewZoom() bool {
	if c.queued.zoom == c.zoom {
		return false
	}

	c.autoCorrectCenter(func() {
		c.zoom = c.queued.zoom
	})

	// applied
	return true
}

func (c *Camera) autoCorrectCenter(operation func()) {
	oldCenter := c.center()
	operation()
	newCenter := c.center()
	correctOffsetDiff := oldCenter.Sub(newCenter)

	// move camera to new center
	c.MoveTo(c.position.Add(correctOffsetDiff))
}
