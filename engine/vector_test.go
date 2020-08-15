package engine

import (
	"reflect"
	"testing"
)

func testNormAngle(angle Angle) float64 {
	return NewAngle(RoundTo(clampDeg(angle.Degrees()))).Radians()
}

func testNormVec(vec Vec) Vec {
	return Vec{
		X: RoundTo(vec.X),
		Y: RoundTo(vec.Y),
	}
}

func TestRadian(t *testing.T) {
	type args struct {
		angle Angle
	}
	tests := []struct {
		name string
		args args
		want Angle
	}{
		{
			name: "0",
			args: args{
				angle: NewAngle(0),
			},
			want: Angle(0),
		},
		{
			name: "360",
			args: args{
				angle: NewAngle(360),
			},
			want: Angle(0),
		},
		{
			name: "240",
			args: args{
				angle: NewAngle(240),
			},
			want: Angle(4.1887902047863905),
		},
		{
			name: "125",
			args: args{
				angle: NewAngle(125),
			},
			want: Angle(2.181661564992912),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.angle; got != tt.want {
				t.Errorf("deg2rad() = %v, want %v", got, tt.want)
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
		n Vec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec
	}{
		{
			name: "plus",
			fields: fields{
				X: 4,
				Y: 3,
			},
			args: args{
				n: Vec{
					X: 1,
					Y: 1,
				},
			},
			want: Vec{
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
				n: Vec{
					X: -3,
					Y: -2,
				},
			},
			want: Vec{
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
				n: Vec{
					X: 3.25,
					Y: -1.85,
				},
			},
			want: Vec{
				X: 28.450,
				Y: 12.25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
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
		to Vec
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
				to: Vec{
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
				to: Vec{
					X: -5,
					Y: -3,
				},
			},
			want: NewAngle(165.9637565320735),
		},
		{
			name: "simple",
			fields: fields{
				X: 10,
				Y: 0,
			},
			args: args{
				to: Vec{
					X: 0,
					Y: 10,
				},
			},
			want: NewAngle(90),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.AngleBetween(tt.args.to); RoundTo(float64(got)) != RoundTo(float64(tt.want)) {
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
		to Vec
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
				to: Vec{
					X: 147,
					Y: -63,
				},
			},
			want: 9072,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
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
			name: "7.5h",
			fields: fields{
				X: -1,
				Y: 1,
			},
			want: NewAngle(225),
		},
		{
			name: "1.5h",
			fields: fields{
				X: 1,
				Y: -1,
			},
			want: NewAngle(45),
		},
		{
			name: "left",
			fields: fields{
				X: -1,
				Y: 0,
			},
			want: NewAngle(180),
		},
		{
			name: "right",
			fields: fields{
				X: 1,
				Y: 0,
			},
			want: NewAngle(0),
		},
		{
			name: "top",
			fields: fields{
				X: 0,
				Y: -1,
			},
			want: NewAngle(90),
		},
		{
			name: "bottom",
			fields: fields{
				X: 0,
				Y: 1,
			},
			want: NewAngle(270),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Direction(); testNormAngle(got) != testNormAngle(tt.want) {
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
		to Vec
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
				to: Vec{
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
				to: Vec{
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
				to: Vec{
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
				to: Vec{
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
				to: Vec{
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
				to: Vec{
					X: 2,
					Y: 2,
				},
			},
			want: 1.4142135623730951,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
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
		want   Vec
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
			want: Vec{
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
			want: Vec{
				X: -4,
				Y: -2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Decrease(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrease() = %v, want %v", got, tt.want)
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
		to Vec
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
				to: Vec{
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
				to: Vec{
					X: 4,
					Y: 1,
				},
			},
			want: -8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
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
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Magnitude(); got != tt.want {
				t.Errorf("Magnitude() = %v, want %v", got, tt.want)
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
		want   Vec
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
			want: Vec{
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
			want: Vec{
				X: -4,
				Y: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Scale(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale() = %v, want %v", got, tt.want)
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
		want   Vec
	}{
		{
			name: "only x",
			fields: fields{
				X: 1,
				Y: 0,
			},
			want: Vec{
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
			want: Vec{
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
			want: Vec{
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
			want: Vec{
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
			want: Vec{
				X: 0.7071067811865475,
				Y: 0.7071067811865475,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
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
		want   Vec
	}{
		{
			name: "x right",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				distance: 5,
				angle:    NewAngle(0),
			},
			want: Vec{
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
				angle:    NewAngle(0),
			},
			want: Vec{
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
				angle:    NewAngle(90),
			},
			want: Vec{
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
				angle:    NewAngle(270),
			},
			want: Vec{
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
				angle:    NewAngle(270),
			},
			want: Vec{
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
				angle:    NewAngle(270),
			},
			want: Vec{
				X: 0,
				Y: -4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			got := v.PolarOffset(tt.args.distance, tt.args.angle)
			got.X = RoundTo(got.X)
			got.Y = RoundTo(got.Y)

			if !reflect.DeepEqual(got, tt.want) {
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
		want   Vec
	}{
		{
			name: "turn to zero",
			fields: fields{
				X: 1,
				Y: 1,
			},
			args: args{
				angle: NewAngle(-45),
			},
			want: Vec{
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
				angle: NewAngle(90),
			},
			want: Vec{
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
				angle: NewAngle(-90),
			},
			want: Vec{
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
				angle: NewAngle(-90),
			},
			want: Vec{
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
				angle: NewAngle(-45),
			},
			want: Vec{
				X: 7.07107,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}

			got := v.Rotate(tt.args.angle)
			got.X = RoundTo(got.X)
			got.Y = RoundTo(got.Y)
			tt.want.X = RoundTo(tt.want.X)
			tt.want.Y = RoundTo(tt.want.Y)

			if !reflect.DeepEqual(got, tt.want) {
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
		orig  Vec
		angle Angle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vec
	}{
		{
			name: "simple",
			fields: fields{
				X: 1,
				Y: 0,
			},
			args: args{
				orig: Vec{
					X: 0,
					Y: 0,
				},
				angle: NewAngle(180),
			},
			want: Vec{
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
				orig: Vec{
					X: 0,
					Y: 0,
				},
				angle: NewAngle(180),
			},
			want: Vec{
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
				orig: Vec{
					X: 0,
					Y: 0,
				},
				angle: NewAngle(-90),
			},
			want: Vec{
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
				orig: Vec{
					X: 0,
					Y: 0,
				},
				angle: NewAngle(-90),
			},
			want: Vec{
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
				orig: Vec{
					X: 0,
					Y: 0,
				},
				angle: NewAngle(180),
			},
			want: Vec{
				X: 2,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.RotateAround(tt.args.orig, tt.args.angle); !reflect.DeepEqual(testNormVec(got), testNormVec(tt.want)) {
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
			v := Vec{
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
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.RoundY(); got != tt.want {
				t.Errorf("RoundY() = %v, want %v", got, tt.want)
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
		want Vec
	}{
		{
			name: "simple",
			args: args{
				y: 2,
			},
			want: Vec{
				X: 0,
				Y: -2,
			},
		},
		{
			name: "simple 2",
			args: args{
				y: -2,
			},
			want: Vec{
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
		want Vec
	}{
		{
			name: "simple",
			args: args{
				x: 2.01,
			},
			want: Vec{
				X: -2.01,
				Y: 0,
			},
		},
		{
			name: "simple 2",
			args: args{
				x: 2,
			},
			want: Vec{
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
		list []Vec
	}
	tests := []struct {
		name string
		args args
		want Vec
	}{
		{
			name: "all pos",
			args: args{
				list: []Vec{
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
			want: Vec{
				X: 6,
				Y: 6,
			},
		},
		{
			name: "all neg",
			args: args{
				list: []Vec{
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
			want: Vec{
				X: -6,
				Y: -6,
			},
		},
		{
			name: "zero final",
			args: args{
				list: []Vec{
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
			want: Vec{
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
		want Vec
	}{
		{
			name: "right",
			args: args{
				a: 0,
			},
			want: Vec{
				X: 1,
				Y: 0,
			},
		},
		{
			name: "top",
			args: args{
				a: NewAngle(90),
			},
			want: Vec{
				X: 0,
				Y: -1,
			},
		},
		{
			name: "left",
			args: args{
				a: NewAngle(180),
			},
			want: Vec{
				X: -1,
				Y: 0,
			},
		},
		{
			name: "bottom",
			args: args{
				a: NewAngle(270),
			},
			want: Vec{
				X: 0,
				Y: 1,
			},
		},
		{
			name: "bottom",
			args: args{
				a: NewAngle(45),
			},
			want: Vec{
				X: 0.70711,
				Y: -0.70711,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VectorTowards(tt.args.a); !reflect.DeepEqual(testNormVec(got), testNormVec(tt.want)) {
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
		to Vec
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
				to: Vec{
					X: 1,
					Y: 0,
				},
			},
			want: NewAngle(0),
		},
		{
			name: "top on 12h",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vec{
					X: 0,
					Y: -1,
				},
			},
			want: NewAngle(90),
		},
		{
			name: "left at 9h",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vec{
					X: -1,
					Y: 0,
				},
			},
			want: NewAngle(180),
		},
		{
			name: "bottom at 6h",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vec{
					X: 0,
					Y: 1,
				},
			},
			want: NewAngle(270),
		},
		{
			name: "real",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vec{
					X: 1,
					Y: -1,
				},
			},
			want: NewAngle(45),
		},
		{
			name: "real negative",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				to: Vec{
					X: -1,
					Y: 1,
				},
			},
			want: NewAngle(225),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.AngleTo(tt.args.to); testNormAngle(got) != testNormAngle(tt.want) {
				t.Errorf("AngleTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
