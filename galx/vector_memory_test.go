package galx

import (
	"math"
	"reflect"
	"testing"
)

func TestVec2_Data(t *testing.T) {
	type fields struct {
		X float32
		Y float32
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "v2 zero",
			fields: fields{
				X: 0,
				Y: 0,
			},
			want: []byte{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "v2 pi",
			fields: fields{
				X: math.Pi / 2,
				Y: math.Pi,
			},
			want: []byte{219, 15, 201, 63, 219, 15, 73, 64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data() = %v, want %v", got, tt.want)
			}
		})
	}
}
