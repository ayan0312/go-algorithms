package array

import (
	"testing"
)

func TestAdd(t *testing.T) {
	da := New()
	da.Add(3)
	da.Add(2)
	da.Add(1)
	da.Add(4)

	num, err := da.Element(3)
	if err != nil || num != 4 {
		t.Error()
	}
}

func TestInsert(t *testing.T) {
	da := New()
	da.Add(1)
	da.Add(2)
	da.Add(3)
	da.Add(5)
	da.Add(6)
	da.Insert(4, 3)

	for i := 0; i < da.size; i++ {
		num, err := da.Element(i)
		if err != nil {
			t.Error(err)
		}

		if num != i+1 {
			t.Error()
		}
	}

	if da.size != 6 {
		t.Error()
	}
}

func TestInsert2(t *testing.T) {
	da := New()
	da.Add(1)
	da.Add(2)
	da.Add(3)
	da.Add(5)
	da.Add(6)
	da.Insert2(4, 3)

	for i := 0; i < da.size; i++ {
		num, err := da.Element(i)
		if err != nil {
			t.Error(err)
		}

		if num != i+1 {
			t.Error()
		}
	}

	if da.size != 6 {
		t.Error()
	}
}
