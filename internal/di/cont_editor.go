package di

import (
	"github.com/fe3dback/galaxy/scope/editor"
	"github.com/fe3dback/galaxy/scope/editor/components/control"
	"github.com/fe3dback/galaxy/scope/editor/components/debug"
	"github.com/fe3dback/galaxy/scope/editor/components/gui"
)

func (c *Container) ProvideEditorManager() *editor.Manager {
	if c.memstate.editor.manager != nil {
		return c.memstate.editor.manager
	}

	// todo auto register all editor components
	manager := editor.NewManager(
		// gui
		&gui.Help{},
		&gui.Settings{},
		&gui.Hierarchy{},

		// components
		&debug.Grid{},
		&control.Camera{},
		&control.Transform{}, // control level: 1
		&control.Select{},    // control level: 2
	)

	c.memstate.editor.manager = manager
	return c.memstate.editor.manager
}
