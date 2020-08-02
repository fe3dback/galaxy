package engine

import (
	"testing"
)

func TestAngle_Add(t *testing.T) {
	type args struct {
		t Angle
	}
	tests := []struct {
		name string
		a    Angle
		args args
		want Angle
	}{
		{
			name: "add",
			a:    90,
			args: args{
				t: Anglef(5),
			},
			want: 95,
		},
		{
			name: "sub",
			a:    90,
			args: args{
				t: Anglef(-5),
			},
			want: 85,
		},
		{
			name: "under-clockwise",
			a:    360,
			args: args{
				t: Anglef(5),
			},
			want: 5,
		},
		{
			name: "clockwise",
			a:    0,
			args: args{
				t: Anglef(-5),
			},
			want: 355,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Add(tt.args.t); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_ToFloat(t *testing.T) {
	tests := []struct {
		name string
		a    Angle
		want float64
	}{
		{
			name: "t1",
			a:    Anglef(90),
			want: 90,
		},
		{
			name: "t2",
			a:    Anglef(0),
			want: 0,
		},
		{
			name: "n1",
			a:    Anglef(-5),
			want: 355,
		},
		{
			name: "n2",
			a:    Anglef(-3605),
			want: 355,
		},
		{
			name: "p1",
			a:    Anglef(3605),
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.ToFloat(); got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnglef(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want Angle
	}{
		{
			name: "simple",
			args: args{
				f: 90,
			},
			want: 90,
		},
		{
			name: "zero",
			args: args{
				f: 0,
			},
			want: 0,
		},
		{
			name: "mod 1",
			args: args{
				f: 360,
			},
			want: 0,
		},
		{
			name: "mod 2",
			args: args{
				f: 720,
			},
			want: 0,
		},
		{
			name: "mod 3",
			args: args{
				f: 721,
			},
			want: 1,
		},
		{
			name: "negative 1",
			args: args{
				f: -1,
			},
			want: 359,
		},
		{
			name: "negative 2",
			args: args{
				f: -3601,
			},
			want: 359,
		},
		{
			name: "positive",
			args: args{
				f: 3602,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anglef(tt.args.f); got != tt.want {
				t.Errorf("Anglef() = %v, want %v", got, tt.want)
			}
		})
	}
}
