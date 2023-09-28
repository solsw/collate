package collate

import (
	"maps"
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

// MapEqualer type implements the [Equaler] interface for [map]s.
//
// [map]: https://go.dev/ref/spec#Map_types
type MapEqualer[M ~map[K]V, K, V comparable] struct{}

// Equal implements the [Equaler] interface.
func (MapEqualer[M, K, V]) Equal(m1, m2 M) bool {
	return maps.Equal[M, M, K, V](m1, m2)
}

// check that MapEqualer implements the Equaler interface
var _ Equaler[map[int]string] = MapEqualer[map[int]string, int, string]{}
