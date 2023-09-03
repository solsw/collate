package collate

import (
	"cmp"
)

// The Order type implements the [Equaler], [Lesser] and [Comparer] interfaces for [ordered] types.
//
// [ordered]: https://pkg.go.dev/cmp#Ordered
type Order[T cmp.Ordered] struct{}

// Equal implements the [Equaler] interface.
func (Order[T]) Equal(x, y T) bool {
	return x == y
}

// Less implements the [Lesser] interface.
func (Order[T]) Less(x, y T) bool {
	return cmp.Less(x, y)
}

// Compare implements the [Comparer] interface.
func (Order[T]) Compare(x, y T) int {
	return cmp.Compare(x, y)
}
