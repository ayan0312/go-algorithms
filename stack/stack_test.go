package stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New()

	if !s.IsEmpty() ||
		s.depth != 0 ||
		s.Count() != 0 {
		t.Error()
	}
}
