package ui

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
)

type UI struct {
	layers []Layer
}

func NewUI(layers ...Layer) *UI {
	ui := &UI{
		layers: layers,
	}

	return ui
}

func (u *UI) OnDraw(r engine.Renderer) error {
	for _, layer := range u.layers {
		err := layer.OnDraw(r)
		if err != nil {
			return fmt.Errorf("draw ui: %v", err)
		}
	}

	return nil
}

func (u *UI) OnUpdate(s engine.State) error {
	for _, layer := range u.layers {
		err := layer.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("draw ui: %v", err)
		}
	}

	return nil
}
