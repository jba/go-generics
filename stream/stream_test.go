package stream

import (
	"math"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSlice(t *testing.T) {
	want := []int{2, 3, 5, 7, 11}
	s := New(want...)
	got := s.Slice()
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRange(t *testing.T) {
	for _, test := range []struct {
		low, high, step int
		want            []int
	}{
		{1, 3, 1, []int{1, 2}},
		{1, 1, 1, nil},
		{0, 5, 2, []int{0, 2, 4}},
		{5, 0, -2, []int{5, 3, 1}},
		{5, 1, 1, nil},
		{1, 5, -1, nil},
	} {
		got := Range(test.low, test.high, test.step).Slice()
		if !slices.Equal(got, test.want) {
			t.Errorf("Range(%d, %d, %d) = %v, want %v", test.low, test.high, test.step, got, test.want)
		}
	}
}

func TestKeep(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	got := s.Keep(func(x int) bool { return x%2 == 0 }).Slice()
	want := []int{2, 4}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	s := New(1, 4, 9, 16)
	got := Map(s, func(x int) float64 { return math.Sqrt(float64(x)) }).Slice()
	want := []float64{1, 2, 3, 4}
	if !slices.Equal(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestMapConcat(t *testing.T) {
	s := Range(0, 5, 1)
	got := MapConcat(s, func(x int) []int {
		if x%2 == 0 {
			return []int{x, -x}
		}
		return nil
	}).Slice()
	want := []int{0, 0, 2, -2, 4, -4}
	if !slices.Equal(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
