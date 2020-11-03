package sorts

import (
	"main/stack"
	"math/rand"
)

func randomArea(begin int, end int) int {
	num := rand.Intn(end)
	if num < begin {
		return begin
	}
	return num
}

func Partition(arr []int, begin int, end int) int {
	pivotIndex := randomArea(begin, end)
	pivot := arr[pivotIndex]
	arr[begin], arr[pivotIndex] = arr[pivotIndex], arr[begin]

	left := begin + 1
	right := end

	for left <= right {
		for (left <= end) && (arr[left] <= pivot) {
			left++
		}
		for (right >= begin) && (arr[right] > pivot) {
			right--
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[begin], arr[right] = arr[right], arr[begin]
	return right
}

func PartitionSingle(arr []int, begin int, end int) int {
	pivotIndex := randomArea(begin, end)
	pivot := arr[pivotIndex]
	arr[begin], arr[pivotIndex] = arr[pivotIndex], arr[begin]

	mark := begin

	for i := begin + 1; i <= end; i++ {
		if arr[i] < pivot {
			mark++
			arr[i], arr[mark] = arr[mark], arr[i]
		}
	}

	arr[begin], arr[mark] = arr[mark], arr[begin]
	return mark
}

func QuickSortRecurse(arr []int, begin int, end int, callback func([]int, int, int) int) {
	if begin >= end {
		return
	}
	pivotIndex := callback(arr, begin, end)
	QuickSortRecurse(arr, begin, pivotIndex-1, callback)
	QuickSortRecurse(arr, pivotIndex+1, end, callback)
}

func QuickSortStack(arr []int, begin int, end int, callback func([]int, int, int) int) {
	if begin >= end {
		return
	}
	quickSortStack := stack.New()
	rootParam := make([]int, 2)
	rootParam[0] = begin
	rootParam[1] = end
	quickSortStack.Push(rootParam)

	for !quickSortStack.IsEmpty() {
		param := quickSortStack.Pop()

		paramBegin := param[0]
		paramEnd := param[1]
		pivotIndex := callback(arr, paramBegin, paramEnd)

		if paramBegin < pivotIndex-1 {
			leftParam := make([]int, 2)
			leftParam[0] = paramBegin
			leftParam[1] = pivotIndex - 1
			quickSortStack.Push(leftParam)
		}

		if paramEnd > pivotIndex+1 {
			rightParam := make([]int, 2)
			rightParam[0] = pivotIndex + 1
			rightParam[1] = paramEnd
			quickSortStack.Push(rightParam)
		}
	}
}
