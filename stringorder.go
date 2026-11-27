package collate

import (
	"strings"
	"unicode"
)

// StringOrder type implements the [Equaler], [Lesser] and [Comparer] interfaces
// for [string]s modified according to the mapping function, if any.
//
// [string]: https://go.dev/ref/spec#String_types
type StringOrder struct {
	Mapping func(string) string
}

// Equal implements the [Equaler] interface.
func (so StringOrder) Equal(x, y string) bool {
	if so.Mapping != nil {
		return so.Mapping(x) == so.Mapping(y)
	}
	return x == y
}

// Less implements the [Lesser] interface.
func (so StringOrder) Less(x, y string) bool {
	if so.Mapping != nil {
		x = so.Mapping(x)
		y = so.Mapping(y)
	}
	return strings.Compare(x, y) < 0
}

// Compare implements the [Comparer] interface.
func (so StringOrder) Compare(x, y string) int {
	if so.Mapping != nil {
		x = so.Mapping(x)
		y = so.Mapping(y)
	}
	return strings.Compare(x, y)
}

// CaseInsensitiveOrder uses Unicode simple case folding, so its Equal method
// agrees with [strings.EqualFold].
var CaseInsensitiveOrder = StringOrder{Mapping: func(s string) string {
	return strings.Map(foldRune, s)
}}

func foldRune(r rune) rune {
	if r <= unicode.MaxASCII {
		return unicode.ToLower(r)
	}
	// Pick the smallest lowercase member of the fold class when possible.
	// If the class has no lowercase member, use its smallest rune.
	best := r
	for member := unicode.SimpleFold(r); member != r; member = unicode.SimpleFold(member) {
		lower := unicode.IsLower(member)
		if lower == unicode.IsLower(best) {
			best = min(best, member)
		} else if lower {
			best = member
		}
	}
	return best
}

// check that StringOrder implements the Equaler, Lesser and Comparer interfaces
var (
	_ Equaler[string]  = StringOrder{}
	_ Lesser[string]   = StringOrder{}
	_ Comparer[string] = StringOrder{}
)
