package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/event"
)

type (
	Camera struct {
		position galx.Vec
		width    int
		height   int
		scale    float64

		dispatcher *event.Dispatcher
		queued     queued
	}

	queued struct {
		width  int
		height int
		scale  float64
	}
)

func NewCamera(dispatcher *event.Dispatcher, defaultWidth int, defaultHeight int) *Camera {
	cam := &Camera{
		dispatcher: dispatcher,

		position: galx.Vec{},
		width:    defaultWidth,
		height:   defaultHeight,
		scale:    1,
	}

	cam.queued = queued{
		width:  cam.width,
		height: cam.height,
		scale:  cam.scale,
	}

	dispatcher.OnFrameEnd(cam.onFrameEnd)
	dispatcher.OnWindowResized(cam.onWindowResize)
	return cam
}

func (c *Camera) Position() galx.Vec {
	return c.position
}

func (c *Camera) Screen2World(screen galx.Vec) galx.Vec {
	return screen.Decrease(c.scale).Add(c.position)
}

func (c *Camera) World2Screen(world galx.Vec) galx.Vec {
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

func (c *Camera) MoveTo(p galx.Vec) {
	c.position = p
}

func (c *Camera) CenterOn(p galx.Vec) {
	c.MoveTo(galx.Vec{
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

func (c *Camera) center() galx.Vec {
	return c.position.Add(galx.Vec{
		X: (float64(c.width) / c.scale) / 2,
		Y: (float64(c.height) / c.scale) / 2,
	})
}

func (c *Camera) dispatchUpdate() {
	c.dispatcher.PublishEventCameraUpdate(event.CameraUpdateEvent{
		Width:  c.Width(),
		Height: c.Height(),
		Scale:  c.Scale(),
	})
}

func (c *Camera) onFrameEnd(_ event.FrameEndEvent) error {
	updated := c.applyNewScale() || c.applyNewSize()

	if updated {
		// update renderer properties
		c.dispatchUpdate()
	}

	return nil
}

func (c *Camera) onWindowResize(ev event.WindowResizedEvent) error {
	c.Resize(ev.NewWidth, ev.NewHeight)

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
