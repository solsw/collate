package collate

import (
	"cmp"
	"strings"
)

// CaseInsensitiveEqualer is a case-insensitive [Equaler] for [string].
//
// [string]: https://go.dev/ref/spec#String_types
var CaseInsensitiveEqualer Equaler[string] = EqualerFunc[string](CaseInsensitiveEqual)

// CaseInsensitiveLesser is a case-insensitive [Lesser] for [string].
//
// [string]: https://go.dev/ref/spec#String_types
var CaseInsensitiveLesser Lesser[string] = LesserFunc[string](CaseInsensitiveLess)

// CaseInsensitiveComparer is a case-insensitive [Comparer] for [string].
//
// [string]: https://go.dev/ref/spec#String_types
var CaseInsensitiveComparer Comparer[string] = ComparerFunc[string](CaseInsensitiveCompare)

// CaseInsensitiveEqual reports whether the string 's1' is case-insensitively equal to 's2'.
func CaseInsensitiveEqual(s1, s2 string) bool {
	// strings.EqualFold(x, y) is not used to comply with CaseInsensitiveLess and CaseInsensitiveCompare
	return strings.ToLower(s1) == strings.ToLower(s2)
}

// CaseInsensitiveLess reports whether the string 's1' is case-insensitively less than 's2'.
func CaseInsensitiveLess(s1, s2 string) bool {
	return cmp.Less(strings.ToLower(s1), strings.ToLower(s2))
}

// CaseInsensitiveCompare compares case-insensitively the specified strings.
// (See [Comparer.Compare] for return values.)
func CaseInsensitiveCompare(s1, s2 string) int {
	return cmp.Compare(strings.ToLower(s1), strings.ToLower(s2))
}
