package engine

import "github.com/fe3dback/galaxy/render"

type (
	Drawer interface {
		OnDraw(*render.Renderer) error
	}

	Updater interface {
		OnUpdate(Moment) error
	}
)
