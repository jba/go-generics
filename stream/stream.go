// Package stream implements operations on streams of values.
package stream

import "constraints"

// A Stream[T] is a possibly infinite sequence of values of type T. It is
// implemented by a function that returns either a valid T and true, or the zero
// value for T and false.
type Stream[T any] func() (T, bool)

// New returns a Stream[T] consisting of the given values.
func New[T any](vals ...T) Stream[T] {
	return func() (T, bool) {
		if len(vals) == 0 {
			var z T
			return z, false
		}
		t := vals[0]
		vals = vals[1:]
		return t, true
	}
}

// Range returns a stream over numeric values from low (inclusive) to high (exclusive),
// in intervals of step.
// Range panics if step is zero.
func Range[T constraints.Integer | constraints.Float](low, high, step T) Stream[T] {
	if step == 0 {
		panic("0 step")
	}
	i := low
	return func() (T, bool) {
		if (step > 0 && i >= high) || (step < 0 && i <= high) {
			var z T
			return z, false
		}
		r := i
		i += step
		return r, true
	}
}

// Slice collects all the values of s into a slice.
func (s Stream[T]) Slice() []T {
	var c []T
	for {
		next, ok := s()
		if !ok {
			return c
		}
		c = append(c, next)
	}
}

// Keep returns a stream that contains the values of s for which f returns true.
func (s Stream[T]) Keep(f func(T) bool) Stream[T] {
	return func() (T, bool) {
		for {
			next, ok := s()
			if !ok {
				return next, false
			}
			if f(next) {
				return next, true
			}
		}
	}
}

// Remove returns a stream that contains all the values of s for which
// f returns false.
func (s Stream[T]) Remove(f func(T) bool) Stream[T] {
	return s.Keep(func(x T) bool { return !f(x) })
}

// Apply invokes f for each element of s.
func (s Stream[T]) Apply(f func(T)) {
	for {
		next, ok := s()
		if !ok {
			break
		}
		f(next)
	}
}

// Map returns the stream that results from applying f to each element of s.
func Map[T, U any](s Stream[T], f func(T) U) Stream[U] {
	return func() (U, bool) {
		next, ok := s()
		if !ok {
			var u U
			return u, false
		}
		return f(next), true
	}
}

// MapConcat applies f to each element of s. It returns the stream consisting
// of all the elements returned by f concatenated together.
func MapConcat[T, U any](s Stream[T], f func(T) []U) Stream[U] {
	var us []U
	return func() (U, bool) {
		for {
			if len(us) > 0 {
				u := us[0]
				us = us[1:]
				return u, true
			}
			next, ok := s()
			if !ok {
				var u U
				return u, false
			}
			us = f(next)
		}
	}
}
