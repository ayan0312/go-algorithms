package singlelinkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New()

	if s.size != 0 {
		t.Error()
	}
}
