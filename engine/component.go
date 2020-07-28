package engine

type Component interface {
	Drawer
	Updater
}

// todo: codegen for Component getter with typecast
