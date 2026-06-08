package collate

import (
	"testing"
)

func TestSliceOrder_Equal(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal", args: args{x: []int{1, 2, 3}, y: []int{1, 2, 3}}, want: true},
		{name: "both empty", args: args{x: []int{}, y: nil}, want: true},
		{name: "different element", args: args{x: []int{1, 2, 3}, y: []int{1, 9, 3}}, want: false},
		{name: "different length", args: args{x: []int{1, 2}, y: []int{1, 2, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (SliceOrder[[]int, int]{}).Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("SliceOrder.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceOrder_Less(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal", args: args{x: []int{1, 2, 3}, y: []int{1, 2, 3}}, want: false},
		{name: "less by element", args: args{x: []int{1, 2, 3}, y: []int{1, 9, 3}}, want: true},
		{name: "less by prefix", args: args{x: []int{1, 2}, y: []int{1, 2, 3}}, want: true},
		{name: "greater", args: args{x: []int{1, 9}, y: []int{1, 2, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (SliceOrder[[]int, int]{}).Less(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("SliceOrder.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceOrder_Compare(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "equal", args: args{x: []int{1, 2, 3}, y: []int{1, 2, 3}}, want: 0},
		{name: "less by element", args: args{x: []int{1, 2, 3}, y: []int{1, 9, 3}}, want: -1},
		{name: "less by prefix", args: args{x: []int{1, 2}, y: []int{1, 2, 3}}, want: -1},
		{name: "greater by element", args: args{x: []int{1, 9}, y: []int{1, 2, 3}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (SliceOrder[[]int, int]{}).Compare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("SliceOrder.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
