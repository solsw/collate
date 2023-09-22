package collate

import (
	"cmp"
	"slices"
)

// SliceOrder type implements the [Equaler], [Lesser] and [Comparer] interfaces
// for [slice]s of [ordered] elements.
//
// [ordered]: https://pkg.go.dev/cmp#Ordered
// [slice]: https://go.dev/ref/spec#Slice_types
type SliceOrder[S ~[]E, E cmp.Ordered] struct{}

// Equal implements the [Equaler] interface.
func (so SliceOrder[S, E]) Equal(x, y S) bool {
	return slices.Equal[S, E](x, y)
}

// Less implements the [Lesser] interface.
func (so SliceOrder[S, E]) Less(x, y S) bool {
	return slices.Compare[S, E](x, y) < 0
}

// Compare implements the [Comparer] interface.
func (so SliceOrder[S, E]) Compare(x, y S) int {
	return slices.Compare[S, E](x, y)
}
