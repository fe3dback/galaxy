package di

import (
	"github.com/fe3dback/galaxy/scope/editor"
)

func (c *Container) ProvideEditorManager() *editor.Manager {
	if c.memstate.editor.manager != nil {
		return c.memstate.editor.manager
	}

	// todo auto register all editor components
	manager := editor.NewManager(
		// todo: turn on after new vulkan renderer
		// gui
		// &gui.Scene{},
		// &gui.Help{},
		// &gui.Settings{},
		// &gui.Hierarchy{},
		//
		// // components
		// &debug.Grid{},
		// &control.Camera{},
		// &control.Transform{}, // control level: 1
		// &control.Select{},    // control level: 2
	)

	c.memstate.editor.manager = manager
	return c.memstate.editor.manager
}
