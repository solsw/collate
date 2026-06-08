package collate

import (
	"testing"
)

func TestEqualerFunc_Equal(t *testing.T) {
	eq := EqualerFunc[int](func(x, y int) bool { return x == y })
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
		{name: "not equal", args: args{x: 1, y: 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eq.Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("EqualerFunc.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeepEqualer_Equal(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal", args: args{x: []int{1, 2}, y: []int{1, 2}}, want: true},
		{name: "not equal", args: args{x: []int{1, 2}, y: []int{1, 3}}, want: false},
		{name: "nil and empty", args: args{x: nil, y: []int{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (DeepEqualer[[]int]{}).Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("DeepEqualer.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapEqualer_Equal(t *testing.T) {
	type args struct {
		m1 map[int]string
		m2 map[int]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal", args: args{m1: map[int]string{1: "a", 2: "b"}, m2: map[int]string{2: "b", 1: "a"}}, want: true},
		{name: "nil and empty", args: args{m1: nil, m2: map[int]string{}}, want: true},
		{name: "different value", args: args{m1: map[int]string{1: "a"}, m2: map[int]string{1: "z"}}, want: false},
		{name: "different keys", args: args{m1: map[int]string{1: "a"}, m2: map[int]string{2: "a"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (MapEqualer[map[int]string, int, string]{}).Equal(tt.args.m1, tt.args.m2); got != tt.want {
				t.Errorf("MapEqualer.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
