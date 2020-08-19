package collision

import (
	"testing"

	"github.com/fe3dback/galaxy/engine"
)

// convert [center [x,y], w, h] -> [min[x,y], max[x,y]]
func testNormalizeRect(rect engine.Rect) engine.Rect {
	aHalfWidth := rect.Max.X / 2
	aHalfHeight := rect.Max.Y / 2
	return engine.Rect{
		Min: engine.Vec{
			X: rect.Min.X - aHalfWidth,
			Y: rect.Min.Y - aHalfHeight,
		},
		Max: engine.Vec{
			X: rect.Min.X + aHalfWidth,
			Y: rect.Min.Y + aHalfHeight,
		},
	}
}

func TestCollideBoxToBox(t *testing.T) {
	type args struct {
		a engine.Rect
		b engine.Rect
	}

	centeredBoxS1 := engine.Rect{
		Min: engine.Vec{X: 0, Y: 0}, // x, y
		Max: engine.Vec{X: 2, Y: 2}, // w, h
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
				b: engine.Rect{Min: engine.Vec{X: -3, Y: -3}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "top right corner not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 3, Y: -3}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom left corner not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -3, Y: 3}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom right corner not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 3, Y: 3}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		// S1: not collide alignment
		{
			name: "top not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 0, Y: -3}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 0, Y: 3}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "left not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -3, Y: 0}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "right not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 3, Y: 0}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},

		// S1.5: pixel perfect sides collide
		{
			name: "top collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 0, Y: -2}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 0, Y: 2}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "left collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -2, Y: 2}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "right collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 2, Y: 0}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},

		// S1.75: pixel perfect float offset not collide
		{
			name: "top float not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 0, Y: -2.00001}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom float not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 0, Y: 2.00001}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "left float not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -2.00001, Y: 0}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "right float not collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 2.00001, Y: 0}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},

		// S2: pixel perfect collide (corners collide)
		{
			name: "top left corner collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -2, Y: -2}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "top right corner collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 2, Y: -2}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom left corner collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -2, Y: 2}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom right corner collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 2, Y: 2}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},

		// S2: small overlap collide
		{
			name: "top left small overlap collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -1.5, Y: -1.5}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "top right small overlap collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 1.5, Y: -1.5}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom left small overlap collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -1.5, Y: 1.5}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},
		{
			name: "bottom right small overlap collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 1.5, Y: 1.5}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: true,
		},

		// S2: close but not collide
		{
			name: "top left close float, no collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -2.00001, Y: -2.00001}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "top right close float, no collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 2.00001, Y: -2.00001}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom left close float, no collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -2.00001, Y: 2.00001}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},
		{
			name: "bottom right close float, no collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 2.00001, Y: 2.00001}, Max: engine.Vec{X: 2, Y: 2}},
			},
			want: false,
		},

		// special cases
		{
			name: "overlap collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 0, Y: 0}, Max: engine.Vec{X: 1, Y: 1}},
			},
			want: true,
		},
		{
			name: "overlap small size collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 0, Y: 0}, Max: engine.Vec{X: 0.0001, Y: 0.0001}},
			},
			want: true,
		},
		{
			name: "overlap small size offset TL collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: -1, Y: -1}, Max: engine.Vec{X: 0.0001, Y: 0.0001}},
			},
			want: true,
		},
		{
			name: "overlap small size offset BR collide",
			args: args{
				a: centeredBoxS1,
				b: engine.Rect{Min: engine.Vec{X: 1, Y: 1}, Max: engine.Vec{X: 0.0001, Y: 0.0001}},
			},
			want: true,
		},

		// offset of box
		{
			name: "offset corner collide",
			args: args{
				a: engine.Rect{Min: engine.Vec{X: 3, Y: 3}, Max: engine.Vec{X: 1, Y: 1}},
				b: engine.Rect{Min: engine.Vec{X: 2, Y: 2}, Max: engine.Vec{X: 1, Y: 1}},
			},
			want: true,
		},
		{
			name: "offset corner pixel no collide",
			args: args{
				a: engine.Rect{Min: engine.Vec{X: 3, Y: 3}, Max: engine.Vec{X: 1, Y: 1}},
				b: engine.Rect{Min: engine.Vec{X: 2, Y: 2}, Max: engine.Vec{X: 0.9999, Y: 0.9999}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aNorm := testNormalizeRect(tt.args.a)
			bNorm := testNormalizeRect(tt.args.b)

			if got := Rect2Rect(aNorm, bNorm); got != tt.want {
				t.Errorf("Rect2Rect() = got '%v', want '%v'\nx,y,w,h\n%+v\n%+v\nmin [x,y], max [x,y]\n%+v\n%+v\n",
					got,
					tt.want,
					aNorm,
					bNorm,
					tt.args.a,
					tt.args.b,
				)
			}
		})
	}
}

func TestRect2Point(t *testing.T) {
	type args struct {
		a engine.Rect
		b engine.Vec
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "left corner collide",
			args: args{
				a: engine.Rect{Min: engine.Vec{X: 0, Y: 0}, Max: engine.Vec{X: 2, Y: 2}},
				b: engine.Vec{X: -1, Y: -1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rect2Point(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Rect2Point() = %v, want %v", got, tt.want)
			}
		})
	}
}
