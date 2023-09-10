package collate

import (
	"cmp"
	"strings"
)

// StringOrder type implements the [Equaler], [Lesser] and [Comparer] interfaces
// for [string]s modified according to the mapping function.
//
// [string]: https://go.dev/ref/spec#String_types
type StringOrder struct {
	Mapping func(string) string
}

// Equal implements the [Equaler] interface.
func (so StringOrder) Equal(x, y string) bool {
	return so.Mapping(x) == so.Mapping(y)
}

// Less implements the [Lesser] interface.
func (so StringOrder) Less(x, y string) bool {
	return cmp.Less(so.Mapping(x), so.Mapping(y))
}

// Compare implements the [Comparer] interface.
func (so StringOrder) Compare(x, y string) int {
	return strings.Compare(so.Mapping(x), so.Mapping(y))
}

// CaseInsensitiveOrder is a case-insensitive [StringOrder].
var CaseInsensitiveOrder = StringOrder{strings.ToLower}
