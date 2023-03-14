package collate

var (
	// BoolEqualer is an [Equaler] for [bool].
	//
	// [bool]: https://go.dev/ref/spec#Boolean_types
	BoolEqualer Equaler[bool] = EqualerFunc[bool](func(x, y bool) bool { return x == y })

	boolLesserFunc = LesserFunc[bool](func(x, y bool) bool { return !x && y })

	// BoolLesser is a [Lesser] for [bool].
	//
	// [bool]: https://go.dev/ref/spec#Boolean_types
	BoolLesser Lesser[bool] = boolLesserFunc

	// BoolComparer is a [Comparer] for [bool].
	//
	// [bool]: https://go.dev/ref/spec#Boolean_types
	BoolComparer Comparer[bool] = boolLesserFunc
)
