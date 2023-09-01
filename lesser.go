package collate

// Lesser defines a function to compare the specified objects of type T.
type Lesser[T any] interface {

	// Less reports whether the first object is less than the second.
	Less(T, T) bool
}

// The LesserFunc type is an adapter to allow the use of ordinary functions as a [Lesser].
// If f is a function with the appropriate signature, LesserFunc(f) is a [Lesser] that calls f.
// LesserFunc also implements the [Equaler] and [Comparer] interfaces.
type LesserFunc[T any] func(T, T) bool

// Less implements the [Lesser] interface.
func (lsf LesserFunc[T]) Less(x, y T) bool {
	return lsf(x, y)
}

// Equal implements the [Equaler] interface.
func (lsf LesserFunc[T]) Equal(x, y T) bool {
	return lsf.Compare(x, y) == 0
}

// Compare implements the [Comparer] interface.
func (lsf LesserFunc[T]) Compare(x, y T) int {
	if lsf.Less(x, y) {
		return -1
	}
	if lsf.Less(y, x) {
		return +1
	}
	return 0
}
