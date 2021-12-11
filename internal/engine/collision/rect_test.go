package collision

import (
	"testing"

	"github.com/fe3dback/galaxy/galx"
)

// convert [center [x,y], w, h] -> [min[x,y], max[x,y]]
func testNormalizeRect(rect galx.Rect) galx.Rect {
	aHalfWidth := rect.Max.X / 2
	aHalfHeight := rect.Max.Y / 2
	return galx.Rect{
		Min: galx.Vec{
			X: rect.Min.X - aHalfWidth,
			Y: rect.Min.Y - aHalfHeight,
		},
		Max: galx.Vec{
			X: rect.Min.X + aHalfWidth,
			Y: rect.Min.Y + aHalfHeight,
		},
	}
}

func TestCollideBoxToBox(t *testing.T) {
	type args struct {
		a galx.Rect
		b galx.Rect
	}

	centeredBoxS1 := galx.Rect{
		Min: galx.Vec{X: 0, Y: 0}, // x, y
		Max: galx.Vec{X: 2, Y: 2}, // w, h
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// S1: not collide
		{
			name: "top left corner not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -3, Y: -3}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "top right corner not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 3, Y: -3}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom left corner not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -3, Y: 3}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom right corner not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 3, Y: 3}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		// S1: not collide alignment
		{
			name: "top not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 0, Y: -3}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 0, Y: 3}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "left not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -3, Y: 0}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "right not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 3, Y: 0}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},

		// S1.5: pixel perfect sides collide
		{
			name: "top collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 0, Y: -2}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 0, Y: 2}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "left collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -2, Y: 2}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "right collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 2, Y: 0}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},

		// S1.75: pixel perfect float offset not collide
		{
			name: "top float not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 0, Y: -2.00001}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom float not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 0, Y: 2.00001}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "left float not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -2.00001, Y: 0}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "right float not collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 2.00001, Y: 0}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},

		// S2: pixel perfect collide (corners collide)
		{
			name: "top left corner collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -2, Y: -2}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "top right corner collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 2, Y: -2}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom left corner collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -2, Y: 2}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom right corner collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 2, Y: 2}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},

		// S2: small overlap collide
		{
			name: "top left small overlap collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -1.5, Y: -1.5}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "top right small overlap collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 1.5, Y: -1.5}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom left small overlap collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -1.5, Y: 1.5}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom right small overlap collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 1.5, Y: 1.5}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: true,
		},

		// S2: close but not collide
		{
			name: "top left close float, no collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -2.00001, Y: -2.00001}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "top right close float, no collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 2.00001, Y: -2.00001}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom left close float, no collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -2.00001, Y: 2.00001}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom right close float, no collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 2.00001, Y: 2.00001}, Max: galx.Vec{X: 2, Y: 2}},
			},
			want: false,
		},

		// special cases
		{
			name: "overlap collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 0, Y: 0}, Max: galx.Vec{X: 1, Y: 1}},
			},
			want: true,
		},
		{
			name: "overlap small size collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 0, Y: 0}, Max: galx.Vec{X: 0.0001, Y: 0.0001}},
			},
			want: true,
		},
		{
			name: "overlap small size offset TL collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: -1, Y: -1}, Max: galx.Vec{X: 0.0001, Y: 0.0001}},
			},
			want: true,
		},
		{
			name: "overlap small size offset BR collide",
			args: args{
				a: centeredBoxS1,
				b: galx.Rect{Min: galx.Vec{X: 1, Y: 1}, Max: galx.Vec{X: 0.0001, Y: 0.0001}},
			},
			want: true,
		},

		// offset of box
		{
			name: "offset corner collide",
			args: args{
				a: galx.Rect{Min: galx.Vec{X: 3, Y: 3}, Max: galx.Vec{X: 1, Y: 1}},
				b: galx.Rect{Min: galx.Vec{X: 2, Y: 2}, Max: galx.Vec{X: 1, Y: 1}},
			},
			want: true,
		},
		{
			name: "offset corner pixel no collide",
			args: args{
				a: galx.Rect{Min: galx.Vec{X: 3, Y: 3}, Max: galx.Vec{X: 1, Y: 1}},
				b: galx.Rect{Min: galx.Vec{X: 2, Y: 2}, Max: galx.Vec{X: 0.9999, Y: 0.9999}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aNorm := testNormalizeRect(tt.args.a)
			bNorm := testNormalizeRect(tt.args.b)

			variants := [][2]galx.Rect{
				{aNorm, bNorm},
				{bNorm, aNorm},
			}

			for varId, variant := range variants {
				if got := Rect2Rect(variant[0], variant[1]); got != tt.want {
					t.Errorf("Rect2Rect() = got '%v', want '%v'\nVariant: #%v\nx,y,w,h\n%+v\n%+v\nmin [x,y], max [x,y]\n%+v\n%+v\n",
						got,
						tt.want,
						varId+1,
						variant[0],
						variant[1],
						tt.args.a,
						tt.args.b,
					)
				}
			}
		})
	}
}

