package collate

import (
	"testing"
)

func TestCaseInsensitiveOrder_Unicode(t *testing.T) {
	for _, tt := range []struct {
		x, y string
		want int
	}{
		{"Σ", "ς", 0},
		{"ς", "σ", 0},
		{"ſ", "S", 0},
		{"Σa", "ςb", -1},
		{"İ", "i", 1},
		{"[", "A", -1}, // Preserve lowercase ASCII ordering around punctuation.
	} {
		order := CaseInsensitiveOrder
		if got := order.Compare(tt.x, tt.y); got != tt.want {
			t.Errorf("Compare(%q, %q) = %d, want %d", tt.x, tt.y, got, tt.want)
		}
		if got := order.Equal(tt.x, tt.y); got != (tt.want == 0) {
			t.Errorf("Equal(%q, %q) = %t", tt.x, tt.y, got)
		}
		if got := order.Less(tt.x, tt.y); got != (tt.want < 0) {
			t.Errorf("Less(%q, %q) = %t", tt.x, tt.y, got)
		}
	}
}

func TestStringOrder_Equal(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal", args: args{x: "qwerty", y: "qwerty"}, want: true},
		{name: "case differs", args: args{x: "qwerty", y: "qWERTy"}, want: false},
		{name: "not equal", args: args{x: "qwerty", y: "qwert"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (StringOrder{}).Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("StringOrder.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringOrder_Less(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal", args: args{x: "qwerty", y: "qwerty"}, want: false},
		{name: "less", args: args{x: "qwert", y: "qwerty"}, want: true},
		{name: "greater", args: args{x: "qwerty", y: "qwert"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (StringOrder{}).Less(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("StringOrder.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringOrder_Compare(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "equal", args: args{x: "qwerty", y: "qwerty"}, want: 0},
		{name: "less", args: args{x: "qwert", y: "qwerty"}, want: -1},
		{name: "greater", args: args{x: "qwerty", y: "qwert"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (StringOrder{}).Compare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("StringOrder.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaseInsensitiveOrder_Equal(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				s1: "qwerty",
				s2: "qWERTy",
			},
			want: true,
		},
		{name: "2",
			args: args{
				s1: "qwerty",
				s2: "qwert",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaseInsensitiveOrder.Equal(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CaseInsensitiveOrder.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaseInsensitiveOrder_Less(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				s1: "qwerty",
				s2: "qWERTy",
			},
			want: false,
		},
		{name: "2",
			args: args{
				s1: "qwert",
				s2: "qwerty",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaseInsensitiveOrder.Less(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CaseInsensitiveOrder.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaseInsensitiveOrder_Compare(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0",
			args: args{
				s1: "",
				s2: "",
			},
			want: 0,
		},
		{name: "1",
			args: args{
				s1: "qwerty",
				s2: "qWERTy",
			},
			want: 0,
		},
		{name: "2",
			args: args{
				s1: "Qwert",
				s2: "QWeRTY",
			},
			want: -1,
		},
		{name: "3",
			args: args{
				s1: "qwERTY",
				s2: "QWert",
			},
			want: 1,
		},
		{name: "4",
			args: args{
				s1: "qwerty",
				s2: "asdfgh",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaseInsensitiveOrder.Compare(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CaseInsensitiveOrder.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
