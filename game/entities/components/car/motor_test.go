package car

import (
	"reflect"
	"testing"

	"github.com/fe3dback/galaxy/game/units"
)

func Test_engineTorque(t *testing.T) {
	type args struct {
		rpm float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "<min",
			args: args{
				rpm: 0,
			},
			want: torqueMin,
		},
		{
			name: "min",
			args: args{
				rpm: rpmMin,
			},
			want: torqueMin,
		},
		{
			name: "peek",
			args: args{
				rpm: rpmPeek,
			},
			want: torqueMax,
		},
		{
			name: "max",
			args: args{
				rpm: rpmMax,
			},
			want: torqueRedLine,
		},
		{
			name: ">max",
			args: args{
				rpm: rpmMax * 3,
			},
			want: torqueRedLine,
		},
		{
			name: "lerp 1",
			args: args{
				rpm: 2000,
			},
			want: 307.6470588235294,
		},
		{
			name: "lerp 2",
			args: args{
				rpm: 3000,
			},
			want: 325.29411764705884,
		},
		{
			name: "lerp 3",
			args: args{
				rpm: 4000,
			},
			want: 342.94117647058823,
		},
		{
			name: "lerp 4",
			args: args{
				rpm: 4500,
			},
			want: 345.625,
		},
		{
			name: "lerp 5",
			args: args{
				rpm: 5500,
			},
			want: 301.875,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := engineTorque(tt.args.rpm); got != tt.want {
				t.Errorf("engineTorque() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_engineRpm(t *testing.T) {
	type args struct {
		gearIndex    gearInd
		wheelsRadius float64
		speed        units.SpeedKmH
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "s1",
			args: args{
				gearIndex:    1,
				wheelsRadius: 0.33,
				speed:        units.SpeedKmH(20),
			},
			want: 1462.4892407026166,
		},
		{
			name: "s2",
			args: args{
				gearIndex:    2,
				wheelsRadius: 0.33,
				speed:        units.SpeedKmH(40),
			},
			want: 1957.3164274065093,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := engineRpm(tt.args.gearIndex, tt.args.wheelsRadius, tt.args.speed); got != tt.want {
				t.Errorf("engineRpm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateLongForce(t *testing.T) {
	type args struct {
		velocity      Vec
		driveForce    Vec
		directionUnit Vec
		isBraking     bool
	}
	tests := []struct {
		name string
		args args
		want Vec
	}{
		{
			name: "s1",
			args: args{
				velocity: Vec{
					X: 10,
					Y: 0,
				},
				driveForce: Vec{
					X: 300,
					Y: 0,
				},
				directionUnit: Vec{
					X: 1,
					Y: 0,
				},
				isBraking: false,
			},
			want: Vec{
				X: 130,
				Y: 0,
			},
		},
		{
			name: "s2",
			args: args{
				velocity: Vec{
					X: 15,
					Y: 0,
				},
				driveForce: Vec{
					X: 450,
					Y: 0,
				},
				directionUnit: Vec{
					X: 1,
					Y: 0,
				},
				isBraking: false,
			},
			want: Vec{
				X: 163,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateLongForce(tt.args.velocity, tt.args.driveForce, tt.args.directionUnit, tt.args.isBraking)

			if got := result.longitudinalForce.Round(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateLongForce() got, want: \n%v\n%v", got, tt.want)
			}
		})
	}
}
