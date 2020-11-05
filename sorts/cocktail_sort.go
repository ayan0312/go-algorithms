package sorts

func CocktailSort(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		isSorted := true
		for j := i; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
		isSorted = true
		for j := len(arr) - i - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}
