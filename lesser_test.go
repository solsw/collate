package collate

import (
	"testing"
)

// intLesser is a LesserFunc over int used to exercise the derived Equal and Compare methods.
var intLesser = LesserFunc[int](func(x, y int) bool { return x < y })

func TestLesserFunc_Less(t *testing.T) {
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
			if got := intLesser.Less(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("LesserFunc.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLesserFunc_Equal(t *testing.T) {
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
			if got := intLesser.Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("LesserFunc.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLesserFunc_Compare(t *testing.T) {
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
			if got := intLesser.Compare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("LesserFunc.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
