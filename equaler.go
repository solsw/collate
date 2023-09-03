package collate

import (
	"reflect"
)

// Equaler defines a function to compare the objects of type T for equality.
type Equaler[T any] interface {

	// Equal reports whether the specified objects are equal.
	Equal(T, T) bool
}

// The EqualerFunc type is an adapter to allow the use of ordinary functions as an [Equaler].
// If f is a function with the appropriate signature, EqualerFunc(f) is an [Equaler] that calls f.
type EqualerFunc[T any] func(T, T) bool

// Equal implements the [Equaler] interface.
func (eqf EqualerFunc[T]) Equal(x, y T) bool {
	return eqf(x, y)
}

// DeepEqualer is an [Equaler] implementation that is a generic wrapper around [reflect.DeepEqual].
type DeepEqualer[T any] struct{}

// Equal implements the [Equaler] interface.
func (DeepEqualer[T]) Equal(x, y T) bool {
	return reflect.DeepEqual(x, y)
}
