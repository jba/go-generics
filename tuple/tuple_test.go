package tuple

import (
	"fmt"
	"testing"
)

func TestVs(t *testing.T) {
	t3 := New3("x", 8, true)
	if got, want := t3.V0, "x"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := t3.V1, 8; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := t3.V2, true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSpread(t *testing.T) {
	t3 := New3("x", 8, true)
	g1, g2, g3 := t3.Spread()
	w1, w2, w3 := "x", 8, true
	if g1 != w1 || g2 != w2 || g3 != w3 {
		t.Errorf("got (%v, %v, %v), want (%v, %v, %v)", g1, g2, g3, w1, w2, w3)
	}
}

func TestString(t *testing.T) {
	tu := New6(1, "x", false, 8.0, -3, 2+3i)
	got := fmt.Sprintf("%v", tu)
	want := "<1, x, false, 8, -3, (2+3i)>"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
