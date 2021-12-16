package list

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPushBack(t *testing.T) {
	l := &List[int]{}
	check(t, l, "")
	l.PushBack(1)
	check(t, l, "1")
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	check(t, l, "1 2 3 4")
	l.PushBack(5)
	check(t, l, "1 2 3 4 | 5")
	l.PushBack(6)
	check(t, l, "1 2 3 4 | 5 6")
}

func TestPushFront(t *testing.T) {
	l := &List[int]{}
	for i := 10; i > 0; i-- {
		l.PushFront(i)
	}
	check(t, l, "1 2 3 4 | 5 6 | 7 8 9 10")
}

func TestInsert(t *testing.T) {
	l := newList(10, 20, 30, 40, 50)
	check(t, l, "10 20 30 40 | 50")
	// Insert at start of non-full node.
	l.InsertBefore(45, find(l, 50))
	check(t, l, "10 20 30 40 | 45 50")
	// Insert at end of full node.
	l.InsertAfter(43, find(l, 40))
	check(t, l, "10 20 30 40 | 43 45 50")
	// Insert at end of no-full node.
	l.InsertAfter(60, find(l, 50))
	check(t, l, "10 20 30 40 | 43 45 50 60")
	// Force a split.
	l.InsertAfter(25, find(l, 20))
	check(t, l, "10 20 25 | 30 40 | 43 45 50 60")
	// Force a move.
	l.InsertAfter(27, find(l, 25))
	check(t, l, "10 20 25 27 | 30 40 | 43 45 50 60")
	l.InsertBefore(15, find(l, 20))
	check(t, l, "10 15 20 | 25 27 30 40 | 43 45 50 60")
	// TODO: force a move, case of copying from i.
}
func TestRemove(t *testing.T) {
	l := &List[int]{}
	var els []Element[int]
	for i := 0; i < 10; i++ {
		els = append(els, l.PushBack(i))
	}
	check(t, l, "0 1 2 3 | 4 5 6 7 | 8 9")
	l.Remove(els[2])
	check(t, l, "0 1 3 | 4 5 6 7 | 8 9")
	l.Remove(l.Front())
	check(t, l, "1 3 | 4 5 6 7 | 8 9")
	l.Remove(l.Back())
	check(t, l, "1 3 | 4 5 6 7 | 8")
	l.Remove(l.Back())
	check(t, l, "1 3 | 4 5 6 7")
}

func slice[T any](l *List[T]) []T {
	var els []T
	for n := l.front; n != nil; n = n.next {
		els = append(els, n.vals[:n.len]...)
	}
	return els
}

func check(t *testing.T, l *List[int], want string) {
	t.Helper()
	got := structure(l)
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func str(l *List[int]) string {
	var buf bytes.Buffer
	io.WriteString(&buf, "[")
	first := true
	for n := l.front; n != nil; n = n.next {
		for i := 0; i < n.len; i++ {
			if !first {
				io.WriteString(&buf, ", ")
			}
			first = false
			fmt.Fprintf(&buf, "%d", n.vals[i])
		}
	}
	io.WriteString(&buf, "]")
	return buf.String()
}

func structure(l *List[int]) string {
	var buf bytes.Buffer
	for n := l.front; n != nil; n = n.next {
		for i := 0; i < n.len; i++ {
			fmt.Fprintf(&buf, "%d ", n.vals[i])
		}
		if n.next != nil {
			io.WriteString(&buf, "| ")
		}
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	return buf.String()
}

func newList(vals ...int) *List[int] {
	l := &List[int]{}
	for _, v := range vals {
		l.PushBack(v)
	}
	return l
}

func find(l *List[int], target int) Element[int] {
	for e := l.Front(); !e.Zero(); e = e.Next() {
		if e.Value() == target {
			return e
		}
	}
	panic("not found")
}
