package sorts

func BubbleSort(arr []int) {
	lastExchangeIndex := 0
	sortBorder := len(arr) - 1
	for i := 0; i < len(arr)-1; i++ {
		isSorted := true
		for j := 0; j < sortBorder; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSorted = false
				lastExchangeIndex = j
			}
		}
		sortBorder = lastExchangeIndex
		if isSorted {
			break
		}
	}
}
