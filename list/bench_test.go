package list

import (
	clist "container/list"
	"testing"
)

const N = 100_000

func BenchmarkIterateContainerList(b *testing.B) {
	l := clist.New()
	for j := 0; j < N; j++ {
		l.PushBack(l)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for e := l.Front(); e != nil; e = e.Next() {
		}
	}
}
