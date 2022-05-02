package galx

import (
	"reflect"
	"testing"
)

func TestCircle_BoundingBox(t *testing.T) {
	type fields struct {
		Pos    Vec2d
		Radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Rect
	}{
		{
			name: "s1",
			fields: fields{
				Pos: Vec2d{
					X: 0,
					Y: 0,
				},
				Radius: 2,
			},
			want: Rect{
				TL: Vec2d{
					X: -2,
					Y: -2,
				},
				BR: Vec2d{
					X: 2,
					Y: 2,
				},
			},
		},
		{
			name: "s1 neg",
			fields: fields{
				Pos: Vec2d{
					X: 0,
					Y: 0,
				},
				Radius: -2,
			},
			want: Rect{
				TL: Vec2d{
					X: -2,
					Y: -2,
				},
				BR: Vec2d{
					X: 2,
					Y: 2,
				},
			},
		},
		{
			name: "s2",
			fields: fields{
				Pos: Vec2d{
					X: -2,
					Y: -3,
				},
				Radius: 1.5,
			},
			want: Rect{
				TL: Vec2d{
					X: -3.5,
					Y: -4.5,
				},
				BR: Vec2d{
					X: -0.5,
					Y: -1.5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				Pos:    tt.fields.Pos,
				Radius: tt.fields.Radius,
			}
			if got := c.BoundingBox(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoundingBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_DistanceTo(t *testing.T) {
	type fields struct {
		Pos    Vec2d
		Radius float64
	}
	type args struct {
		to Circle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "s1",
			fields: fields{
				Pos: Vec2d{
					X: 0,
					Y: 0,
				},
				Radius: 1,
			},
			args: args{
				Circle{
					Pos: Vec2d{
						X: 3,
						Y: 0,
					},
					Radius: 1,
				},
			},
			want: 3,
		},
		{
			name: "s2",
			fields: fields{
				Pos: Vec2d{
					X: 0,
					Y: 0,
				},
				Radius: 1,
			},
			args: args{
				Circle{
					Pos: Vec2d{
						X: -3,
						Y: 0,
					},
					Radius: 1,
				},
			},
			want: 3,
		},
		{
			name: "s3",
			fields: fields{
				Pos: Vec2d{
					X: 2,
					Y: 2,
				},
				Radius: 4,
			},
			args: args{
				Circle{
					Pos: Vec2d{
						X: -5,
						Y: 2,
					},
					Radius: 3,
				},
			},
			want: 7,
		},
		{
			name: "s4",
			fields: fields{
				Pos: Vec2d{
					X: -2,
					Y: -2,
				},
				Radius: -2,
			},
			args: args{
				Circle{
					Pos: Vec2d{
						X: 2,
						Y: 2,
					},
					Radius: 2,
				},
			},
			want: 5.656854249492381,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Pos:    tt.fields.Pos,
				Radius: tt.fields.Radius,
			}
			if got := c.DistanceTo(tt.args.to); got != tt.want {
				t.Errorf("DistanceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
