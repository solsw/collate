package collate

// Comparer defines a function to compare the specified objects of type T.
type Comparer[T any] interface {
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.icomparer-1

	// Compare compares the specified objects and returns negative if the first one is less than the second,
	// zero if the first one is equal to the second and positive if the first one is greater than the second.
	Compare(T, T) int
}

// The ComparerFunc type is an adapter to allow the use of ordinary functions as a Comparer.
// If f is a function with the appropriate signature, ComparerFunc(f) is a Comparer that calls f.
// ComparerFunc also implements the Equaler and Lesser interfaces.
type ComparerFunc[T any] func(T, T) int

// Compare implements the Comparer interface.
func (cmpf ComparerFunc[T]) Compare(x, y T) int {
	return cmpf(x, y)
}

// Equal implements the Equaler interface.
func (cmpf ComparerFunc[T]) Equal(x, y T) bool {
	return cmpf.Compare(x, y) == 0
}

// Less implements the Lesser interface.
func (cmpf ComparerFunc[T]) Less(x, y T) bool {
	return cmpf.Compare(x, y) < 0
}
