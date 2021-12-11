package ui

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
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

func (u *UI) OnDraw(r galx.Renderer) error {
	for _, layer := range u.layers {
		err := layer.OnDraw(r)
		if err != nil {
			return fmt.Errorf("draw ui: %w", err)
		}
	}

	return nil
}

func (u *UI) OnUpdate(s galx.State) error {
	for _, layer := range u.layers {
		err := layer.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("draw ui: %w", err)
		}
	}

	return nil
}
