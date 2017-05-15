package zsort

func SelectSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
