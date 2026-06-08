package collate

import (
	"cmp"
	"testing"
)

// intComparer is a ComparerFunc over int used to exercise the derived Equal and Less methods.
var intComparer = ComparerFunc[int](cmp.Compare[int])

func TestComparerFunc_Equal(t *testing.T) {
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
			if got := intComparer.Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ComparerFunc.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparerFunc_Less(t *testing.T) {
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
			if got := intComparer.Less(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ComparerFunc.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparerFunc_Compare(t *testing.T) {
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
			if got := intComparer.Compare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ComparerFunc.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
