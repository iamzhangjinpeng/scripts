package zsort

func getMax(arr []int) int {
	l := len(arr)
	max := arr[0]
	for i := 1; i < l; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func getMin(arr []int) int {
	l := len(arr)
	min := arr[0]
	for i := 1; i < l; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}
