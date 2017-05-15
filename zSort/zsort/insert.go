package zsort

func InsertSort(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		temp := arr[i]
		index := i
		for index >= 1 && arr[index-1] > temp {
			arr[index] = arr[index-1]
			index = index - 1
		}
		arr[index] = temp
	}
	return arr
}
