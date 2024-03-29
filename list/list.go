// Package list implements an unrolled doubly-linked list.
package list

type Element[T any] struct {
	n *node[T]
	i int
}

func (e Element[T]) Zero() bool {
	return e.n == nil && e.i == 0
}

// Any modification to the list invalidates Value.
func (e Element[T]) Value() T {
	return e.n.vals[e.i]
}

func (e Element[T]) Next() Element[T] {
	i := e.i + 1
	if i < e.n.len {
		return Element[T]{e.n, i}
	}
	return Element[T]{e.n.next, 0}
}

func (e *Element[T]) Prev() Element[T] {
	i := e.i - 1
	if i >= 0 {
		return Element[T]{e.n, i}
	}
	return Element[T]{e.n.prev, 0}
}

const nodeSize = 4

type node[T any] struct {
	vals       [nodeSize]T
	len        int
	next, prev *node[T]
}

type List[T any] struct {
	front, back *node[T]
	len         int
}

func (l *List[T]) Front() Element[T] {
	return Element[T]{l.front, 0}
}

func (l *List[T]) Back() Element[T] {
	if l.len == 0 {
		return Element[T]{}
	}
	return Element[T]{l.back, l.back.len}
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) InsertAfter(v T, e Element[T]) Element[T] {
	// assert e.i < e.n.len
	return l.insertAt(v, e.n, e.i+1)
}

func (l *List[T]) InsertBefore(v T, e Element[T]) Element[T] {
	return l.insertAt(v, e.n, e.i)
}

func (l *List[T]) insertAt(v T, n *node[T], i int) Element[T] {
	l.len++
	if n.len == nodeSize {
		// This node is full. Add the element to the next node if there's room.
		// Otherwise, split this node into two.
		if n.next == nil {
			l.addNodeAfter(n)
		}
		assert(n.next != nil)
		if n.next.len == nodeSize {
			// Split this node.
			const halfSize = nodeSize / 2
			n2 := l.addNodeAfter(n)
			copy(n2.vals[:], n.vals[halfSize:])
			n2.len = halfSize
			n.len = halfSize
			if i > halfSize {
				n = n2
				i -= halfSize
			}
		} else if i == n.len {
			n = n.next
			i = 0
		} else {
			// Try to move i and above; move less there's not enough room.
			m := n.len - i
			if nodeSize-n.next.len < m {
				m = nodeSize - n.next.len
			}
			copy(n.next.vals[m:], n.next.vals[:n.next.len])
			copy(n.next.vals[:m], n.vals[n.len-m:n.len])
			n.len -= m
			n.next.len += m
			// We still want to insert at position i in n.
		}
	}
	assert(n.len < nodeSize)
	assert(i <= n.len)
	n.insertNonFull(v, i)
	return Element[T]{n, i}
}

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func (n *node[T]) insertNonFull(v T, i int) {
	assert(n.len < nodeSize)
	copy(n.vals[i+1:], n.vals[i:n.len])
	n.vals[i] = v
	n.len++
}

func (l *List[T]) PushFront(v T) Element[T] {
	if l.front == nil {
		l.init()
	}
	return l.insertAt(v, l.front, 0)
}

func (l *List[T]) PushBack(v T) Element[T] {
	if l.back == nil {
		l.init()
	}
	return l.insertAt(v, l.back, l.back.len)
}

func (l *List[T]) init() {
	n := &node[T]{}
	l.front = n
	l.back = n
}

func (l *List[T]) addNodeAfter(n *node[T]) *node[T] {
	n2 := &node[T]{next: n.next, prev: n}
	n.next = n2
	if n2.next == nil {
		l.back = n2
	} else {
		n2.next.prev = n2
	}
	return n2
}

func (l *List[T]) Remove(e Element[T]) {
	//  TODO: merge half-full nodes
	if e.n.len == 1 {
		l.unsplice(e.n)
	} else {
		copy(e.n.vals[e.i:], e.n.vals[e.i+1:])
		e.n.len--
	}
}

func (l *List[T]) unsplice(n *node[T]) {
	if n.next == nil {
		l.back = n.prev
	} else {
		n.next.prev = n.prev
	}
	if n.prev == nil {
		l.front = n.next
	} else {
		n.prev.next = n.next
	}
}
