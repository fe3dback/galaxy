package ui

import (
	"fmt"
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

func (u *UI) OnDraw() error {
	for _, layer := range u.layers {
		err := layer.OnDraw()
		if err != nil {
			return fmt.Errorf("draw ui: %v", err)
		}
	}

	return nil
}
