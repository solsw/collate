package collate

import (
	"time"
)

// TimeEqualer is an [Equaler] for [time.Time].
var TimeEqualer Equaler[time.Time] = EqualerFunc[time.Time](TimeEqual)

// TimeLesser is a [Lesser] for [time.Time].
var TimeLesser Lesser[time.Time] = LesserFunc[time.Time](TimeLess)

// TimeComparer is a [Comparer] for [time.Time].
var TimeComparer Comparer[time.Time] = ComparerFunc[time.Time](TimeCompare)

// TimeEqual reports whether 't1' and 't2' represent the same time instant.
func TimeEqual(t1, t2 time.Time) bool {
	return t1.Equal(t2)
}

// TimeLess reports whether the time instant 't1' is before 't2'.
func TimeLess(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// TimeCompare compares the specified time instants.
func TimeCompare(t1, t2 time.Time) int {
	return t1.Compare(t2)
}
