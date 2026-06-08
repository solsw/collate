package collate

import (
	"testing"
)

func TestOrder_Equal(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal", args: args{x: 2, y: 2}, want: true},
		{name: "less", args: args{x: 1, y: 2}, want: false},
		{name: "greater", args: args{x: 3, y: 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (Order[int]{}).Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Order.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_Less(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal", args: args{x: 2, y: 2}, want: false},
		{name: "less", args: args{x: 1, y: 2}, want: true},
		{name: "greater", args: args{x: 3, y: 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (Order[int]{}).Less(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Order.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_Compare(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "equal", args: args{x: 2, y: 2}, want: 0},
		{name: "less", args: args{x: 1, y: 2}, want: -1},
		{name: "greater", args: args{x: 3, y: 2}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (Order[int]{}).Compare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Order.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
