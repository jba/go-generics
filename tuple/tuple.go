// Package tuple provides tuples of values.
// A tuple of size N represents N values, each of a different type.
package tuple

import "fmt"

// T2 is a tuple of two elements.
type T2[A, B any] struct {
	V0 A
	V1 B
}

// T3 is a tuple of three elements.
type T3[A, B, C any] struct {
	V0 A
	V1 B
	V2 C
}

// T4 is a tuple of four elements.
type T4[A, B, C, D any] struct {
	V0 A
	V1 B
	V2 C
	V3 D
}

// T5 is a tuple of five elements.
type T5[A, B, C, D, E any] struct {
	V0 A
	V1 B
	V2 C
	V3 D
	V4 E
}

// New2 returns a new T2.
func New2[A, B any](a A, b B) T2[A, B] {
	return T2[A, B]{a, b}
}

// New3 returns a new T3.
func New3[A, B, C any](a A, b B, c C) T3[A, B, C] {
	return T3[A, B, C]{a, b, c}
}

// New4 returns a new T4.
func New4[A, B, C, D any](a A, b B, c C, d D) T4[A, B, C, D] {
	return T4[A, B, C, D]{a, b, c, d}
}

// New5 returns a new T5.
func New5[A, B, C, D, E any](a A, b B, c C, d D, e E) T5[A, B, C, D, E] {
	return T5[A, B, C, D, E]{a, b, c, d, e}
}

// Spread returns the elements of its receiver as separate return values.
func (t T2[A, B]) Spread() (A, B) { return t.V0, t.V1 }

// Spread returns the elements of its receiver as separate return values.
func (t T3[A, B, C]) Spread() (A, B, C) { return t.V0, t.V1, t.V2 }

// Spread returns the elements of its receiver as separate return values.
func (t T4[A, B, C, D]) Spread() (A, B, C, D) { return t.V0, t.V1, t.V2, t.V3 }

// Spread returns the elements of its receiver as separate return values.
func (t T5[A, B, C, D, E]) Spread() (A, B, C, D, E) { return t.V0, t.V1, t.V2, t.V3, t.V4 }

func (t T2[A, B]) String() string    { return fmt.Sprintf("<%v, %v>", t.V0, t.V1) }
func (t T3[A, B, C]) String() string { return fmt.Sprintf("<%v, %v, %v>", t.V0, t.V1, t.V2) }

func (t T4[A, B, C, D]) String() string {
	return fmt.Sprintf("<%v, %v, %v, %v>",
		t.V0, t.V1, t.V2, t.V3)
}

func (t T5[A, B, C, D, E]) String() string {
	return fmt.Sprintf("<%v, %v, %v, %v, %v>",
		t.V0, t.V1, t.V2, t.V3, t.V4)
}
