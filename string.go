package collate

import (
	"strings"
)

var (
	// CaseInsensitiveEqualer is a case insensitive [Equaler] for [string].
	//
	// [string]: https://go.dev/ref/spec#String_types
	CaseInsensitiveEqualer Equaler[string] = EqualerFunc[string](func(x, y string) bool {
		// strings.EqualFold(x, y) not used to comply with CaseInsensitiveLesser and CaseInsensitiveComparer
		return strings.ToLower(x) == strings.ToLower(y)
	})

	// CaseInsensitiveLesser is a case insensitive [Lesser] for [string].
	//
	// [string]: https://go.dev/ref/spec#String_types
	CaseInsensitiveLesser Lesser[string] = LesserFunc[string](func(x, y string) bool {
		return strings.ToLower(x) < strings.ToLower(y)
	})

	// CaseInsensitiveComparer is a case insensitive [Comparer] for [string].
	//
	// [string]: https://go.dev/ref/spec#String_types
	CaseInsensitiveComparer Comparer[string] = ComparerFunc[string](func(x, y string) int {
		lx := strings.ToLower(x)
		ly := strings.ToLower(y)
		if lx < ly {
			return -1
		}
		if lx > ly {
			return +1
		}
		return 0
	})
)
