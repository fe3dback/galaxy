package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/utils"
)

// Proxy user for engine input normalization, (for example convert rect to screen space)
// also only this renderer should check clip rect's

func (r *Renderer) transRect(rect galx.Rect) sdl.Rect {
	if !rect.Valid() {
		rect = rect.Normalize()
	}

	size := rect.BR.Sub(rect.TL)
	return Rect{
		X: int32(r.screenX(rect.TL.X)),
		Y: int32(r.screenY(rect.TL.Y)),
		W: int32(size.X),
		H: int32(size.Y),
	}
}

func (r *Renderer) transCircle(circle galx.Circle) (sdl.Point, float64) {
	return r.transPoint(circle.Pos), circle.Radius
}

func (r *Renderer) transRectPtr(rect galx.Rect) *sdl.Rect {
	rRect := r.transRect(rect)
	return &rRect
}

func (r *Renderer) transPoint(point galx.Vec2d) sdl.Point {
	return Point{
		X: int32(r.screenX(point.X)),
		Y: int32(r.screenY(point.Y)),
	}
}

func (r *Renderer) transLine(line galx.Line) []sdl.Point {
	norm := line.Normalize()
	return []sdl.Point{
		r.transPoint(norm.A),
		r.transPoint(norm.B),
	}
}

func transformColor(color galx.Color) sdl.Color {
	r, g, b, a := color.Split()
	return sdl.Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

// ==================================================
// Primitive
// ==================================================

func (r *Renderer) DrawSquare(color galx.Color, rect galx.Rect) {
	norm := r.transRect(rect)
	if !r.isRectInsideCamera(norm) {
		return
	}

	r.internalDrawSquare(color, norm)
}

func (r *Renderer) DrawSquareFilled(color galx.Color, rect galx.Rect) {
	norm := r.transRect(rect)
	if !r.isRectInsideCamera(norm) {
		return
	}

	r.internalDrawSquareFilled(color, norm)
}

func (r *Renderer) DrawCircle(color galx.Color, circle galx.Circle) {
	norm := r.transRect(circle.BoundingBox())
	if !r.isRectInsideCamera(norm) {
		return
	}

	pos, radius := r.transCircle(circle)
	r.internalDrawCircle(color, pos, radius)
}

func (r *Renderer) DrawSquareEx(color galx.Color, angle galx.Angle, rect galx.Rect) {
	norm := r.transRect(rect)
	if !r.isRectInsideCamera(norm) {
		return
	}

	center := galx.Vec2d{
		X: float64(norm.X + norm.W/2),
		Y: float64(norm.Y + norm.H/2),
	}

	vertices := [4]galx.Vec2d{
		galx.Vec2d{
			X: rect.TL.X,
			Y: rect.TL.Y,
		}.RotateAround(center, angle),

		galx.Vec2d{
			X: rect.TL.X + rect.BR.X,
			Y: rect.TL.Y,
		}.RotateAround(center, angle),

		galx.Vec2d{
			X: rect.TL.X + rect.BR.X,
			Y: rect.TL.Y + rect.BR.Y,
		}.RotateAround(center, angle),

		galx.Vec2d{
			X: rect.TL.X,
			Y: rect.TL.Y + rect.BR.Y,
		}.RotateAround(center, angle),
	}

	r.internalDrawLines(color, []sdl.Point{
		r.transPoint(vertices[0]),
		r.transPoint(vertices[1]),
		r.transPoint(vertices[2]),
		r.transPoint(vertices[3]),
		r.transPoint(vertices[0]),
	})
}

func (r *Renderer) DrawLine(color galx.Color, line galx.Line) {
	norm := r.transLine(line)
	if !r.isLineInsideCamera(norm) {
		return
	}

	r.internalDrawLines(color, norm)
}

func (r *Renderer) DrawPoint(color galx.Color, vec galx.Vec2d) {
	p := r.transPoint(vec)
	if !r.isPointInsideCamera(p) {
		return
	}

	r.internalDrawPoint(color, p)
}

func (r *Renderer) DrawVector(color galx.Color, dist float64, vec galx.Vec2d, angle galx.Angle) {
	target := vec.PolarOffset(dist, angle)
	line := galx.Line{
		A: vec,
		B: target,
	}

	normLine := r.transLine(line)
	if !r.isLineInsideCamera(normLine) {
		return
	}

	// draw
	counterDeg := angle.Add(galx.NewAngle(180))
	arrowLeft := galx.Line{
		A: target,
		B: target.PolarOffset(6, counterDeg.Add(galx.NewAngle(-30))),
	}
	arrowRight := galx.Line{
		A: target,
		B: target.PolarOffset(6, counterDeg.Add(galx.NewAngle(+30))),
	}

	r.internalDrawLines(color, normLine)
	r.internalDrawLines(color, r.transLine(arrowLeft))
	r.internalDrawLines(color, r.transLine(arrowRight))
}

func (r *Renderer) DrawCrossLines(color galx.Color, size int, vec galx.Vec2d) {
	p := r.transPoint(vec)
	if !r.isPointInsideCamera(p) {
		return
	}

	r.internalDrawCrossLines(color, int32(size), p)
}

// ==================================================
// Sprite
// ==================================================

func (r *Renderer) DrawSprite(res consts.AssetsPath, vec galx.Vec2d) {
	r.drawTexture(
		res,
		Rect{},
		Rect{X: int32(vec.X), Y: int32(vec.Y)},
		galx.NewAngle(0),
	)
}

func (r *Renderer) DrawSpriteAngle(res consts.AssetsPath, vec galx.Vec2d, angle galx.Angle) {
	r.drawTexture(
		res,
		Rect{},
		Rect{X: int32(vec.X), Y: int32(vec.Y)},
		angle,
	)
}

func (r *Renderer) DrawSpriteEx(res consts.AssetsPath, src, dest galx.Rect, angle galx.Angle) {
	normSrc := r.transRect(src)
	normDest := r.transRect(dest)
	r.drawTexture(res, normSrc, normDest, angle)
}

func (r *Renderer) drawTexture(res consts.AssetsPath, src, dest Rect, angle galx.Angle) {
	defer utils.CheckPanic(fmt.Sprintf("draw sprite `%s`", res))

	texture := r.getTexture(res)

	if dest.W == 0 {
		dest.W = texture.Width
	}
	if dest.H == 0 {
		dest.H = texture.Height
	}

	// check is visible
	if !r.isRectInsideCamera(dest) {
		return
	}

	if src.W == 0 {
		src.W = texture.Width
	}
	if src.H == 0 {
		src.H = texture.Height
	}

	r.internalDrawTexture(texture.Tex, src, dest, angle)
}

// ==================================================
// Text
// ==================================================

func (r *Renderer) DrawText(fontId consts.AssetsPath, color galx.Color, vec galx.Vec2d, text string) {
	norm := Rect{
		X: int32(vec.X),
		Y: int32(vec.Y),
		W: avgTextWidthOptRender,  // todo magic numbers
		H: avgTextHeightOptRender, // todo magic numbers
	}

	if !r.isRectInsideCamera(norm) {
		return
	}

	r.internalDrawText(fontId, color, text, r.transPoint(vec))
}

// ==================================================
// Other
// ==================================================