func TestRect2Point(t *testing.T) {
	type args struct {
		a galx.Rect
		b galx.Vec
	}

	centered2x2 := galx.Rect{Min: galx.Vec{X: 0, Y: 0}, Max: galx.Vec{X: 2, Y: 2}}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// corners collide
		{
			name: "top left corner collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: -1, Y: -1},
			},
			want: true,
		},
		{
			name: "top right corner collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 1, Y: -1},
			},
			want: true,
		},
		{
			name: "bottom left corner collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: -1, Y: 1},
			},
			want: true,
		},
		{
			name: "bottom right corner collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 1, Y: 1},
			},
			want: true,
		},

		// sides collide
		{
			name: "top side collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 0, Y: -1},
			},
			want: true,
		},
		{
			name: "bottom side collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 0, Y: 1},
			},
			want: true,
		},
		{
			name: "left side collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: -1, Y: 0},
			},
			want: true,
		},
		{
			name: "right side collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 1, Y: 0},
			},
			want: true,
		},

		// corners offset not collide
		{
			name: "top left corner offset not collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: -1.00001, Y: -1},
			},
			want: false,
		},
		{
			name: "top right corner offset not collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 1, Y: -1.00001},
			},
			want: false,
		},
		{
			name: "bottom left corner offset not collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: -1, Y: 1.00001},
			},
			want: false,
		},
		{
			name: "bottom right corner offset not collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 1.00001, Y: 1},
			},
			want: false,
		},

		// sides offset not collide
		{
			name: "top side offset not collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 0, Y: -1.00001},
			},
			want: false,
		},
		{
			name: "bottom side offset not collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 0, Y: 1.00001},
			},
			want: false,
		},
		{
			name: "left side offset not collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: -1.00001, Y: 0},
			},
			want: false,
		},
		{
			name: "right side offset not collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 1.00001, Y: 0},
			},
			want: false,
		},

		// special
		{
			name: "inside collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 0, Y: 0},
			},
			want: true,
		},
		{
			name: "inside tl collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: -0.9999, Y: -0.9999},
			},
			want: true,
		},
		{
			name: "inside br collide",
			args: args{
				a: centered2x2,
				b: galx.Vec{X: 0.9999, Y: 0.9999},
			},
			want: true,
		},

		// offset not collide
		{
			name: "inside not centered tl collide",
			args: args{
				a: galx.Rect{
					Min: galx.Vec{X: -1, Y: -1}, // x,y
					Max: galx.Vec{X: 2, Y: 2},   // w,h
				},
				b: galx.Vec{X: -2, Y: -2},
			},
			want: true,
		},
		{
			name: "inside not centered tl not collide",
			args: args{
				a: galx.Rect{
					Min: galx.Vec{X: 0, Y: 0},             // x,y
					Max: galx.Vec{X: 1.99999, Y: 1.99999}, // w,h
				},
				b: galx.Vec{X: -2, Y: -2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normBox := testNormalizeRect(tt.args.a)

			if got := Rect2Point(normBox, tt.args.b); got != tt.want {
				t.Errorf("Rect2Point() = %v, want %v (%+v)\nnorm %+v\n raw %+v",
					got,
					tt.want,
					tt.args.b,
					normBox,
					tt.args.a,
				)
			}
		})
	}
}

func TestRect2Circle(t *testing.T) {
	type args struct {
		r galx.Rect
		c galx.Circle
	}

	centeredCircle2R := galx.Circle{Pos: galx.Vec{}, Radius: 2}
	centeredBox2R := galx.Rect{
		Min: galx.Vec{},           // x,y
		Max: galx.Vec{X: 2, Y: 2}, // w,h
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s1",
			args: args{
				r: centeredBox2R,
				c: centeredCircle2R,
			},
			want: true,
		},
		{
			name: "s2",
			args: args{
				r: centeredBox2R,
				c: galx.Circle{Pos: galx.Vec{X: 2, Y: 0}, Radius: 2},
			},
			want: true,
		},
		{
			name: "s3",
			args: args{
				r: centeredBox2R,
				c: galx.Circle{Pos: galx.Vec{X: 3, Y: 0}, Radius: 2},
			},
			want: true,
		},
		{
			name: "s3+",
			args: args{
				r: centeredBox2R,
				c: galx.Circle{Pos: galx.Vec{X: 3.0001, Y: 0}, Radius: 2},
			},
			want: false,
		},
		{
			name: "s3-",
			args: args{
				r: centeredBox2R,
				c: galx.Circle{Pos: galx.Vec{X: 2.9999, Y: 0}, Radius: 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		normBox := testNormalizeRect(tt.args.r)

		t.Run(tt.name, func(t *testing.T) {
			if got := Rect2Circle(normBox, tt.args.c); got != tt.want {
				t.Errorf("Rect2Circle() = %v, want %v", got, tt.want)
			}
		})
	}
}
