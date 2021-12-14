package control

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
)

const cameraMaxZoom = 10
const cameraMinZoom = 0.5

type Camera struct {
	settings settingsPane
	camera   galx.Camera

	// settings
	zoom        float64
	prevZoom    float64
	cameraSpeed float64
}

func NewCamera(settings settingsPane) *Camera {
	return &Camera{
		settings:    settings,
		zoom:        1,
		prevZoom:    1,
		cameraSpeed: 50.0,
	}
}

func (c *Camera) OnUpdate(state galx.State) error {
	c.camera = state.Camera()
	c.displaySettings()

	// zoom
	lastScroll := state.Mouse().ScrollLastOffset()
	if lastScroll != 0 {
		c.zoom = galx.Clamp(c.zoom+lastScroll*cameraMinZoom, cameraMinZoom, cameraMaxZoom)
	}

	// move camera
	speed := c.cameraSpeed
	if state.Movement().Shift() {
		speed *= 5
	}

	state.Camera().MoveTo(
		state.Camera().Position().Add(state.Movement().Vector().Scale(speed)),
	)

	if c.prevZoom != c.zoom {
		state.Camera().ZoomView(c.zoom)
		c.prevZoom = c.zoom
	}

	return nil
}

func (c *Camera) displaySettings() {
	c.settings.Extend("Camera", 10, func() {
		// props
		speed := int32(c.cameraSpeed)
		zoom := float32(c.zoom)

		imgui.DragIntV("Speed", &speed, 1, 1, 512, "%d", imgui.SliderFlagsNone)
		imgui.DragFloatV("Zoom", &zoom, cameraMinZoom, cameraMinZoom, cameraMaxZoom, "%.1f", imgui.SliderFlagsNone)

		c.cameraSpeed = float64(speed)
		c.zoom = float64(zoom)

		// pos info
		imgui.Text(fmt.Sprintf("X = %.0f", c.camera.Position().X))
		imgui.Text(fmt.Sprintf("Y = %.0f", c.camera.Position().Y))
		imgui.Text(fmt.Sprintf("W = %d", c.camera.Width()))
		imgui.Text(fmt.Sprintf("H = %d", c.camera.Height()))
	})
}

func (c *Camera) OnDraw(_ galx.Renderer) error {
	return nil
}
