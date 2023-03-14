package collate

import (
	"reflect"
)

// DeepEqualer is an [Equaler] implementation that is a generic wrapper around [reflect.DeepEqual].
type DeepEqualer[T any] struct{}

// Equal implements the [Equaler] interface.
func (DeepEqualer[T]) Equal(x, y T) bool {
	return reflect.DeepEqual(x, y)
}
