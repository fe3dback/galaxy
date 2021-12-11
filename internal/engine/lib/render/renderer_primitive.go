package render

import (
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/utils"

	"github.com/veandco/go-sdl2/sdl"
)

func (r *Renderer) internalDrawSquare(color galx.Color, rect sdl.Rect) {
	r.SetDrawColor(color)
	err := r.ref.DrawLines([]sdl.Point{
		{X: rect.X, Y: rect.Y},
		{X: rect.X + rect.W, Y: rect.Y},
		{X: rect.X + rect.W, Y: rect.Y + rect.H},
		{X: rect.X, Y: rect.Y + rect.H},
		{X: rect.X, Y: rect.Y},
	})
	utils.Check("draw square", err)
}

func (r *Renderer) internalDrawCircle(color galx.Color, pos sdl.Point, radius float64) {
	r.SetDrawColor(color)

	diameter := int(radius * 2)

	x := int32(radius - 1)
	var y int32 = 0
	tx := 1
	ty := 1
	err := tx - diameter

	for {
		if x >= y {
			break
		}

		centreX := pos.X
		centreY := pos.Y

		//  Each of the following renders an octant of the circle
		r.internalDrawPoint(color, sdl.Point{X: centreX + x, Y: centreY - y})
		r.internalDrawPoint(color, sdl.Point{X: centreX + x, Y: centreY + y})
		r.internalDrawPoint(color, sdl.Point{X: centreX - x, Y: centreY - y})
		r.internalDrawPoint(color, sdl.Point{X: centreX - x, Y: centreY + y})
		r.internalDrawPoint(color, sdl.Point{X: centreX + y, Y: centreY - x})
		r.internalDrawPoint(color, sdl.Point{X: centreX + y, Y: centreY + x})
		r.internalDrawPoint(color, sdl.Point{X: centreX - y, Y: centreY - x})
		r.internalDrawPoint(color, sdl.Point{X: centreX - y, Y: centreY + x})

		if err <= 0 {
			y++
			err += ty
			ty += 2
		}

		if err > 0 {
			x--
			tx += 2
			err += tx - diameter
		}
	}
}

func (r *Renderer) internalDrawLines(color galx.Color, line []sdl.Point) {
	r.SetDrawColor(color)
	err := r.ref.DrawLines(line)
	utils.Check("draw line", err)
}

func (r *Renderer) internalDrawPoint(color galx.Color, pos sdl.Point) {
	r.SetDrawColor(color)
	err := r.ref.DrawPoint(pos.X, pos.Y)
	utils.Check("draw point", err)
}

// -------------------------------------------
// extended function (based on originals)
// -------------------------------------------

func (r *Renderer) internalDrawCrossLines(color galx.Color, size int32, pos sdl.Point) {
	r.internalDrawLines(
		color,
		[]sdl.Point{
			{X: pos.X - size, Y: pos.Y - size},
			{X: pos.X + size, Y: pos.Y + size},
		},
	)
	r.internalDrawLines(
		color,
		[]sdl.Point{
			{X: pos.X - size, Y: pos.Y + size},
			{X: pos.X + size, Y: pos.Y - size},
		},
	)
}
