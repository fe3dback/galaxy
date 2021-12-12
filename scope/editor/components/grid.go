package components

import (
	"fmt"
	"math"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type Grid struct {
	size   float64
	camera galx.Camera
}

func NewGrid() *Grid {
	return &Grid{
		size: 64,
	}
}

func (c *Grid) OnUpdate(s galx.State) error {
	c.camera = s.Camera()
	return nil
}

func (c *Grid) OnDraw(r galx.Renderer) error {
	if !r.Gizmos().Secondary() {
		return nil
	}

	startAt := c.camera.Position()
	endAt := galx.Vec{
		X: startAt.X + float64(c.camera.Width()),
		Y: startAt.Y + float64(c.camera.Height()),
	}

	for x := startAt.X - c.size; x < endAt.X+c.size; x += c.size {
		for y := startAt.Y - c.size; y < endAt.Y+c.size; y += c.size {
			rX := math.Floor(x/c.size) * c.size
			rY := math.Floor(y/c.size) * c.size

			r.DrawLine(galx.ColorGray1, galx.Line{
				A: galx.Vec{X: rX, Y: rY},
				B: galx.Vec{X: rX + c.size, Y: rY},
			})
			r.DrawLine(galx.ColorGray1, galx.Line{
				A: galx.Vec{X: rX, Y: rY},
				B: galx.Vec{X: rX, Y: rY + c.size},
			})

			r.DrawText(consts.AssetDefaultFont, galx.ColorGray1, fmt.Sprintf("%.0fx%.0f", rX, rY), galx.Vec{X: rX, Y: rY})
		}
	}

	return nil
}
