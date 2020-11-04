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
