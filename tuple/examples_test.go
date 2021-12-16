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
	return joinGen(ga, gb, tuple.Join2[A, B])
}

func CombineGens3[A, B, C any](ga Gen[A], gb Gen[B], gc Gen[C]) Gen[tuple.T3[A, B, C]] {
	return joinGen(CombineGens2(ga, gb), gc, tuple.Join3[A, B, C])
}

func CombineGens4[A, B, C, D any](ga Gen[A], gb Gen[B], gc Gen[C], gd Gen[D]) Gen[tuple.T4[A, B, C, D]] {
	return joinGen(CombineGens3(ga, gb, gc), gd, tuple.Join4[A, B, C, D])
}

func joinGen[T, V, R any](gt Gen[T], gv Gen[V], join func(T, V) R) Gen[R] {
	return func() R { return join(gt(), gv()) }
}

////////////////////////////////////////////////////////////////
// Alternative to CombineGensN: define CheckN in terms of Check.

func Check2[A, B any](max int, pred func(A, B) bool, genA Gen[A], genB Gen[B]) (A, B, bool) {
	p := func(t tuple.T2[A, B]) bool { return pred(t.Spread()) }
	g := func() tuple.T2[A, B] { return tuple.New2(genA(), genB()) }
	r, ok := Check(max, p, g)
	return r.V0(), r.V1(), ok
}
