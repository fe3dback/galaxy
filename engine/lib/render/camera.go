package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
)

type Camera struct {
	position engine.Vec
	width    int
	height   int
}

func NewCamera() *Camera {
	return &Camera{
		position: engine.Vec{},
		width:    320,
		height:   240,
	}
}

func (c *Camera) Position() engine.Vec {
	return c.position
}

func (c *Camera) Width() int {
	return c.width
}

func (c *Camera) Height() int {
	return c.height
}

func (c *Camera) MoveTo(p engine.Vec) {
	c.position = p
}

func (c *Camera) CenterOn(p engine.Vec) {
	c.MoveTo(engine.Vec{
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
