// Package tuple provides tuples of values.
// A tuple of size N represents N values, each of a different type.
package tuple2

import "fmt"

// T2 is a tuple of two elements.
type T2[A, B any] struct {
	T A
	V B
}

// T3 is a tuple of three elements.
type T3[A, B, C any] struct {
	T T2[A, B]
	V C
}

// T4 is a tuple of four elements.
type T4[A, B, C, D any] struct {
	T T3[A, B, C]
	V D
}

// New2 returns a new T2.
func New2[A, B any](a A, b B) T2[A, B] {
	return Join2(a, b)
}

// New3 returns a new T3.
func New3[A, B, C any](a A, b B, c C) T3[A, B, C] {
	return Join3(New2(a, b), c)
}

// New4 returns a new T4.
func New4[A, B, C, D any](a A, b B, c C, d D) T4[A, B, C, D] {
	return Join4(New3(a, b, c), d)
}

// Join2 returns a T2 consisting of the elements t and v.
// Join2 is identical to New2.
func Join2[A, B any](t A, v B) T2[A, B] { return T2[A, B]{t, v} }

// Join3 returns a T3 consisting of the T2 t and the value v.
func Join3[A, B, C any](t T2[A, B], v C) T3[A, B, C] { return T3[A, B, C]{t, v} }

// Join4 returns a T4 consisting of the T3 t and the value v.
func Join4[A, B, C, D any](t T3[A, B, C], v D) T4[A, B, C, D] { return T4[A, B, C, D]{t, v} }

// V0 returns the first element of its receiver tuple.
func (t T2[A, B]) V0() A { return t.T }

// V1 returns the second element of its receiver tuple.
func (t T2[A, B]) V1() B { return t.V }

func (t T3[A, B, C]) V0() A { return t.T.T }
func (t T3[A, B, C]) V1() B { return t.T.V }
func (t T3[A, B, C]) V2() C { return t.V }

func (t T4[A, B, C, D]) V0() A { return t.T.T.T }
func (t T4[A, B, C, D]) V1() B { return t.T.T.V }
func (t T4[A, B, C, D]) V2() C { return t.T.V }
func (t T4[A, B, C, D]) V3() D { return t.V }

// Spread returns the elements of its receiver as separate return values.
func (t T2[A, B]) Spread() (A, B) { return t.V0(), t.V1() }

// Spread returns the elements of its receiver as separate return values.
func (t T3[A, B, C]) Spread() (A, B, C) { return t.V0(), t.V1(), t.V2() }

// Spread returns the elements of its receiver as separate return values.
func (t T4[A, B, C, D]) Spread() (A, B, C, D) { return t.V0(), t.V1(), t.V2(), t.V3() }

func (t T2[A, B]) String() string    { return fmt.Sprintf("<%v, %v>", t.V0(), t.V1()) }
func (t T3[A, B, C]) String() string { return fmt.Sprintf("<%v, %v, %v>", t.V0(), t.V1(), t.V2()) }
func (t T4[A, B, C, D]) String() string {
	return fmt.Sprintf("<%v, %v, %v, %v>",
		t.V0(), t.V1(), t.V2(), t.V3())
}

// Nth returns the nth element ,0-based, of its receiver.
func (t T2[A, B]) Nth(i int) interface{} {
	switch i {
	case 0:
		return t.T
	case 1:
		return t.V
	default:
		panic("bad index")
	}
}

// Nth returns the nth element ,0-based, of its receiver.
func (t T3[A, B, C]) Nth(i int) interface{} {
	if i == 2 {
		return t.V
	}
	return t.T.Nth(i)
}

// Nth returns the nth element ,0-based, of its receiver.
func (t T4[A, B, C, D]) Nth(i int) interface{} {
	if i == 3 {
		return t.V
	}
	return t.T.Nth(i)
}
