package collate

import (
	"testing"
)

func TestCaseInsensitiveEqual(t *testing.T) {
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
			if got := CaseInsensitiveEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CaseInsensitiveEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaseInsensitiveLess(t *testing.T) {
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
			if got := CaseInsensitiveLess(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CaseInsensitiveLess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaseInsensitiveCompare(t *testing.T) {
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
			if got := CaseInsensitiveCompare(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CaseInsensitiveCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
