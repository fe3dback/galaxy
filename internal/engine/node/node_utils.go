package node

import (
	"github.com/fe3dback/galaxy/galx"
)

func (n *Node) BoundingBox(padding float64) galx.Rect {
	pos := n.AbsPosition()
	minX := pos.X - padding
	minY := pos.Y - padding
	maxX := pos.X + padding
	maxY := pos.Y + padding

	for _, child := range n.Child() {
		chBBox := child.BoundingBox(padding)

		if chBBox.Min.X < minX {
			minX = chBBox.Min.X
		}
		if chBBox.Min.Y < minY {
			minY = chBBox.Min.Y
		}

		if chBBox.Max.X > maxX {
			maxX = chBBox.Max.X
		}
		if chBBox.Max.Y > maxY {
			maxY = chBBox.Max.Y
		}
	}

	return galx.Rect{
		Min: galx.Vec{
			X: minX,
			Y: minY,
		},
		Max: galx.Vec{
			X: maxX,
			Y: maxY,
		},
	}
}
