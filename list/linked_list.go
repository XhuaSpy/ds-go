package list

import "fmt"

type nodeDoubly[T any] struct {
	data T
	next *nodeDoubly[T]
	prev *nodeDoubly[T]
}

func (n nodeDoubly[T]) String() string {
	return fmt.Sprint("%v", n.data)
}

type LinkedList[T any] struct {
	head nodeDoubly[T]
	tail nodeDoubly[T]
}


