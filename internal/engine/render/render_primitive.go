package render

import "github.com/fe3dback/galaxy/galx"

func (r *Render) DrawSquare(color galx.Color, rect galx.Rect) {
	var vertPos [4]galx.Vec2d

	for ind, vec := range rect.VerticesClockWise() {
		transform := r.project(r.cam(vec))
		vertPos[ind] = galx.Vec2d{
			X: transform.X,
			Y: transform.Y,
		}
	}

	r.renderer.DrawRect(
		vertPos,
		[4]galx.Vec3d{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
			{1, 0, 1},
		},
	)
}

func (r *Render) DrawSquareFilled(color galx.Color, rect galx.Rect) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) DrawCircle(color galx.Color, circle galx.Circle) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) DrawSquareEx(color galx.Color, angle galx.Angle, rect galx.Rect) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) DrawLine(color galx.Color, line galx.Line) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) DrawVector(color galx.Color, dist float64, vec galx.Vec2d, angle galx.Angle) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) DrawCrossLines(color galx.Color, size int, vec galx.Vec2d) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) DrawPoint(color galx.Color, vec galx.Vec2d) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) FillRect(rect galx.Rect) {
	// TODO implement me
	panic("implement me")
}
