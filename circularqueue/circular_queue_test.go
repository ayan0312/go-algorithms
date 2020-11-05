package circularqueue

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New(10)

	s.EnQueue(1)
	el, err := s.DeQueue()
	if err != nil || el.item != 1 {
		t.Error()
	}
}

func TestEnQueue(t *testing.T) {
	s := New(10)

	for i := 0; i < 9; i++ {
		err := s.EnQueue(i)
		if err != nil {
			t.Error()
		}
	}

	for i := 0; i < 6; i++ {
		el, err := s.DeQueue()
		if err != nil || el.item != i {
			t.Error()
		}
	}

	for i := 9; i < 13; i++ {
		err := s.EnQueue(i)
		if err != nil {
			t.Error()
		}
	}
}

func TestDeQueue(t *testing.T) {
	s := New(10)

	for i := 0; i < 9; i++ {
		s.EnQueue(i)
	}

	for i := 0; i < 6; i++ {
		el, err := s.DeQueue()
		if err != nil || el.item != i {
			t.Error()
		}
	}

	for i := 9; i < 13; i++ {
		s.EnQueue(i)
	}

	for i := 6; i < 12; i++ {
		el, err := s.DeQueue()
		if err != nil || el.item != i {
			t.Error()
		}
	}
}
