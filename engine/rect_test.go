package engine

import (
	"reflect"
	"testing"
)

func TestRect_Edges(t *testing.T) {
	type fields struct {
		Min Vec
		Max Vec
	}
	tests := []struct {
		name   string
		fields fields
		want   [4]Line
	}{
		{
			name: "simple",
			fields: fields{
				Min: Vec{
					X: -2,
					Y: -2,
				},
				Max: Vec{
					X: 2,
					Y: 2,
				},
			},
			want: [4]Line{
				{
					A: Vec{
						X: -2,
						Y: -2,
					},
					B: Vec{
						X: -2,
						Y: 2,
					},
				},
				{
					A: Vec{
						X: -2,
						Y: 2,
					},
					B: Vec{
						X: 2,
						Y: 2,
					},
				},
				{
					A: Vec{
						X: 2,
						Y: 2,
					},
					B: Vec{
						X: 2,
						Y: -2,
					},
				},
				{
					A: Vec{
						X: 2,
						Y: -2,
					},
					B: Vec{
						X: -2,
						Y: -2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := r.Edges(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Edges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Height(t *testing.T) {
	type fields struct {
		Min Vec
		Max Vec
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "t1",
			fields: fields{
				Min: Vec{
					X: 0,
					Y: -4,
				},
				Max: Vec{
					X: 0,
					Y: 4,
				},
			},
			want: 8,
		},
		{
			name: "t2",
			fields: fields{
				Min: Vec{
					X: 3,
					Y: -4,
				},
				Max: Vec{
					X: -3,
					Y: 4,
				},
			},
			want: 8,
		},
		{
			name: "t3",
			fields: fields{
				Min: Vec{
					X: 3,
					Y: 20,
				},
				Max: Vec{
					X: -3,
					Y: -2,
				},
			},
			want: -22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := r.Height(); got != tt.want {
				t.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Normalize(t *testing.T) {
	type fields struct {
		Min Vec
		Max Vec
	}
	tests := []struct {
		name   string
		fields fields
		want   Rect
	}{
		{
			name: "simple",
			fields: fields{
				Min: Vec{
					X: 3,
					Y: 2,
				},
				Max: Vec{
					X: -3,
					Y: -2,
				},
			},
			want: Rect{
				Min: Vec{
					X: -3,
					Y: -2,
				},
				Max: Vec{
					X: 3,
					Y: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := r.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Vertices(t *testing.T) {
	type fields struct {
		Min Vec
		Max Vec
	}
	tests := []struct {
		name   string
		fields fields
		want   [4]Vec
	}{
		{
			name: "simple",
			fields: fields{
				Min: Vec{
					X: -2,
					Y: -2,
				},
				Max: Vec{
					X: 2,
					Y: 2,
				},
			},
			want: [4]Vec{
				{
					X: -2,
					Y: -2,
				},
				{
					X: -2,
					Y: 2,
				},
				{
					X: 2,
					Y: 2,
				},
				{
					X: 2,
					Y: -2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := r.Vertices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vertices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Width(t *testing.T) {
	type fields struct {
		Min Vec
		Max Vec
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "s1",
			fields: fields{
				Min: Vec{
					X: -2,
					Y: 0,
				},
				Max: Vec{
					X: 2,
					Y: 0,
				},
			},
			want: 4,
		},
		{
			name: "s2",
			fields: fields{
				Min: Vec{
					X: -3,
					Y: 0,
				},
				Max: Vec{
					X: 1,
					Y: 0,
				},
			},
			want: 4,
		},
		{
			name: "s3",
			fields: fields{
				Min: Vec{
					X: -5,
					Y: 0,
				},
				Max: Vec{
					X: -3,
					Y: 0,
				},
			},
			want: 2,
		},
		{
			name: "s4",
			fields: fields{
				Min: Vec{
					X: -5,
					Y: 16,
				},
				Max: Vec{
					X: -3,
					Y: 12,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := r.Width(); got != tt.want {
				t.Errorf("Width() = %v, want %v", got, tt.want)
			}
		})
	}
}
