package sorts

func CountingSort(arr []int) {
	max := arr[0]
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	d := max - min
	countArr := make([]int, d+1)
	for i := 0; i < len(arr); i++ {
		countArr[arr[i]-min]++
	}

	for i := 1; i < len(countArr); i++ {
		countArr[i] += countArr[i-1]
	}

	sortedArr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		countIndex := arr[i] - min
		sortedArr[countArr[countIndex]-1] = arr[i]
		countArr[countIndex]--
	}

	copy(arr, sortedArr)
}
