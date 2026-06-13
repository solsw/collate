# collate
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/collate.svg)](https://pkg.go.dev/github.com/solsw/collate)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/collate)

Package `collate` contains comparison-related entities: a small set of generic
interfaces for comparing values, function adapters that turn ordinary functions
into those interfaces, and ready-made implementations for ordered types, slices,
strings and `time.Time`.

## Installation

```sh
go get github.com/solsw/collate
```

```go
import "github.com/solsw/collate"
```

## Interfaces

The package is built around three interfaces, each generic over the type `T`
being compared.

| Interface | Method | Reports |
| --- | --- | --- |
| [`Equaler[T]`](https://github.com/solsw/collate/blob/main/equaler.go) | `Equal(T, T) bool` | whether two values are equal |
| [`Lesser[T]`](https://github.com/solsw/collate/blob/main/lesser.go) | `Less(T, T) bool` | whether the first value is less than the second |
| [`Comparer[T]`](https://github.com/solsw/collate/blob/main/comparer.go) | `Compare(T, T) int` | negative / zero / positive ordering |

`Compare` returns a negative number if the first value is less than the second,
zero if they are equal, and a positive number if the first is greater — the same
convention used by the standard library [`cmp`](https://pkg.go.dev/cmp) package.

## Function adapters

Each interface has a matching `…Func` adapter so a plain function can satisfy the
interface. Notably, the adapters cross-implement the other interfaces where the
information is available:

| Adapter | Underlying func | Implements |
| --- | --- | --- |
| [`EqualerFunc[T]`](https://github.com/solsw/collate/blob/main/equaler.go) | `func(T, T) bool` | `Equaler` |
| [`LesserFunc[T]`](https://github.com/solsw/collate/blob/main/lesser.go) | `func(T, T) bool` | `Equaler`, `Lesser`, `Comparer` |
| [`ComparerFunc[T]`](https://github.com/solsw/collate/blob/main/comparer.go) | `func(T, T) int` | `Equaler`, `Lesser`, `Comparer` |

Because `ComparerFunc` carries full ordering information, it can derive both
`Equal` and `Less`. `LesserFunc` derives `Compare` by calling the function twice
(`x < y` and `y < x`) and `Equal` from that.

## Ready-made implementations

| Type / value | Works on | Notes |
| --- | --- | --- |
| [`Order[T cmp.Ordered]`](https://github.com/solsw/collate/blob/main/order.go) | ordered types | wraps `cmp.Less` / `cmp.Compare` |
| [`SliceOrder[S ~[]E, E cmp.Ordered]`](https://github.com/solsw/collate/blob/main/sliceorder.go) | slices of ordered elements | wraps `slices.Equal` / `slices.Compare` |
| [`StringOrder`](https://github.com/solsw/collate/blob/main/stringorder.go) | strings | optional `Mapping func(string) string` applied before comparison |
| [`CaseInsensitiveOrder`](https://github.com/solsw/collate/blob/main/stringorder.go) | strings | a `StringOrder` using `strings.ToLower` |
| [`DeepEqualer[T]`](https://github.com/solsw/collate/blob/main/equaler.go) | any type | wraps `reflect.DeepEqual` (`Equaler` only) |
| [`MapEqualer[M ~map[K]V, K, V comparable]`](https://github.com/solsw/collate/blob/main/equaler.go) | maps | wraps `maps.Equal` (`Equaler` only) |
| [`TimeEqualer` / `TimeLesser` / `TimeComparer`](https://github.com/solsw/collate/blob/main/time.go) | `time.Time` | use `Time.Equal` / `Before` / `Compare` |

## Usage

### Ordered values

```go
var o collate.Order[int]
o.Less(1, 2)    // true
o.Equal(2, 2)   // true
o.Compare(3, 2) // +1
```

### Case-insensitive strings

```go
collate.CaseInsensitiveOrder.Equal("Go", "go") // true

// Custom mapping (e.g. trim before comparing).
so := collate.StringOrder{Mapping: strings.TrimSpace}
so.Equal("hello ", "hello") // true
```

### Slices

```go
var so collate.SliceOrder[[]int, int]
so.Equal([]int{1, 2}, []int{1, 2})   // true
so.Compare([]int{1, 2}, []int{1, 3}) // -1
```

### Adapting a function

```go
// A Comparer from an ordinary function...
byLen := collate.ComparerFunc[string](func(a, b string) int {
	return cmp.Compare(len(a), len(b))
})
byLen.Less("ab", "abc") // true
byLen.Equal("ab", "cd") // true (same length)
```

### Time

```go
collate.TimeLesser.Less(t1, t2)   // t1 before t2
collate.TimeComparer.Compare(t1, t2)
```

## Documentation

Full API reference is available on
[pkg.go.dev](https://pkg.go.dev/github.com/solsw/collate).
