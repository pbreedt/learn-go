package generics

// Generic interface
type Getter[T any] interface {
	Get() T
}

// For concrete type:
// ------------------
type StringStruct struct {
	Val string
}

func (m StringStruct) Get() string {
	return m.Val
}

// This is ok
func NewStringGetter() Getter[string] {
	return StringStruct{}
}

// This won't work: according to interface, StringStruct MUST implement 'Get() T' even if T is string
// func NewGetter[T any]() Getter[T] {
//     return StringStruct{} // doesn't compile, even if T is string
// }

// For generic type:
// -----------------
type GenericStruct[T any] struct {
	Val T
}

func (m GenericStruct[T]) Get() T {
	return m.Val
}

func NewGetter[T any]() Getter[T] {
	return GenericStruct[T]{} // ok, GenericStruct implements 'Get() T'
}
