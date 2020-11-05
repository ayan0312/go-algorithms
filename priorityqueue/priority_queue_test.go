package priorityqueue

import (
	"fmt"
	"testing"
)

func TestMaxPriorityQueue(t *testing.T) {
	h := NewMax()

	h.Insert(*NewItem(8, 14))
	h.Insert(*NewItem(7, 232))
	h.Insert(*NewItem(6, 12))
	h.Insert(*NewItem(3, 53))
	h.Insert(*NewItem(1, 123))
	h.Insert(*NewItem(0, 45))
	h.Insert(*NewItem(2, 436))
	h.Insert(*NewItem(4, 117))
	h.Insert(*NewItem(9, 1438))
	h.Insert(*NewItem(5, 159))

	sorted := make([]Item, 0)

	for h.Len() > 0 {
		sorted = append(sorted, h.Extract())
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i].Priority < sorted[i+1].Priority {
			fmt.Println(sorted)
			t.Error()
		}
	}
}

func TestMinPriorityQueue(t *testing.T) {
	h := NewMin()

	h.Insert(*NewItem(8, 14))
	h.Insert(*NewItem(7, 232))
	h.Insert(*NewItem(6, 12))
	h.Insert(*NewItem(3, 53))
	h.Insert(*NewItem(1, 123))
	h.Insert(*NewItem(0, 45))
	h.Insert(*NewItem(2, 436))
	h.Insert(*NewItem(4, 117))
	h.Insert(*NewItem(9, 1438))
	h.Insert(*NewItem(5, 159))

	sorted := make([]Item, 0)
	for h.Len() > 0 {
		sorted = append(sorted, h.Extract())
	}

	for i := 0; i < len(sorted)-2; i++ {
		if sorted[i].Priority > sorted[i+1].Priority {
			fmt.Println(sorted)
			t.Error()
		}
	}
}
