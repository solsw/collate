package collate

import (
	"golang.org/x/exp/constraints"
)

// The Order type implements the [Equaler], [Lesser] and [Comparer] interfaces for ordered types.
type Order[T constraints.Ordered] struct{}

// Equal implements the [Equaler] interface.
func (Order[T]) Equal(x, y T) bool {
	return x == y
}

// Less implements the [Lesser] interface.
func (Order[T]) Less(x, y T) bool {
	return x < y
}

// Compare implements the [Comparer] interface.
func (Order[T]) Compare(x, y T) int {
	if x < y {
		return -1
	}
	if x > y {
		return +1
	}
	return 0
}
