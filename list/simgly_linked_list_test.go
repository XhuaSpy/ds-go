package list

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedListInsertLast(t *testing.T) {
	t.Run("add int", func(t *testing.T) {
		list := SimpleLinkedList[int]{}

		a := 1
		b := 2

		list.InsertLast(a)
		list.InsertLast(b)

		value, ok := list.Get(0)

		if !ok && value != a {
			t.Errorf("got %q want %q", value, a)
		}

		value, ok = list.Get(1)

		if !ok && value != b {
			t.Errorf("got %q want %q", value, b)
		}
	})

}

func TestSinglyLinkedListString(t *testing.T) {
	list := SimpleLinkedList[int]{}

	a, b := 1, 2

	list.InsertLast(a)
	list.InsertLast(b)

	if list.String() != fmt.Sprintf("[%d,%d]", a, b) {
		t.Errorf("got %q want %q", fmt.Sprintf("[%d,%d]", a, b), list.String())
	}
}
