package sorts

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	arr := []float32{4.32, 0.3223, 3.65, 2.6, 3.4, 0.13, 4.86, 2.334, 5.964, 4.123, 2.453}[:]
	targetArr := []float32{0.13, 0.3223, 2.334, 2.453, 2.6, 3.4, 3.65, 4.123, 4.32, 4.86, 5.964}[:]

	BucketSort(arr)
	if !reflect.DeepEqual(arr, targetArr) {
		t.Error()
	}

	arr2 := []float32{7, 6, 3, 2, 6, 5, 3, 1}[:]
	targetArr2 := []float32{1, 2, 3, 3, 5, 6, 6, 7}[:]

	BucketSort(arr2)
	if !reflect.DeepEqual(arr2, targetArr2) {
		t.Error()
	}
}
