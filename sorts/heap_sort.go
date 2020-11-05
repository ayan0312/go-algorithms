package sorts

import "main/heap"

type Int int

func (x Int) Less(than heap.Item) bool {
	return x < than.(Int)
}

type Float32 float32

func (x Float32) Less(than heap.Item) bool {
	return x < than.(Float32)
}

func HeapSort(arr []int) {
	h := heap.NewMin()
	for i := 0; i < len(arr); i++ {
		h.Insert(Int(arr[i]))
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = int(h.Extract().(Int))
	}
}

func HeapSortFloat32(arr []float32) {
	h := heap.NewMin()
	for i := 0; i < len(arr); i++ {
		h.Insert(Float32(arr[i]))
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = float32(h.Extract().(Float32))
	}
}
