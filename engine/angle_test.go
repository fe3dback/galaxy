package engine

import (
	"reflect"
	"testing"
)

func TestBasicAngle(t *testing.T) {
	for want := 0.0; want < 360; want += 0.1 {
		got := NewAngle(want).Degrees()

		// corrects for test only
		// 0.1999 -> 0.2
		// but in real life we want more precision
		// lib will round float numbers to pow(5)
		wantCorrected := RoundTo(want)

		if got != wantCorrected {
			t.Errorf("Circle conv, got `%v` want `%v (%v)`", got, wantCorrected, want)
			break
		}
	}
}

func TestAngle_Add(t *testing.T) {
	type args struct {
		t Angle
	}
	tests := []struct {
		name string
		a    Angle
		args args
		want float64
	}{
		{
			name: "add",
			a:    NewAngle(90),
			args: args{
				t: NewAngle(5),
			},
			want: NewAngle(95).Degrees(),
		},
		{
			name: "sub",
			a:    NewAngle(90),
			args: args{
				t: NewAngle(-5),
			},
			want: NewAngle(85).Degrees(),
		},
		{
			name: "under-clockwise",
			a:    NewAngle(360),
			args: args{
				t: NewAngle(5),
			},
			want: NewAngle(5).Degrees(),
		},
		{
			name: "clockwise",
			a:    NewAngle(0),
			args: args{
				t: NewAngle(-5),
			},
			want: NewAngle(355).Degrees(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Add(tt.args.t); got.Degrees() != tt.want {
				t.Errorf("Add() = %v, want %v", got.Degrees(), tt.want)
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
			a:    NewAngle(90),
			want: 90,
		},
		{
			name: "t2",
			a:    NewAngle(0),
			want: 0,
		},
		{
			name: "n1",
			a:    NewAngle(-5),
			want: 355,
		},
		{
			name: "n2",
			a:    NewAngle(-3605),
			want: 355,
		},
		{
			name: "p1",
			a:    NewAngle(3605),
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Degrees(); got != tt.want {
				t.Errorf("Degrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromDegrees(t *testing.T) {
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
			want: 1.5707963267948966,
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
			want: 0.017453292519943295,
		},
		{
			name: "negative 1",
			args: args{
				f: -1,
			},
			want: 6.265732014659643,
		},
		{
			name: "negative 2",
			args: args{
				f: -3601,
			},
			want: 6.265732014659643,
		},
		{
			name: "positive",
			args: args{
				f: 3602,
			},
			want: 0.03490658503988659,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAngle(tt.args.f); got != tt.want {
				t.Errorf("NewAngle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Unit(t *testing.T) {
	tests := []struct {
		name string
		a    Angle
		want Vec
	}{
		{
			name: "right",
			a:    Angle0,
			want: Vec{
				X: 1,
				Y: 0,
			},
		},
		{
			name: "top",
			a:    Angle90,
			want: Vec{
				X: 0,
				Y: -1,
			},
		},
		{
			name: "right",
			a:    Angle180,
			want: Vec{
				X: -1,
				Y: 0,
			},
		},
		{
			name: "bottom",
			a:    Angle270,
			want: Vec{
				X: 0,
				Y: 1,
			},
		},
		{
			name: "45",
			a:    Angle45,
			want: Vec{
				X: 0.7071,
				Y: -0.7071,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Unit(); !reflect.DeepEqual(testNormVec(got), testNormVec(tt.want)) {
				t.Errorf("Unit() = %v, want %v", got, tt.want)
			}
		})
	}
}
