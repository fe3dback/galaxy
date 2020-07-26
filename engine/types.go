package engine

type (
	Drawer interface {
		OnDraw() error
	}

	Updater interface {
		OnUpdate(deltaTime float64) error
	}
)
