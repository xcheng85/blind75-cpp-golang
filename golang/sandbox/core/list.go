package core

import "fmt"

type node[T any] struct {
	value T
	next  *node[T]
}

type list[T any] struct {
	head *node[T]
	tail *node[T]
}

// ctor pattern in go
func NewList[T any]() *list[T] {
	dummy := new(node[T])
	tail := new(node[T])
	dummy.next = tail
	tail.next = nil
	return &list[T]{dummy, tail}
}

// point receiver

func (l *list[T]) PushFront(v T) {
	newNode := node[T]{
		value: v,
		next:  nil,
	}
	tmp := l.head.next
	l.head.next = &newNode
	newNode.next = tmp
}

// ostream in c++
func (l list[T]) Debug() {
	curr := l.head.next

	for {
		if curr == l.tail {
			break
		}
		fmt.Println(curr.value)
		curr = curr.next
	}
}
