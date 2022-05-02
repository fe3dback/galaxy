package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
	event2 "github.com/fe3dback/galaxy/internal/engine/event"
)

type (
	Camera struct {
		position galx.Vec2d
		width    int
		height   int
		scale    float64

		dispatcher *event2.Dispatcher
		queued     queued
	}

	queued struct {
		width  int
		height int
		scale  float64
	}
)

func NewCamera(dispatcher *event2.Dispatcher) *Camera {
	cam := &Camera{
		dispatcher: dispatcher,

		position: galx.Vec2d{},
		width:    320,
		height:   240,
		scale:    1,
	}

	cam.queued = queued{
		width:  cam.width,
		height: cam.height,
		scale:  cam.scale,
	}

	dispatcher.OnFrameEnd(cam.onFrameEnd)
	return cam
}

func (c *Camera) Position() galx.Vec2d {
	return c.position
}

func (c *Camera) Screen2World(screen galx.Vec2d) galx.Vec2d {
	return screen.Decrease(c.scale).Add(c.position)
}

func (c *Camera) World2Screen(world galx.Vec2d) galx.Vec2d {
	return world.Scale(c.scale).Sub(c.position)
}

func (c *Camera) Width() int {
	return c.width
}

func (c *Camera) Height() int {
	return c.height
}

func (c *Camera) Scale() float64 {
	return c.scale
}

func (c *Camera) MoveTo(p galx.Vec2d) {
	c.position = p
}

func (c *Camera) CenterOn(p galx.Vec2d) {
	c.MoveTo(galx.Vec2d{
		X: p.X - (float64(c.width)/c.scale)/2,
		Y: p.Y - (float64(c.height)/c.scale)/2,
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
	c.queued.scale = galx.RoundTo(
		galx.Clamp(scale, 0.25, 10),
	)
}

func (c *Camera) center() galx.Vec2d {
	return c.position.Add(galx.Vec2d{
		X: (float64(c.width) / c.scale) / 2,
		Y: (float64(c.height) / c.scale) / 2,
	})
}

func (c *Camera) dispatchUpdate() {
	c.dispatcher.PublishEventCameraUpdate(event2.CameraUpdateEvent{
		Width:  c.Width(),
		Height: c.Height(),
		Scale:  c.Scale(),
	})
}

func (c *Camera) onFrameEnd(_ event2.FrameEndEvent) error {
	updated := c.applyNewScale() || c.applyNewSize()

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

func (c *Camera) applyNewScale() bool {
	if c.queued.scale == c.scale {
		return false
	}

	c.autoCorrectCenter(func() {
		c.scale = c.queued.scale
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
