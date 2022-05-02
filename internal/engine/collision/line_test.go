package collision

import (
	"testing"

	"github.com/fe3dback/galaxy/galx"
)

func TestLine2Line(t *testing.T) {

	h0 := 0.0 // height 0
	h1 := 1.0 // height 1

	x0 := 0.0 // hor offset 0
	x1 := 1.0 // hor offset 1

	tests := []struct {
		name string
		a    [4]float64
		b    [4]float64
		want bool
	}{
		{
			name: "zero points collide",
			a:    [4]float64{0, 0, 0, 0},
			b:    [4]float64{0, 0, 0, 0},
			want: true,
		},
		{
			name: "parallel h no collide",
			a:    [4]float64{2, h0, 5, h0},
			b:    [4]float64{2, h1, 5, h1},
			want: false,
		},
		{
			name: "parallel h collide",
			a:    [4]float64{2, h0, 5, h0},
			b:    [4]float64{2, h0, 5, h0},
			want: true,
		},
		{
			name: "parallel v no collide",
			a:    [4]float64{x0, 2, x0, 8},
			b:    [4]float64{x1, 2, x1, 8},
			want: false,
		},
		{
			name: "parallel v collide",
			a:    [4]float64{x0, 2, x0, 8},
			b:    [4]float64{x0, 2, x0, 8},
			want: true,
		},
		{
			name: "point collide",
			a:    [4]float64{1, 1, 1, 1},
			b:    [4]float64{1, 1, 1, 1},
			want: true,
		},
		{
			name: "point neg collide",
			a:    [4]float64{-1, -1, -1, -1},
			b:    [4]float64{-1, -1, -1, -1},
			want: true,
		},
		{
			name: "cross +",
			a:    [4]float64{0, -1, 0, 1},
			b:    [4]float64{-1, 0, 1, 0},
			want: true,
		},
		{
			name: "cross top T collide",
			a:    [4]float64{0, -1, 0, 1},
			b:    [4]float64{-1, -1, 1, -1},
			want: true,
		},
		{
			name: "cross top T offset no collide",
			a:    [4]float64{0, -1, 0, 1},
			b:    [4]float64{-1, -1.0001, 1, -1.0001},
			want: false,
		},
		{
			name: "cross bottom ⊥ collide",
			a:    [4]float64{0, -1, 0, 1},
			b:    [4]float64{-1, 1, 1, 1},
			want: true,
		},
		{
			name: "cross bottom ⊥ offset no collide",
			a:    [4]float64{0, -1, 0, 1},
			b:    [4]float64{-1, 1.0001, 1, 1.0001},
			want: false,
		},
		{
			name: "cross left |- collide",
			a:    [4]float64{0, -1, 0, 1},
			b:    [4]float64{0, 0, 1, 0},
			want: true,
		},
		{
			name: "cross left |- offset no collide",
			a:    [4]float64{-0.0001, -1, -0.0001, 1},
			b:    [4]float64{0, 0, 1, 0},
			want: false,
		},
		{
			name: "cross right -| collide",
			a:    [4]float64{0, -1, 0, 1},
			b:    [4]float64{-1, 0, 0, 0},
			want: true,
		},
		{
			name: "cross right -| offset no collide",
			a:    [4]float64{0.0001, -1, 0.0001, 1},
			b:    [4]float64{-1, 0, 0, 0},
			want: false,
		},
		{
			name: "cross angled x collide BL",
			a:    [4]float64{-1, -1, 1, 1},
			b:    [4]float64{-1, 1, 1, -1},
			want: true,
		},
		{
			name: "cross angled x collide TR",
			a:    [4]float64{-1, -1, 1, 1},
			b:    [4]float64{1, -1, -1, 1},
			want: true,
		},
		{
			name: "cross angled x collide BR",
			a:    [4]float64{1, -1, -1, 1},
			b:    [4]float64{1, 1, -1, -1},
			want: true,
		},
		{
			name: "cross angled x collide TL",
			a:    [4]float64{1, -1, -1, 1},
			b:    [4]float64{-1, -1, 1, 1},
			want: true,
		},
		{
			name: "cross angled T collide",
			a:    [4]float64{-1, -1, 1, 1},
			b:    [4]float64{0, 0, 1, -1},
			want: true,
		},
		{
			name: "cross angled T offset no collide",
			a:    [4]float64{-1, -1, 1, 1},
			b:    [4]float64{0.0001, -0.0001, 1, -1},
			want: false,
		},
		{
			name: "cross angled T rev collide",
			a:    [4]float64{-1, -1, 1, 1},
			b:    [4]float64{0, 0, -1, 1},
			want: true,
		},
		{
			name: "cross angled T rev offset no collide",
			a:    [4]float64{-1, -1, 1, 1},
			b:    [4]float64{-0.0001, 0.0001, -1, 1},
			want: false,
		},
		{
			name: "s1 collide",
			a:    [4]float64{0, 0, 10, 1},
			b:    [4]float64{0, -1, 10, 2},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1A := galx.Vec2d{X: tt.a[0], Y: tt.a[1]}
			l1B := galx.Vec2d{X: tt.a[2], Y: tt.a[3]}
			l2A := galx.Vec2d{X: tt.b[0], Y: tt.b[1]}
			l2B := galx.Vec2d{X: tt.b[2], Y: tt.b[3]}

			var variants = make([][2]galx.Line, 0)

			variants = append(variants, [2]galx.Line{
				{A: l1A, B: l1B}, {A: l2A, B: l2B},
			})
			variants = append(variants, [2]galx.Line{
				{A: l1B, B: l1A}, {A: l2B, B: l2A},
			})
			variants = append(variants, [2]galx.Line{
				{A: l2A, B: l2B}, {A: l1A, B: l1B},
			})
			variants = append(variants, [2]galx.Line{
				{A: l2B, B: l2A}, {A: l1B, B: l1A},
			})

			for varId, variant := range variants {
				l1 := variant[0]
				l2 := variant[1]

				if got := Line2Line(l1, l2); got != tt.want {
					t.Errorf("Line2Line() = %v, want %v\nVariant: #%v\n%+v\n%+v",
						got,
						tt.want,
						varId+1,
						l1,
						l2,
					)
				}
			}
		})
	}
}

