package collate

import (
	"cmp"
	"slices"
)

// SliceOrder type implements the [Equaler], [Lesser] and [Comparer] interfaces
// for [slice]s of [ordered] elements.
// Floating-point elements follow the same NaN ordering as [Order].
//
// [ordered]: https://pkg.go.dev/cmp#Ordered
// [slice]: https://go.dev/ref/spec#Slice_types
type SliceOrder[S ~[]E, E cmp.Ordered] struct{}

// Equal implements the [Equaler] interface.
func (so SliceOrder[S, E]) Equal(x, y S) bool {
	return len(x) == len(y) && so.Compare(x, y) == 0
}

// Less implements the [Lesser] interface.
func (SliceOrder[S, E]) Less(x, y S) bool {
	return slices.Compare(x, y) < 0
}

// Compare implements the [Comparer] interface.
func (SliceOrder[S, E]) Compare(x, y S) int {
	return slices.Compare(x, y)
}

// check that SliceOrder implements the Equaler, Lesser and Comparer interfaces
var (
	_ Equaler[[]int]  = SliceOrder[[]int, int]{}
	_ Lesser[[]int]   = SliceOrder[[]int, int]{}
	_ Comparer[[]int] = SliceOrder[[]int, int]{}
)
