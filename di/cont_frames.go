package di

import "github.com/fe3dback/galaxy/system"

func (c *Container) ProvideFrames() *system.Frames {
	if c.memstate.frames != nil {
		return c.memstate.frames
	}

	c.memstate.frames = system.NewFrames(
		c.Flags().TargetFPS,
	)

	return c.memstate.frames
}
