package engine

import (
	"reflect"
	"testing"
)

func TestRadian(t *testing.T) {
	type args struct {
		angle Angle
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0",
			args: args{
				angle: Anglef(0),
			},
			want: 0,
		},
		{
			name: "360",
			args: args{
				angle: Anglef(360),
			},
			want: 0,
		},
		{
			name: "240",
			args: args{
				angle: Anglef(240),
			},
			want: 4.1887902047863905,
		},
		{
			name: "125",
			args: args{
				angle: Anglef(125),
			},
			want: 2.1816615649929116,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Radian(tt.args.angle); got != tt.want {
				t.Errorf("Radian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Add(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		n Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector2D
	}{
		{
			name: "plus",
			fields: fields{
				X: 4,
				Y: 3,
			},
			args: args{
				n: Vector2D{
					X: 1,
					Y: 1,
				},
			},
			want: Vector2D{
				X: 5,
				Y: 4,
			},
		},
		{
			name: "minus",
			fields: fields{
				X: 9,
				Y: 6,
			},
			args: args{
				n: Vector2D{
					X: -3,
					Y: -2,
				},
			},
			want: Vector2D{
				X: 6,
				Y: 4,
			},
		},
		{
			name: "comb",
			fields: fields{
				X: 25.2,
				Y: 14.1,
			},
			args: args{
				n: Vector2D{
					X: 3.25,
					Y: -1.85,
				},
			},
			want: Vector2D{
				X: 28.450,
				Y: 12.25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Add(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_AngleBetween(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		to Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Angle
	}{
		{
			name: "base",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 1,
					Y: 1,
				},
			},
			want: 0,
		},
		{
			name: "reversed",
			fields: fields{
				X: 5,
				Y: 5,
			},
			args: args{
				to: Vector2D{
					X: -5,
					Y: -3,
				},
			},
			want: 165.9637565320735,
		},
		{
			name: "simple",
			fields: fields{
				X: 10,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 0,
					Y: 10,
				},
			},
			want: 90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.AngleBetween(tt.args.to); got != tt.want {
				t.Errorf("AngleBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Cross(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		to Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "simple",
			fields: fields{
				X: 66,
				Y: -90,
			},
			args: args{
				to: Vector2D{
					X: 147,
					Y: -63,
				},
			},
			want: 9072,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Cross(tt.args.to); got != tt.want {
				t.Errorf("Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Direction(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Angle
	}{
		{
			name: "1.5h",
			fields: fields{
				X: 1,
				Y: -1,
			},
			want: 315,
		},
		{
			name: "left",
			fields: fields{
				X: -1,
				Y: 0,
			},
			want: 180,
		},
		{
			name: "right",
			fields: fields{
				X: 1,
				Y: 0,
			},
			want: 0,
		},
		{
			name: "top",
			fields: fields{
				X: 0,
				Y: -1,
			},
			want: 90,
		},
		{
			name: "bottom",
			fields: fields{
				X: 0,
				Y: 1,
			},
			want: 270,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Direction(); got != tt.want {
				t.Errorf("Direction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_DistanceTo(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		to Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "positive x",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 5,
					Y: 0,
				},
			},
			want: 5,
		},
		{
			name: "negative x",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: -4,
					Y: 0,
				},
			},
			want: 4,
		},
		{
			name: "positive y",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 0,
					Y: 3,
				},
			},
			want: 3,
		},
		{
			name: "xx",
			fields: fields{
				X: 10,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 20,
					Y: 0,
				},
			},
			want: 10,
		},
		{
			name: "n to p",
			fields: fields{
				X: -10,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 10,
					Y: 0,
				},
			},
			want: 20,
		},
		{
			name: "n to p",
			fields: fields{
				X: 1,
				Y: 1,
			},
			args: args{
				to: Vector2D{
					X: 2,
					Y: 2,
				},
			},
			want: 1.4142135623730951,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.DistanceTo(tt.args.to); got != tt.want {
				t.Errorf("DistanceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Divide(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector2D
	}{
		{
			name: "simple",
			fields: fields{
				X: 8,
				Y: 4,
			},
			args: args{
				n: 2,
			},
			want: Vector2D{
				X: 4,
				Y: 2,
			},
		},
		{
			name: "negative",
			fields: fields{
				X: 8,
				Y: 4,
			},
			args: args{
				n: -2,
			},
			want: Vector2D{
				X: -4,
				Y: -2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Divide(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Dot(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		to Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "simple",
			fields: fields{
				X: 1,
				Y: 2,
			},
			args: args{
				to: Vector2D{
					X: 3,
					Y: 4,
				},
			},
			want: 11,
		},
		{
			name: "negative",
			fields: fields{
				X: -1,
				Y: -4,
			},
			args: args{
				to: Vector2D{
					X: 4,
					Y: 1,
				},
			},
			want: -8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Dot(tt.args.to); got != tt.want {
				t.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Length(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "simple",
			fields: fields{
				X: 5,
				Y: 0,
			},
			want: 5,
		},
		{
			name: "simple",
			fields: fields{
				X: -5,
				Y: 0,
			},
			want: 5,
		},
		{
			name: "simple",
			fields: fields{
				X: -5,
				Y: 5,
			},
			want: 7.0710678118654755,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Mul(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector2D
	}{
		{
			name: "simple",
			fields: fields{
				X: 2,
				Y: 3,
			},
			args: args{
				n: 3,
			},
			want: Vector2D{
				X: 6,
				Y: 9,
			},
		},
		{
			name: "negative",
			fields: fields{
				X: -2,
				Y: 5,
			},
			args: args{
				n: 2,
			},
			want: Vector2D{
				X: -4,
				Y: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Mul(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Normalize(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector2D
	}{
		{
			name: "only x",
			fields: fields{
				X: 1,
				Y: 0,
			},
			want: Vector2D{
				X: 1,
				Y: 0,
			},
		},
		{
			name: "only x inv",
			fields: fields{
				X: -1,
				Y: 0,
			},
			want: Vector2D{
				X: -1,
				Y: 0,
			},
		},
		{
			name: "only y",
			fields: fields{
				X: 0,
				Y: -1,
			},
			want: Vector2D{
				X: 0,
				Y: -1,
			},
		},
		{
			name: "pos",
			fields: fields{
				X: 1,
				Y: 1,
			},
			want: Vector2D{
				X: 0.7071067811865475,
				Y: 0.7071067811865475,
			},
		},
		{
			name: "pos 2x",
			fields: fields{
				X: 10,
				Y: 10,
			},
			want: Vector2D{
				X: 0.7071067811865475,
				Y: 0.7071067811865475,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_PolarOffset(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		distance float64
		angle    Angle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector2D
	}{
		{
			name: "x right",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				distance: 5,
				angle:    0,
			},
			want: Vector2D{
				X: 5,
				Y: 0,
			},
		},
		{
			name: "x left",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				distance: -5,
				angle:    0,
			},
			want: Vector2D{
				X: -5,
				Y: 0,
			},
		},
		{
			name: "y top",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				distance: 2,
				angle:    90,
			},
			want: Vector2D{
				X: 0,
				Y: -2,
			},
		},
		{
			name: "y bottom",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				distance: 2,
				angle:    270,
			},
			want: Vector2D{
				X: 0,
				Y: 2,
			},
		},
		{
			name: "y bottom offset",
			fields: fields{
				X: 0,
				Y: -2,
			},
			args: args{
				distance: 2,
				angle:    270,
			},
			want: Vector2D{
				X: 0,
				Y: 0,
			},
		},
		{
			name: "y bottom offset neg",
			fields: fields{
				X: 0,
				Y: -2,
			},
			args: args{
				distance: -2,
				angle:    270,
			},
			want: Vector2D{
				X: 0,
				Y: -4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.PolarOffset(tt.args.distance, tt.args.angle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PolarOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Rotate(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		angle Angle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector2D
	}{
		{
			name: "turn to zero",
			fields: fields{
				X: 1,
				Y: 1,
			},
			args: args{
				angle: -45,
			},
			want: Vector2D{
				X: 1.41421,
				Y: 0,
			},
		},
		{
			name: "counter-clockwise",
			fields: fields{
				X: 1,
				Y: 0,
			},
			args: args{
				angle: 90,
			},
			want: Vector2D{
				X: 0,
				Y: -1,
			},
		},
		{
			name: "clockwise",
			fields: fields{
				X: 1,
				Y: 0,
			},
			args: args{
				angle: -90,
			},
			want: Vector2D{
				X: 0,
				Y: 1,
			},
		},
		{
			name: "clockwise to 12h",
			fields: fields{
				X: -1,
				Y: 0,
			},
			args: args{
				angle: -90,
			},
			want: Vector2D{
				X: 0,
				Y: -1,
			},
		},
		{
			name: "complex",
			fields: fields{
				X: 5,
				Y: 5,
			},
			args: args{
				angle: -45,
			},
			want: Vector2D{
				X: 7.07107,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Rotate(tt.args.angle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_RotateAround(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		orig  Vector2D
		angle Angle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector2D
	}{
		{
			name: "simple",
			fields: fields{
				X: 1,
				Y: 0,
			},
			args: args{
				orig: Vector2D{
					X: 0,
					Y: 0,
				},
				angle: 180,
			},
			want: Vector2D{
				X: -1,
				Y: 0,
			},
		},
		{
			name: "simple neg",
			fields: fields{
				X: -1,
				Y: 0,
			},
			args: args{
				orig: Vector2D{
					X: 0,
					Y: 0,
				},
				angle: 180,
			},
			want: Vector2D{
				X: 1,
				Y: 0,
			},
		},
		{
			name: "simple neg",
			fields: fields{
				X: 0,
				Y: -1,
			},
			args: args{
				orig: Vector2D{
					X: 0,
					Y: 0,
				},
				angle: Anglef(-90),
			},
			want: Vector2D{
				X: 1,
				Y: 0,
			},
		},
		{
			name: "simple neg x",
			fields: fields{
				X: 0,
				Y: 1,
			},
			args: args{
				orig: Vector2D{
					X: 0,
					Y: 0,
				},
				angle: -90,
			},
			want: Vector2D{
				X: -1,
				Y: 0,
			},
		},
		{
			name: "simple half circle",
			fields: fields{
				X: -2,
				Y: 0,
			},
			args: args{
				orig: Vector2D{
					X: 0,
					Y: 0,
				},
				angle: 180,
			},
			want: Vector2D{
				X: 2,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.RotateAround(tt.args.orig, tt.args.angle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateAround() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_RoundX(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "simple",
			fields: fields{
				X: 2.2567788,
				Y: 1.123456,
			},
			want: 2,
		},
		{
			name: "lower always",
			fields: fields{
				X: 2.9567788,
				Y: 1.123456,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.RoundX(); got != tt.want {
				t.Errorf("RoundX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_RoundY(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "simple",
			fields: fields{
				X: 2.2567788,
				Y: 1.123456,
			},
			want: 1,
		},
		{
			name: "lower always",
			fields: fields{
				X: 2.9567788,
				Y: 1.99923456,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.RoundY(); got != tt.want {
				t.Errorf("RoundY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_ToPoint(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Point
	}{
		{
			name: "simple",
			fields: fields{
				X: 2,
				Y: 2,
			},
			want: Point{
				X: 2,
				Y: 2,
			},
		},
		{
			name: "always lower",
			fields: fields{
				X: 2.01,
				Y: 2.99,
			},
			want: Point{
				X: 2,
				Y: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.ToPoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorForward(t *testing.T) {
	type args struct {
		y float64
	}
	tests := []struct {
		name string
		args args
		want Vector2D
	}{
		{
			name: "simple",
			args: args{
				y: 2,
			},
			want: Vector2D{
				X: 0,
				Y: -2,
			},
		},
		{
			name: "simple 2",
			args: args{
				y: -2,
			},
			want: Vector2D{
				X: 0,
				Y: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VectorForward(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VectorForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorLeft(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want Vector2D
	}{
		{
			name: "simple",
			args: args{
				x: 2.01,
			},
			want: Vector2D{
				X: -2.01,
				Y: 0,
			},
		},
		{
			name: "simple 2",
			args: args{
				x: 2,
			},
			want: Vector2D{
				X: -2,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VectorLeft(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VectorLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorSum(t *testing.T) {
	type args struct {
		list []Vector2D
	}
	tests := []struct {
		name string
		args args
		want Vector2D
	}{
		{
			name: "all pos",
			args: args{
				list: []Vector2D{
					{
						X: 1,
						Y: 1,
					},
					{
						X: 2,
						Y: 2,
					},
					{
						X: 3,
						Y: 3,
					},
				},
			},
			want: Vector2D{
				X: 6,
				Y: 6,
			},
		},
		{
			name: "all neg",
			args: args{
				list: []Vector2D{
					{
						X: -1,
						Y: -1,
					},
					{
						X: -2,
						Y: -2,
					},
					{
						X: -3,
						Y: -3,
					},
				},
			},
			want: Vector2D{
				X: -6,
				Y: -6,
			},
		},
		{
			name: "zero final",
			args: args{
				list: []Vector2D{
					{
						X: 1,
						Y: -1,
					},
					{
						X: -1,
						Y: 1,
					},
					{
						X: -2,
						Y: 2,
					},
					{
						X: 2,
						Y: -2,
					},
				},
			},
			want: Vector2D{
				X: 0,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VectorSum(tt.args.list...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VectorSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorTowards(t *testing.T) {
	type args struct {
		a Angle
	}
	tests := []struct {
		name string
		args args
		want Vector2D
	}{
		{
			name: "right",
			args: args{
				a: 0,
			},
			want: Vector2D{
				X: 1,
				Y: 0,
			},
		},
		{
			name: "top",
			args: args{
				a: 90,
			},
			want: Vector2D{
				X: 0,
				Y: -1,
			},
		},
		{
			name: "left",
			args: args{
				a: 180,
			},
			want: Vector2D{
				X: -1,
				Y: 0,
			},
		},
		{
			name: "bottom",
			args: args{
				a: 270,
			},
			want: Vector2D{
				X: 0,
				Y: 1,
			},
		},
		{
			name: "bottom",
			args: args{
				a: 45,
			},
			want: Vector2D{
				X: 0.70711,
				Y: -0.70711,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VectorTowards(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VectorTowards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_AngleTo(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		to Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Angle
	}{
		{
			name: "zero in 3h",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 1,
					Y: 0,
				},
			},
			want: 0,
		},
		{
			name: "top on 12h",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 0,
					Y: -1,
				},
			},
			want: 90,
		},
		{
			name: "left at 9h",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: -1,
					Y: 0,
				},
			},
			want: 180,
		},
		{
			name: "bottom at 6h",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 0,
					Y: 1,
				},
			},
			want: 270,
		},
		{
			name: "real",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: 1,
					Y: -1,
				},
			},
			want: 45,
		},
		{
			name: "real negative",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vector2D{
					X: -1,
					Y: 1,
				},
			},
			want: 225,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.AngleTo(tt.args.to); got != tt.want {
				t.Errorf("AngleTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
