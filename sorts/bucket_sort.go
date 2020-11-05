package sorts

func BucketSort(arr []float32) {
	max := arr[0]
	min := arr[0]
	length := len(arr)
	for i := 1; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	d := max - min
	bucketArr := make([][]float32, length)

	for i := 0; i < length; i++ {
		num := int((arr[i] - min) * (float32(length) - 1) / d)
		bucketArr[num] = append(bucketArr[num], arr[i])
	}

	for i := 0; i < len(bucketArr); i++ {
		HeapSortFloat32(bucketArr[i])
	}

	sortedArr := make([]float32, length)
	index := 0
	for i := 0; i < len(bucketArr); i++ {
		bucket := bucketArr[i]
		for j := 0; j < len(bucket); j++ {
			sortedArr[index] = bucket[j]
			index++
		}
	}
	copy(arr, sortedArr)
}
