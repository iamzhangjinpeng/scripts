package zsort

func CountSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	max := getMax(arr)
	bucket := make([]int, max+1)
	index := 0

	for i := 0; i < l; i++ {
		bucket[arr[i]] += 1
	}
	for j := 0; j < max+1; j++ {
		for bucket[j] > 0 {
			arr[index] = j
			index += 1
			bucket[j] -= 1
		}
	}

	return arr
}
