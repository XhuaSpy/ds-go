package list

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	next *node[T]
	data T
}

func (n *node[T]) String() string {
	return fmt.Sprintf("%v", n.data)
}

type SimpleLinkedList[T any] struct {
	head *node[T]
	tail *node[T]
}

func (l *SimpleLinkedList[T]) InsertLast(data T) {
	new_nodo := &node[T]{data: data}

	if l.head != nil {
		l.tail.next = new_nodo
		l.tail = new_nodo
		return
	}

	l.head = new_nodo
	l.tail = new_nodo
}

func (l *SimpleLinkedList[T]) Get(id int) (T, bool) {
	var zero T

	if l.head == nil {
		return zero, false
	}

	current := l.head
	var count int

	for current != nil && count <= id {
		if count == id {
			return current.data, true
		}

		count++
		current = current.next
	}

	return zero, true
}

func (l *SimpleLinkedList[T]) ToString() {
	if l.head == nil {
		return
	}

	current := l.head

	for current != nil {
		println(current.data)
		current = current.next
	}
}

func (l SimpleLinkedList[T]) String() string {
	var sb strings.Builder
	current := l.head

	sb.WriteString("[")

	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.data))
		if current.next != nil {
			sb.WriteString(", ")
		}
		current = current.next
	}

	sb.WriteString("]")
	return sb.String()
}
