package doublelinkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New()

	if s.Count() != 0 {
		t.Error()
	}
}

func TestDoubleLinkedList(t *testing.T) {
	ll := New()

	for i := 0; i < 15; i++ {
		ll.Append(i)
	}

	for i := 0; i < 15; i++ {
		el, err := ll.Delete(0)
		if err != nil || el.Data != i {
			t.Error()
		}
	}
}
