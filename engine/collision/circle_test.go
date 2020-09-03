package collision

import (
	"testing"

	"github.com/fe3dback/galaxy/engine"
)

func TestCircle2Circle(t *testing.T) {
	type args struct {
		a engine.Circle
		b engine.Circle
	}

	center2R := engine.Circle{
		Pos:    engine.Vec{X: 0, Y: 0},
		Radius: 2,
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "inside collide", args: args{
			a: center2R,
			b: engine.Circle{Pos: engine.Vec{X: 0, Y: 0}, Radius: 1},
		}, want: true},
		{name: "outside collide", args: args{
			a: center2R,
			b: engine.Circle{Pos: engine.Vec{X: 0, Y: 0}, Radius: 3},
		}, want: true},

		// sides collide
		{name: "top side collide", args: args{
			a: center2R,
			b: engine.Circle{Pos: engine.Vec{X: 0, Y: -3}, Radius: 1},
		}, want: true},
		{name: "bottom side collide", args: args{
			a: center2R,
			b: engine.Circle{Pos: engine.Vec{X: 0, Y: 3}, Radius: 1},
		}, want: true},
		{name: "left side collide", args: args{
			a: center2R,
			b: engine.Circle{Pos: engine.Vec{X: -3, Y: 0}, Radius: 1},
		}, want: true},
		{name: "right side collide", args: args{
			a: center2R,
			b: engine.Circle{Pos: engine.Vec{X: 3, Y: 0}, Radius: 1},
		}, want: true},

		// sides offset not collide
		{name: "top side collide", args: args{
			a: center2R,
			b: engine.Circle{Pos: engine.Vec{X: 0, Y: -3.0001}, Radius: 1},
		}, want: false},

		// special
		{name: "eq collide", args: args{
			a: engine.Circle{Pos: engine.Vec{X: 5, Y: 1}, Radius: 0.5},
			b: engine.Circle{Pos: engine.Vec{X: 6, Y: 1}, Radius: 0.5},
		}, want: true},
		{name: "small offset no collide", args: args{
			a: engine.Circle{Pos: engine.Vec{X: 5, Y: 1}, Radius: 0.49999},
			b: engine.Circle{Pos: engine.Vec{X: 6, Y: 1}, Radius: 0.5},
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Circle2Circle(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Circle2Circle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle2Point(t *testing.T) {
	type args struct {
		c engine.Circle
		p engine.Vec
	}

	center2R := engine.Circle{
		Pos: engine.Vec{
			X: 0,
			Y: 0,
		},
		Radius: 2,
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s1",
			args: args{
				c: center2R,
				p: engine.Vec{
					X: 0,
					Y: 0,
				},
			},
			want: true,
		},
		{
			name: "left",
			args: args{
				c: center2R,
				p: engine.Vec{
					X: -2,
					Y: 0,
				},
			},
			want: true,
		},
		{
			name: "top",
			args: args{
				c: center2R,
				p: engine.Vec{
					X: 0,
					Y: -2,
				},
			},
			want: true,
		},
		{
			name: "left no collide",
			args: args{
				c: center2R,
				p: engine.Vec{
					X: 0,
					Y: -2.0001,
				},
			},
			want: false,
		},
		{
			name: "bottom offset collide",
			args: args{
				c: engine.Circle{
					Pos: engine.Vec{
						X: 0,
						Y: 10,
					},
					Radius: 3,
				},
				p: engine.Vec{
					X: 0,
					Y: 13,
				},
			},
			want: true,
		},
		{
			name: "bottom offset+ collide",
			args: args{
				c: engine.Circle{
					Pos: engine.Vec{
						X: 0,
						Y: 10,
					},
					Radius: 3,
				},
				p: engine.Vec{
					X: 0,
					Y: 12.99999,
				},
			},
			want: true,
		},
		{
			name: "bottom offset- no collide",
			args: args{
				c: engine.Circle{
					Pos: engine.Vec{
						X: 0,
						Y: 10,
					},
					Radius: 3,
				},
				p: engine.Vec{
					X: 0,
					Y: 13.00001,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Circle2Point(tt.args.c, tt.args.p); got != tt.want {
				t.Errorf("Circle2Point() = %v, want %v", got, tt.want)
			}
		})
	}
}
