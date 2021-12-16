// Package optional provides Opt[T] values, which represent
// either a value of type T or no value.
package optional

// An Opt[T] represents either a value of type T, or no value.
// The zero value of an Opt[T] represents no value.
type Opt[T any] struct {
	Value   T
	Present bool
}

// New returns an Opt[T] with value x.
func New[T any](x T) Opt[T] {
	return Opt[T]{Value: x, Present: true}
}

// Set sets the value of o to x.
func (o *Opt[T]) Set(x T) {
	o.Value = x
	o.Present = true
}

// ValueOr returns o.Value if it is present. Otherwise it returns other.
func (o Opt[T]) ValueOr(other T) T {
	if o.Present {
		return o.Value
	}
	return other
}

// Must returns T if o has a value. Otherwise it panics.
func (o Opt[T]) Must() T {
	if !o.Present {
		panic("missing value")
	}
	return o.Value
}
