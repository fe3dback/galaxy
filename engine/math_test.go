package engine

import (
	"math"
	"testing"
)

func Test_roundTo(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "zero",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "s 1",
			args: args{
				n: 0.11,
			},
			want: 0.11,
		},
		{
			name: "s 2",
			args: args{
				n: 0.021,
			},
			want: 0.021,
		},
		{
			name: "s 3",
			args: args{
				n: 0.031,
			},
			want: 0.031,
		},
		{
			name: "s 4",
			args: args{
				n: 0.0041,
			},
			want: 0.0041,
		},
		{
			name: "s 4 ok floor",
			args: args{
				n: 0.00412,
			},
			want: 0.0041,
		},
		{
			name: "s 4 ok ceil",
			args: args{
				n: 0.0041901,
			},
			want: 0.0042,
		},
		{
			name: "s 4 ok ceil z",
			args: args{
				n: 0.00419999999,
			},
			want: 0.004200,
		},
		{
			name: "pi",
			args: args{
				n: math.Pi,
			},
			want: 3.1416,
		},
		{
			name: "normal number",
			args: args{
				n: 15,
			},
			want: 15,
		},
		{
			name: "neg number",
			args: args{
				n: -math.Pi,
			},
			want: -3.141600,
		},
		{
			name: "real 1",
			args: args{
				n: 0.19999999,
			},
			want: 0.2000000,
		},
		{
			name: "real 2",
			args: args{
				n: 0.33333333333,
			},
			want: 0.3333000000,
		},
		{
			name: "real 3",
			args: args{
				n: 0.10000001,
			},
			want: 0.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundTo(tt.args.n); got != tt.want {
				t.Errorf("RoundTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLerp(t *testing.T) {
	type args struct {
		a float64
		b float64
		t float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "zero",
			args: args{
				a: 0,
				b: 1,
				t: 0,
			},
			want: 0,
		},
		{
			name: "full",
			args: args{
				a: 0,
				b: 1,
				t: 1,
			},
			want: 1,
		},
		{
			name: "less zero",
			args: args{
				a: 0,
				b: 1,
				t: -10,
			},
			want: 0,
		},
		{
			name: "more full",
			args: args{
				a: 0,
				b: 1,
				t: 10,
			},
			want: 1,
		},
		{
			name: "complex 1",
			args: args{
				a: 5,
				b: 15,
				t: 1,
			},
			want: 15,
		},
		{
			name: "complex 2",
			args: args{
				a: 5,
				b: 15,
				t: 0,
			},
			want: 5,
		},
		{
			name: "float 1",
			args: args{
				a: 0,
				b: 1,
				t: 0.5,
			},
			want: 0.5,
		},
		{
			name: "float 2",
			args: args{
				a: 0,
				b: 1,
				t: 0.75,
			},
			want: 0.75,
		},
		{
			name: "float 3",
			args: args{
				a: 0,
				b: 1,
				t: 0.99,
			},
			want: 0.99,
		},
		{
			name: "complex float 1",
			args: args{
				a: 5,
				b: 10,
				t: 0.5,
			},
			want: 7.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lerp(tt.args.a, tt.args.b, tt.args.t); got != tt.want {
				t.Errorf("Lerp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	type args struct {
		n   float64
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "simple",
			args: args{
				n:   50,
				min: 0,
				max: 100,
			},
			want: 50,
		},
		{
			name: "s1",
			args: args{
				n:   150,
				min: 0,
				max: 100,
			},
			want: 100,
		},
		{
			name: "s2",
			args: args{
				n:   -50,
				min: 0,
				max: 100,
			},
			want: 0,
		},
		{
			name: "complex",
			args: args{
				n:   -25,
				min: -100,
				max: 100,
			},
			want: -25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.args.n, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLerpf(t *testing.T) {
	type args struct {
		originA float64
		originB float64
		targetA float64
		targetB float64
		origin  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "min",
			args: args{
				originA: 0,
				originB: 1,
				targetA: 5,
				targetB: 10,
				origin:  0,
			},
			want: 5,
		},
		{
			name: "max",
			args: args{
				originA: 0,
				originB: 1,
				targetA: 5,
				targetB: 10,
				origin:  1,
			},
			want: 10,
		},
		{
			name: "n min",
			args: args{
				originA: 5,
				originB: 10,
				targetA: 0,
				targetB: 1,
				origin:  5,
			},
			want: 0,
		},
		{
			name: "n max",
			args: args{
				originA: 5,
				originB: 10,
				targetA: 0,
				targetB: 1,
				origin:  10,
			},
			want: 1,
		},
		{
			name: "resize",
			args: args{
				originA: 0,
				originB: 1,
				targetA: 0,
				targetB: 1024,
				origin:  0.5,
			},
			want: 512,
		},
		{
			name: "resize back",
			args: args{
				originA: 0,
				originB: 1024,
				targetA: 0,
				targetB: 1,
				origin:  512,
			},
			want: 0.5,
		},
		{
			name: "negative resize back",
			args: args{
				originA: 0,
				originB: -1024,
				targetA: 0,
				targetB: 1,
				origin:  -512,
			},
			want: 0.5,
		},
		{
			name: "negative resize back full",
			args: args{
				originA: 0,
				originB: -1024,
				targetA: 0,
				targetB: 1,
				origin:  -1024,
			},
			want: 1,
		},
		{
			name: "same",
			args: args{
				originA: 1,
				originB: 1,
				targetA: 1,
				targetB: 1,
				origin:  1,
			},
			want: 1,
		},
		{
			name: "denormalize A",
			args: args{
				originA: 0,
				originB: 1,
				targetA: -1,
				targetB: 1,
				origin:  0.5,
			},
			want: 0,
		},
		{
			name: "denormalize min",
			args: args{
				originA: 0,
				originB: 1,
				targetA: -1,
				targetB: 1,
				origin:  0,
			},
			want: -1,
		},
		{
			name: "denormalize max",
			args: args{
				originA: 0,
				originB: 1,
				targetA: -1,
				targetB: 1,
				origin:  1,
			},
			want: 1,
		},
		{
			name: "reversed x min",
			args: args{
				originA: -1,
				originB: 1,
				targetA: 1,
				targetB: 0,
				origin:  1,
			},
			want: 0,
		},
		{
			name: "reversed x max",
			args: args{
				originA: -1,
				originB: 1,
				targetA: 1,
				targetB: 0,
				origin:  -1,
			},
			want: 1,
		},
		{
			name: "reversed x center",
			args: args{
				originA: -1,
				originB: 1,
				targetA: 1,
				targetB: 0,
				origin:  0,
			},
			want: 0.5,
		},
		{
			name: "reversed x bigger",
			args: args{
				originA: -1,
				originB: 1,
				targetA: 1,
				targetB: 0,
				origin:  0.5,
			},
			want: 0.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lerpf(tt.args.originA, tt.args.originB, tt.args.targetA, tt.args.targetB, tt.args.origin); got != tt.want {
				t.Errorf("Lerpf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLerpInverse(t *testing.T) {
	type args struct {
		a float64
		b float64
		t float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "simple",
			args: args{
				a: 0,
				b: 2,
				t: 1,
			},
			want: 0.5,
		},
		{
			name: "simple negative",
			args: args{
				a: -1,
				b: 1,
				t: 0,
			},
			want: 0.5,
		},
		{
			name: "scaled 1",
			args: args{
				a: 0,
				b: 1024,
				t: 768,
			},
			want: 0.75,
		},
		{
			name: "scaled 2",
			args: args{
				a: 0,
				b: 1024,
				t: 256,
			},
			want: 0.25,
		},
		{
			name: "scaled negative",
			args: args{
				a: -1024,
				b: 1024,
				t: 512,
			},
			want: 0.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LerpInverse(tt.args.a, tt.args.b, tt.args.t); got != tt.want {
				t.Errorf("LerpInverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomRange(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "s1",
			args: args{
				min: 1,
				max: 2,
			},
		},
		{
			name: "s2",
			args: args{
				min: 1,
				max: 1,
			},
		},
		{
			name: "s3",
			args: args{
				min: -1,
				max: 1,
			},
		},
		{
			name: "s4",
			args: args{
				min: -2,
				max: -1,
			},
		},
		{
			name: "s5",
			args: args{
				min: -1,
				max: -2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomRange(tt.args.min, tt.args.max)

			if tt.args.min > tt.args.max {
				tt.args.min, tt.args.max = tt.args.max, tt.args.min
			}

			if got > tt.args.max {
				t.Errorf("RandomRange() = %v > %v", got, tt.args.max)
			}
			if got < tt.args.min {
				t.Errorf("RandomRange() = %v < %v", got, tt.args.min)
			}
		})
	}
}
