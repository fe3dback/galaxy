package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
)

type Camera struct {
	position engine.Vector2D
	width    int
	height   int
}

func NewCamera() *Camera {
	return &Camera{
		position: engine.Vector2D{},
		width:    320,
		height:   240,
	}
}

func (c *Camera) Position() engine.Vector2D {
	return c.position
}

func (c *Camera) Width() int {
	return c.width
}

func (c *Camera) Height() int {
	return c.height
}

func (c *Camera) MoveTo(p engine.Vector2D) {
	c.position = p
}

func (c *Camera) CenterOn(p engine.Vector2D) {
	c.MoveTo(engine.Vector2D{
		X: float64(p.RoundX() - c.width/2),
		Y: float64(p.RoundY() - c.height/2),
	})
}

func (c *Camera) Resize(width, height int) {
	if width < 1 || height < 1 {
		panic(fmt.Sprintf("can`t resize camera to %d x %d", width, height))
	}

	c.width = width
	c.height = height
}
