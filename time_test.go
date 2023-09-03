package collate

import (
	"testing"
	"time"
)

func TestTimeEqual(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal",
			args: args{
				t1: time.Date(0, 1, 2, 3, 4, 5, 6, time.UTC),
				t2: time.Date(0, 1, 2, 3, 4, 5, 6, time.UTC),
			},
			want: true,
		},
		{name: "not equal",
			args: args{
				t1: time.Now(),
				t2: time.Now().Add(time.Hour),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeEqual(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("TimeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeLess(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "equal",
			args: args{
				t1: time.Date(0, 1, 2, 3, 4, 5, 6, time.UTC),
				t2: time.Date(0, 1, 2, 3, 4, 5, 6, time.UTC),
			},
			want: false,
		},
		{name: "more",
			args: args{
				t1: time.Now(),
				t2: time.Now().Add(-time.Hour),
			},
			want: false,
		},
		{name: "less",
			args: args{
				t1: time.Now(),
				t2: time.Now().AddDate(0, 1, 2),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeLess(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("TimeLess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeCompare(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0",
			args: args{
				t1: time.Time{},
				t2: time.Time{},
			},
			want: 0,
		},
		{name: "equal",
			args: args{
				t1: time.Date(0, 1, 2, 3, 4, 5, 6, time.UTC),
				t2: time.Date(0, 1, 2, 3, 4, 5, 6, time.UTC),
			},
			want: 0,
		},
		{name: "more",
			args: args{
				t1: time.Now(),
				t2: time.Now().Add(-time.Hour),
			},
			want: 1,
		},
		{name: "less",
			args: args{
				t1: time.Now(),
				t2: time.Now().AddDate(0, 1, 2),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeCompare(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("TimeCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
