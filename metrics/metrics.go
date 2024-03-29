// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metrics provides tracking arbitrary metrics composed of
// values of comparable types.
package metrics

import "sync"

// Metric tracks metrics of values of some type.
type Metric[T comparable] struct {
	mu sync.Mutex
	m  map[T]int
}

// Add adds another instance of some value.
func (m *Metric[T]) Add(v T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.m == nil {
		m.m = make(map[T]int)
	}
	m.m[v]++
}

// Count returns the number of instances we've seen of v.
func (m *Metric[T]) Count(v T) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.m[v]
}

// Metrics returns all the values we've seen, in an indeterminate order.
func (m *Metric[T]) Metrics() []T {
	// TODO: use maps.Keys when available.
	var keys []T
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}
