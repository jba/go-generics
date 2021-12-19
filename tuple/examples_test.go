// This example demonstrates how a property-testing framework
// could benefit from tuples.
// This design is simplified from github.com/leanovate/gopter.
package tuple_test

import "github.com/jba/go-generics/tuple"

// Check generates up to max values by calling gen, and returns the first
// value for which pred returns false, along with a second return value of true.
// If pred never returns false, the second return value is false.
func Check[T any](max int, pred func(T) bool, gen Gen[T]) (T, bool) {
	for i := 0; i < max; i++ {
		v := gen()
		if !pred(v) {
			return v, true
		}
	}
	var z T
	return z, false
}

// Gen is a generator: each time it is called, it produces a new value of type T.
type Gen[T any] func() T

// CombineGens2 converts two generators into a single generator that returns a tuple.
func CombineGens2[A, B any](ga Gen[A], gb Gen[B]) Gen[tuple.T2[A, B]] {
	return func() tuple.T2[A, B] {
		return tuple.New2(ga(), gb())
	}
}

func CombineGens3[A, B, C any](ga Gen[A], gb Gen[B], gc Gen[C]) Gen[tuple.T3[A, B, C]] {
	return func() tuple.T3[A, B, C] {
		return tuple.New3(ga(), gb(), gc())
	}
}

func CombineGens4[A, B, C, D any](ga Gen[A], gb Gen[B], gc Gen[C], gd Gen[D]) Gen[tuple.T4[A, B, C, D]] {
	return func() tuple.T4[A, B, C, D] {
		return tuple.New4(ga(), gb(), gc(), gd())
	}
}

////////////////////////////////////////////////////////////////
// Alternative to CombineGensN: define CheckN in terms of Check.

func Check2[A, B any](max int, pred func(A, B) bool, genA Gen[A], genB Gen[B]) (A, B, bool) {
	p := func(t tuple.T2[A, B]) bool { return pred(t.Spread()) }
	g := func() tuple.T2[A, B] { return tuple.New2(genA(), genB()) }
	r, ok := Check(max, p, g)
	return r.V0, r.V1, ok
}
