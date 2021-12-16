package tuple

import (
	"testing"
)

func TestVs(t *testing.T) {
	t3 := New3("x", 8, true)
	if got, want := t3.V0(), "x"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := t3.V1(), 8; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := t3.V2(), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNth(t *testing.T) {
	t3 := New3("x", 8, true)
	if got, want := t3.Nth(0), "x"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := t3.Nth(1), 8; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := t3.Nth(2), true; got != want {
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
