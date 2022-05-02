package galx

import (
	"reflect"
	"testing"
)

func TestRect_Edges(t *testing.T) {
	type fields struct {
		Min Vec2d
		Max Vec2d
	}
	tests := []struct {
		name   string
		fields fields
		want   [4]Line
	}{
		{
			name: "simple",
			fields: fields{
				Min: Vec2d{
					X: -2,
					Y: -2,
				},
				Max: Vec2d{
					X: 2,
					Y: 2,
				},
			},
			want: [4]Line{
				{
					A: Vec2d{
						X: -2,
						Y: -2,
					},
					B: Vec2d{
						X: -2,
						Y: 2,
					},
				},
				{
					A: Vec2d{
						X: -2,
						Y: 2,
					},
					B: Vec2d{
						X: 2,
						Y: 2,
					},
				},
				{
					A: Vec2d{
						X: 2,
						Y: 2,
					},
					B: Vec2d{
						X: 2,
						Y: -2,
					},
				},
				{
					A: Vec2d{
						X: 2,
						Y: -2,
					},
					B: Vec2d{
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
				TL: tt.fields.Min,
				BR: tt.fields.Max,
			}
			if got := r.Edges(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Edges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Height(t *testing.T) {
	type fields struct {
		Min Vec2d
		Max Vec2d
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "t1",
			fields: fields{
				Min: Vec2d{
					X: 0,
					Y: -4,
				},
				Max: Vec2d{
					X: 0,
					Y: 4,
				},
			},
			want: 8,
		},
		{
			name: "t2",
			fields: fields{
				Min: Vec2d{
					X: 3,
					Y: -4,
				},
				Max: Vec2d{
					X: -3,
					Y: 4,
				},
			},
			want: 8,
		},
		{
			name: "t3",
			fields: fields{
				Min: Vec2d{
					X: 3,
					Y: 20,
				},
				Max: Vec2d{
					X: -3,
					Y: -2,
				},
			},
			want: 22,
		},
		{
			name: "around 0",
			fields: fields{
				Min: Vec2d{
					X: 0,
					Y: -100,
				},
				Max: Vec2d{
					X: 0,
					Y: 200,
				},
			},
			want: 300,
		},
		{
			name: "less t 0",
			fields: fields{
				Min: Vec2d{
					X: 0,
					Y: -100,
				},
				Max: Vec2d{
					X: 0,
					Y: -150,
				},
			},
			want: 50,
		},
		{
			name: "less t 0 reversed",
			fields: fields{
				Min: Vec2d{
					X: 0,
					Y: -150,
				},
				Max: Vec2d{
					X: 0,
					Y: -100,
				},
			},
			want: 50,
		},
		{
			name: "simple g 0",
			fields: fields{
				Min: Vec2d{
					X: 0,
					Y: 500,
				},
				Max: Vec2d{
					X: 0,
					Y: 250,
				},
			},
			want: 250,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				TL: tt.fields.Min,
				BR: tt.fields.Max,
			}
			if got := r.Height(); got != tt.want {
				t.Errorf("height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Vertices(t *testing.T) {
	type fields struct {
		Min Vec2d
		Max Vec2d
	}
	tests := []struct {
		name   string
		fields fields
		want   [4]Vec2d
	}{
		{
			name: "simple",
			fields: fields{
				Min: Vec2d{
					X: -2,
					Y: -2,
				},
				Max: Vec2d{
					X: 2,
					Y: 2,
				},
			},
			want: [4]Vec2d{
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
				TL: tt.fields.Min,
				BR: tt.fields.Max,
			}
			if got := r.VerticesClockWise(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vertices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Width(t *testing.T) {
	type fields struct {
		Min Vec2d
		Max Vec2d
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "s1",
			fields: fields{
				Min: Vec2d{
					X: -2,
					Y: 0,
				},
				Max: Vec2d{
					X: 2,
					Y: 0,
				},
			},
			want: 4,
		},
		{
			name: "s2",
			fields: fields{
				Min: Vec2d{
					X: -3,
					Y: 0,
				},
				Max: Vec2d{
					X: 1,
					Y: 0,
				},
			},
			want: 4,
		},
		{
			name: "s3",
			fields: fields{
				Min: Vec2d{
					X: -5,
					Y: 0,
				},
				Max: Vec2d{
					X: -3,
					Y: 0,
				},
			},
			want: 2,
		},
		{
			name: "s4",
			fields: fields{
				Min: Vec2d{
					X: -5,
					Y: 16,
				},
				Max: Vec2d{
					X: -3,
					Y: 12,
				},
			},
			want: 2,
		},
		{
			name: "around 0",
			fields: fields{
				Min: Vec2d{
					X: -100,
					Y: 0,
				},
				Max: Vec2d{
					X: 200,
					Y: 0,
				},
			},
			want: 300,
		},
		{
			name: "less t 0",
			fields: fields{
				Min: Vec2d{
					X: -100,
					Y: 0,
				},
				Max: Vec2d{
					X: -150,
					Y: 0,
				},
			},
			want: 50,
		},
		{
			name: "less t 0 reversed",
			fields: fields{
				Min: Vec2d{
					X: -150,
					Y: 0,
				},
				Max: Vec2d{
					X: -100,
					Y: 0,
				},
			},
			want: 50,
		},
		{
			name: "simple g 0",
			fields: fields{
				Min: Vec2d{
					X: 500,
					Y: 0,
				},
				Max: Vec2d{
					X: 250,
					Y: 0,
				},
			},
			want: 250,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				TL: tt.fields.Min,
				BR: tt.fields.Max,
			}
			if got := r.Width(); got != tt.want {
				t.Errorf("width() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Center(t *testing.T) {
	type fields struct {
		Min Vec2d
		Max Vec2d
	}
	tests := []struct {
		name   string
		fields fields
		want   Vec2d
	}{
		{
			name: "zero",
			fields: fields{
				Min: Vec2d{
					X: 0,
					Y: 0,
				},
				Max: Vec2d{
					X: 0,
					Y: 0,
				},
			},
			want: Vec2d{
				X: 0,
				Y: 0,
			},
		},
		{
			name: "x simple",
			fields: fields{
				Min: Vec2d{
					X: 0,
					Y: 0,
				},
				Max: Vec2d{
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
			name: "x simple 2",
			fields: fields{
				Min: Vec2d{
					X: 5,
					Y: 0,
				},
				Max: Vec2d{
					X: 10,
					Y: 0,
				},
			},
			want: Vec2d{
				X: 7.5,
				Y: 0,
			},
		},
		{
			name: "x less zero 1",
			fields: fields{
				Min: Vec2d{
					X: -5,
					Y: 0,
				},
				Max: Vec2d{
					X: 15,
					Y: 0,
				},
			},
			want: Vec2d{
				X: 5,
				Y: 0,
			},
		},
		{
			name: "x reversed",
			fields: fields{
				Min: Vec2d{
					X: 15,
					Y: 0,
				},
				Max: Vec2d{
					X: -5,
					Y: 0,
				},
			},
			want: Vec2d{
				X: 5,
				Y: 0,
			},
		},
		{
			name: "xy",
			fields: fields{
				Min: Vec2d{
					X: -15,
					Y: 5,
				},
				Max: Vec2d{
					X: 5,
					Y: -15,
				},
			},
			want: Vec2d{
				X: -5,
				Y: -5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				TL: tt.fields.Min,
				BR: tt.fields.Max,
			}
			if got := r.Center(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Center() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Scale(t *testing.T) {
	type fields struct {
		Min Vec2d
		Max Vec2d
	}
	type args struct {
		s float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Rect
	}{
		{
			name: "zero",
			fields: fields{
				Min: Vec2d{X: 0, Y: 0},
				Max: Vec2d{X: 0, Y: 0},
			},
			args: args{
				s: 2,
			},
			want: Rect{
				TL: Vec2d{X: 0, Y: 0},
				BR: Vec2d{X: 0, Y: 0},
			},
		},
		{
			name: "one2one",
			fields: fields{
				Min: Vec2d{X: -1, Y: -1},
				Max: Vec2d{X: 1, Y: 1},
			},
			args: args{
				s: 1,
			},
			want: Rect{
				TL: Vec2d{X: -1, Y: -1},
				BR: Vec2d{X: 1, Y: 1},
			},
		},
		{
			name: "one x2.5",
			fields: fields{
				Min: Vec2d{X: -1, Y: -1},
				Max: Vec2d{X: 1, Y: 1},
			},
			args: args{
				s: 2.5,
			},
			want: Rect{
				TL: Vec2d{X: -2.5, Y: -2.5},
				BR: Vec2d{X: 2.5, Y: 2.5},
			},
		},
		{
			name: "reversed one x2.5 (nornalized inside)",
			fields: fields{
				Min: Vec2d{X: 1, Y: 1},
				Max: Vec2d{X: -1, Y: -1},
			},
			args: args{
				s: 2.5,
			},
			want: Rect{
				TL: Vec2d{X: -2.5, Y: -2.5},
				BR: Vec2d{X: 2.5, Y: 2.5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				TL: tt.fields.Min,
				BR: tt.fields.Max,
			}
			if got := r.Scale(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}
