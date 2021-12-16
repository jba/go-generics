// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

import (
	"sort"
	"testing"

	"github.com/jba/go-generics/tuple"
	"golang.org/x/exp/slices"
)

type S struct{ a, b, c string }

func TestMetrics(t *testing.T) {
	m1 := Metric[string]{}
	if got := m1.Count("a"); got != 0 {
		t.Errorf("Count(%q) = %d, want 0", "a", got)
	}
	m1.Add("a")
	m1.Add("a")
	if got := m1.Count("a"); got != 2 {
		t.Errorf("Count(%q) = %d, want 2", "a", got)
	}
	if got, want := m1.Metrics(), []string{"a"}; !slices.Equal(got, want) {
		t.Errorf("Metrics = %v, want %v", got, want)
	}

	m2 := Metric[tuple.T2[int, float64]]{}
	m2.Add(tuple.New2(1, 1.0))
	m2.Add(tuple.New2(2, 2.0))
	m2.Add(tuple.New2(3, 3.0))
	m2.Add(tuple.New2(3, 3.0))
	k := m2.Metrics()

	sort.Slice(k, func(i, j int) bool {
		if k[i].V0() < k[j].V0() {
			return true
		}
		if k[i].V0() > k[j].V0() {
			return false
		}
		return k[i].V1() < k[j].V1()
	})

	w := [](tuple.T2[int, float64]){tuple.New2(1, 1.0), tuple.New2(2, 2.0), tuple.New2(3, 3.0)}
	if !slices.Equal(k, w) {
		t.Errorf("m2.Metrics = %v, want %v", k, w)
	}

	m3 := Metric[tuple.T3[string, S, S]]{}
	m3.Add(tuple.New3("a", S{"d", "e", "f"}, S{"g", "h", "i"}))
	m3.Add(tuple.New3("a", S{"d", "e", "f"}, S{"g", "h", "i"}))
	m3.Add(tuple.New3("a", S{"d", "e", "f"}, S{"g", "h", "i"}))
	m3.Add(tuple.New3("b", S{"d", "e", "f"}, S{"g", "h", "i"}))
	if got := m3.Count(tuple.New3("a", S{"d", "e", "f"}, S{"g", "h", "i"})); got != 3 {
		t.Errorf("Count(%v, %v, %v) = %d, want 3", "a", S{"d", "e", "f"}, S{"g", "h", "i"}, got)
	}
}
