package entity

import "github.com/fe3dback/galaxy/engine"

type Component interface {
	engine.Drawer
	engine.Updater
}

// todo: codegen for Component getter with typecast
