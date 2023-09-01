package collate

// EqualerTo defines a function to compare the instance with the object of type T for equality.
type EqualerTo[T any] interface {

	// EqualTo reports whether the instance is equal to the object of type T.
	EqualTo(T) bool
}
