package di

import (
	"github.com/fe3dback/galaxy/internal/frames"
)

func (c *Container) ProvideFrames() *frames.Frames {
	if c.memstate.frames != nil {
		return c.memstate.frames
	}

	c.memstate.frames = frames.NewFrames(
		c.Flags().TargetFPS(),
	)

	return c.memstate.frames
}
