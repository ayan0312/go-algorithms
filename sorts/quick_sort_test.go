package sorts

import (
	"reflect"
	"testing"
)

func TestQuickSortRecurse(t *testing.T) {
	arr := []int{4, 23, 65, 86, 23, 13, 86, 23, 89, 43, 23}[:]
	targetArr := []int{4, 13, 23, 23, 23, 23, 43, 65, 86, 86, 89}[:]

	arr2 := []int{7, 6, 3, 2, 6, 5, 3, 1}[:]
	targetArr2 := []int{1, 2, 3, 3, 5, 6, 6, 7}[:]

	QuickSortRecurse(arr, 0, len(arr)-1, PartitionSingle)
	if !reflect.DeepEqual(arr, targetArr) {
		t.Error()
	}

	QuickSortRecurse(arr2, 0, len(arr2)-1, Partition)
	if !reflect.DeepEqual(arr2, targetArr2) {
		t.Error()
	}
}

func TestQuickSortStack(t *testing.T) {
	arr := []int{4, 23, 65, 86, 23, 13, 86, 23, 89, 43, 23}[:]
	targetArr := []int{4, 13, 23, 23, 23, 23, 43, 65, 86, 86, 89}[:]

	arr2 := []int{7, 6, 3, 2, 6, 5, 3, 1}[:]
	targetArr2 := []int{1, 2, 3, 3, 5, 6, 6, 7}[:]

	QuickSortStack(arr, 0, len(arr)-1, PartitionSingle)
	if !reflect.DeepEqual(arr, targetArr) {
		t.Error()
	}

	QuickSortStack(arr2, 0, len(arr2)-1, Partition)
	if !reflect.DeepEqual(arr2, targetArr2) {
		t.Error()
	}
}
