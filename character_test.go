package main

import (
	_ "image/png"
	"testing"
)

func TestRectangleCollidesWith(t *testing.T) {
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
		{
			name: "Intersection bottom right",
			fields: fields{
				upperLeft:  Position{0, 0},
				lowerRight: Position{3, 3},
			},
			args: args{
				r2: Rectangle{
					Position{2, 2},
					Position{4, 4},
				}},
			want: true,
		},
		{
			name: "Intersection bottom left",
			fields: fields{
				upperLeft:  Position{3, 3},
				lowerRight: Position{6, 6},
			},
			args: args{
				r2: Rectangle{
					Position{0, 2},
					Position{7, 7},
				}},
			want: true,
		},
		{
			name: "Intersection top left",
			fields: fields{
				upperLeft:  Position{3, 3},
				lowerRight: Position{6, 6},
			},
			args: args{
				r2: Rectangle{
					Position{0, 0},
					Position{4, 4},
				}},
			want: true,
		},
		{
			name: "Intersection top right",
			fields: fields{
				upperLeft:  Position{3, 3},
				lowerRight: Position{6, 6},
			},
			args: args{
				r2: Rectangle{
					Position{2, 2},
					Position{4, 4},
				}},
			want: true,
		},
		{
			name: "Touching at top",
			fields: fields{
				upperLeft:  Position{3, 3},
				lowerRight: Position{6, 6},
			},
			args: args{
				r2: Rectangle{
					Position{2, 2},
					Position{3, 3},
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
