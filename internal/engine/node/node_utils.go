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

		if chBBox.TL.X < minX {
			minX = chBBox.TL.X
		}
		if chBBox.TL.Y < minY {
			minY = chBBox.TL.Y
		}

		if chBBox.BR.X > maxX {
			maxX = chBBox.BR.X
		}
		if chBBox.BR.Y > maxY {
			maxY = chBBox.BR.Y
		}
	}

	return galx.Rect{
		TL: galx.Vec{
			X: minX,
			Y: minY,
		},
		BR: galx.Vec{
			X: maxX,
			Y: maxY,
		},
	}
}
