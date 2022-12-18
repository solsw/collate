package collate

var (
	// BoolEqualer is an Equaler for bool.
	BoolEqualer Equaler[bool] = EqualerFunc[bool](func(x, y bool) bool { return x == y })

	boolLesserFunc = LesserFunc[bool](func(x, y bool) bool { return !x && y })

	// BoolLesser is a Lesser for bool.
	BoolLesser Lesser[bool] = boolLesserFunc

	// BoolComparer is a Comparer for bool.
	BoolComparer Comparer[bool] = boolLesserFunc
)
