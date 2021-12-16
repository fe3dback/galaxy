package control

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/scope/editor/components/gui"
)

const cameraMaxScale = 10
const cameraMinScale = 0.5

type Camera struct {
	settings *gui.Settings
	camera   galx.Camera

	// settings
	scale       float64
	prevScale   float64
	cameraSpeed float64
}

func (c Camera) Id() string {
	return "6055b322-cbe6-4102-bb85-5a52e4f59d42"
}

func (c *Camera) OnCreated(require galx.EditorComponentResolver) {
	c.settings = require(c.settings).(*gui.Settings)
	c.scale = 1
	c.prevScale = 1
	c.cameraSpeed = 50.0
}

func (c *Camera) OnUpdate(state galx.State) error {
	c.camera = state.Camera()
	c.displaySettings()

	// scale
	lastScroll := state.Mouse().ScrollLastOffset()
	if lastScroll != 0 {
		c.scale = galx.Clamp(c.scale+lastScroll*cameraMinScale, cameraMinScale, cameraMaxScale)
	}

	// move camera
	speed := c.cameraSpeed / c.scale
	if state.Movement().Shift() {
		speed *= 5
	}

	state.Camera().MoveTo(
		state.Camera().Position().Add(state.Movement().Vector().Scale(speed)),
	)

	if c.prevScale != c.scale {
		state.Camera().ZoomView(c.scale)
		c.prevScale = c.scale
	}

	return nil
}

func (c *Camera) displaySettings() {
	c.settings.Extend("Camera", 10, func() {
		// props
		speed := int32(c.cameraSpeed)
		scale := float32(c.scale)

		imgui.DragIntV("Speed", &speed, 1, 1, 512, "%d", imgui.SliderFlagsNone)
		imgui.DragFloatV("Scale", &scale, cameraMinScale, cameraMinScale, cameraMaxScale, "%.1f", imgui.SliderFlagsNone)

		c.cameraSpeed = float64(speed)
		c.scale = float64(scale)

		// pos info
		imgui.Text(fmt.Sprintf("X = %.0f", c.camera.Position().X))
		imgui.Text(fmt.Sprintf("Y = %.0f", c.camera.Position().Y))
		imgui.Text(fmt.Sprintf("W = %d", c.camera.Width()))
		imgui.Text(fmt.Sprintf("H = %d", c.camera.Height()))
	})
}
