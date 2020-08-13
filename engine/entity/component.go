package entity

import "github.com/fe3dback/galaxy/engine"

type Component interface {
	engine.Drawer
	engine.Updater
}