func TestLine2Circle(t *testing.T) {
	type args struct {
		line   [4]float64
		circle galx.Circle
	}

	center2R := galx.Circle{
		Pos:    galx.Vec2d{},
		Radius: 2,
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "zero collide",
			args: args{
				line:   [4]float64{0, 0, 0, 0},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "h collide",
			args: args{
				line:   [4]float64{-1, 0, 1, 0},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "h2x collide",
			args: args{
				line:   [4]float64{-2, 0, 2, 0},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "h3x collide",
			args: args{
				line:   [4]float64{-3, 0, 3, 0},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "v collide",
			args: args{
				line:   [4]float64{0, -1, 0, 1},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "v2x collide",
			args: args{
				line:   [4]float64{0, -2, 0, 2},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "v3x collide",
			args: args{
				line:   [4]float64{0, -3, 0, 3},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "top collide",
			args: args{
				line:   [4]float64{-5, -2, 5, -2},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "top -offset collide",
			args: args{
				line:   [4]float64{-5, -1.9999, 5, -1.9999},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "top +offset no collide",
			args: args{
				line:   [4]float64{-5, -2.0001, 5, -2.0001},
				circle: center2R,
			},
			want: false,
		},
		{
			name: "bottom collide",
			args: args{
				line:   [4]float64{-5, 2, 5, 2},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "bottom -offset collide",
			args: args{
				line:   [4]float64{-5, 1.9999, 5, 1.9999},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "bottom +offset no collide",
			args: args{
				line:   [4]float64{-5, 2.0001, 5, 2.0001},
				circle: center2R,
			},
			want: false,
		},
		{
			name: "left collide",
			args: args{
				line:   [4]float64{-2, -10, -2, 10},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "left offset no collide",
			args: args{
				line:   [4]float64{-2.0001, -10, -2.0001, 10},
				circle: center2R,
			},
			want: false,
		},
		{
			name: "right offset collide",
			args: args{
				line:   [4]float64{2, 0, 15, 0},
				circle: center2R,
			},
			want: true,
		},
		{
			name: "right offset no collide",
			args: args{
				line:   [4]float64{2.0001, 0, 15, 0},
				circle: center2R,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lA := galx.Vec2d{X: tt.args.line[0], Y: tt.args.line[1]}
			lB := galx.Vec2d{X: tt.args.line[2], Y: tt.args.line[3]}

			variants := []galx.Line{
				{A: lA, B: lB},
				{A: lB, B: lA},
			}

			for varId, variant := range variants {
				if got := Line2Circle(variant, tt.args.circle); got != tt.want {
					t.Errorf("Line2Circle() = %v, want %v\nVariant: #%v\n%v\n%v",
						got,
						tt.want,
						varId+1,
						tt.args.circle,
						variant,
					)
				}
			}
		})
	}
}
