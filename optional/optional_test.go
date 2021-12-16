package optional

import "testing"

func Test(t *testing.T) {
	x := New(3)
	y := x.Must() + 1
	if y != 4 {
		t.Fatal("y != 4")
	}
	x.Set(7)
	if x.Must() != 7 {
		t.Fatal("x.Must() != 7")
	}

	// ValueOr
	var z Opt[int]
	got := z.ValueOr(1)
	if want := 1; got !=want {
		t.Errorf("z.ValueOr(1) = %d, want %d", got, want)
	}
	z.Set(2)
	got = z.ValueOr(1)
	if want := 2; got !=want {
		t.Errorf("z.ValueOr(1) = %d, want %d", got, want)
	}
}
