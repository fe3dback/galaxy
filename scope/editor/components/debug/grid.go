package debug

import (
	"fmt"
	"math"

	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type Grid struct {
	settings settingsPane
	camera   galx.Camera

	// settings
	sizeX float64
	sizeY float64
}

func NewGrid(settings settingsPane) *Grid {
	return &Grid{
		settings: settings,
		sizeX:    64,
		sizeY:    64,
	}
}

func (c *Grid) OnUpdate(s galx.State) error {
	c.displaySettingsWindow()
	c.camera = s.Camera()
	return nil
}

func (c *Grid) displaySettingsWindow() {
	c.settings.Extend("Grid", 5, func() {
		sizeX := float32(c.sizeX)
		sizeY := float32(c.sizeY)

		imgui.DragFloatV("Size X", &sizeX, 1, 4, 1024, "%.0f", imgui.SliderFlagsNone)
		imgui.DragFloatV("Size Y", &sizeY, 1, 4, 1024, "%.0f", imgui.SliderFlagsNone)

		c.sizeX = float64(sizeX)
		c.sizeY = float64(sizeY)
	})
}

func (c *Grid) OnDraw(r galx.Renderer) error {
	if !r.Gizmos().Debug() {
		return nil
	}

	startAt := c.camera.Position()
	endAt := galx.Vec{
		X: startAt.X + float64(c.camera.Width()),
		Y: startAt.Y + float64(c.camera.Height()),
	}

	for x := startAt.X - c.sizeX; x < endAt.X+c.sizeX; x += c.sizeX {
		for y := startAt.Y - c.sizeY; y < endAt.Y+c.sizeY; y += c.sizeY {
			rX := math.Floor(x/c.sizeX) * c.sizeX
			rY := math.Floor(y/c.sizeY) * c.sizeY

			r.DrawLine(galx.ColorGray1, galx.Line{
				A: galx.Vec{X: rX, Y: rY},
				B: galx.Vec{X: rX + c.sizeX, Y: rY},
			})
			r.DrawLine(galx.ColorGray1, galx.Line{
				A: galx.Vec{X: rX, Y: rY},
				B: galx.Vec{X: rX, Y: rY + c.sizeY},
			})

			if c.sizeX >= 64 && c.sizeY >= 64 {
				r.DrawText(consts.AssetDefaultFont, galx.ColorGray1, fmt.Sprintf("%.0fx%.0f", rX, rY), galx.Vec{X: rX, Y: rY})
			}
		}
	}

	return nil
}