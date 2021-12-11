package entity

import (
	"github.com/fe3dback/galaxy/galx"
)

type (
	Component interface {
		Id() string
		Title() string
		Description() string
	}

	ComponentCycleCreated interface {
		OnCreated(entity galx.GameObject)
	}

	ComponentCycleUpdated interface {
		galx.Updater
	}

	ComponentCycleDrawer interface {
		galx.Drawer
	}
)
