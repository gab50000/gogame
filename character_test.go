package main

import (
	_ "image/png"
	"testing"
)

func TestRectangle_collidesWith(t *testing.T) {
	type fields struct {
		upperLeft  Position
		lowerRight Position
	}
	type args struct {
		r2 Rectangle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Intersecting rectangles",
			fields: fields{
				upperLeft:  Position{0, 0},
				lowerRight: Position{3, 3},
			},
			args: args{r2: Rectangle{
				Position{2, 2},
				Position{3, 3},
			}},
			want: true,
		},
		{
			name: "Non-intersecting rectangles",
			fields: fields{
				upperLeft:  Position{0, 0},
				lowerRight: Position{3, 3},
			},
			args: args{r2: Rectangle{
				Position{4, 4},
				Position{6, 6},
			}},
			want: false,
		},
		{
			name: "Intersecting rectangles along x-axis",
			fields: fields{
				upperLeft:  Position{0, 0},
				lowerRight: Position{3, 3},
			},
			args: args{
				r2: Rectangle{
					Position{0, 4},
					Position{0, 6},
				}},
			want: false,
		},
		{
			name: "Intersecting rectangles along y-axis",
			fields: fields{
				upperLeft:  Position{0, 0},
				lowerRight: Position{3, 3},
			},
			args: args{
				r2: Rectangle{
					Position{10, 1},
					Position{10, 2},
				}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r1 := Rectangle{
				upperLeft:  tt.fields.upperLeft,
				lowerRight: tt.fields.lowerRight,
			}
			if got := r1.collidesWith(tt.args.r2); got != tt.want {
				t.Errorf("Rectangle.collidesWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
