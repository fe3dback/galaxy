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
			if got := roundTo(tt.args.n); got != tt.want {
				t.Errorf("roundTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
