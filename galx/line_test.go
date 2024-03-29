package galx

import (
	"reflect"
	"testing"
)

func TestLine_Center(t *testing.T) {
	type fields struct {
		A Vec2d
		B Vec2d
	}
	tests := []struct {
		name   string
		fields fields
		want   Vec2d
	}{
		{
			name: "x",
			fields: fields{
				A: Vec2d{
					X: 0,
					Y: 0,
				},
				B: Vec2d{
					X: 10,
					Y: 0,
				},
			},
			want: Vec2d{
				X: 5,
				Y: 0,
			},
		},
		{
			name: "x 2",
			fields: fields{
				A: Vec2d{
					X: 5,
					Y: 0,
				},
				B: Vec2d{
					X: 15,
					Y: 0,
				},
			},
			want: Vec2d{
				X: 10,
				Y: 0,
			},
		},
		{
			name: "x 3",
			fields: fields{
				A: Vec2d{
					X: -5,
					Y: 0,
				},
				B: Vec2d{
					X: 5,
					Y: 0,
				},
			},
			want: Vec2d{
				X: 0,
				Y: 0,
			},
		},
		{
			name: "y",
			fields: fields{
				A: Vec2d{
					X: -1,
					Y: -1,
				},
				B: Vec2d{
					X: 1,
					Y: 1,
				},
			},
			want: Vec2d{
				X: 0,
				Y: 0,
			},
		},
		{
			name: "y 2",
			fields: fields{
				A: Vec2d{
					X: -2,
					Y: 2,
				},
				B: Vec2d{
					X: 2,
					Y: -2,
				},
			},
			want: Vec2d{
				X: 0,
				Y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			line := Line{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := line.Center(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Center() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_Length(t *testing.T) {
	type fields struct {
		A Vec2d
		B Vec2d
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "simple",
			fields: fields{
				A: Vec2d{
					X: 0,
					Y: 0,
				},
				B: Vec2d{
					X: 1,
					Y: 0,
				},
			},
			want: 1,
		},
		{
			name: "x 2",
			fields: fields{
				A: Vec2d{
					X: -2,
					Y: 0,
				},
				B: Vec2d{
					X: 2,
					Y: 0,
				},
			},
			want: 4,
		},
		{
			name: "x 3",
			fields: fields{
				A: Vec2d{
					X: -1,
					Y: -1,
				},
				B: Vec2d{
					X: 1,
					Y: 1,
				},
			},
			want: 2.8284271247461903,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			line := Line{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := line.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}
