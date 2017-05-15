package zsort

func ShellSort(arr []int) []int {
	l := len(arr)
	gap := 1
	for gap < l/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < l; i++ {
			temp := arr[i]
			index := i
			for index >= gap && arr[index-gap] > temp {
				arr[index] = arr[index-gap]
				index = index - gap
			}
			arr[index] = temp
		}
		gap = gap / 3
	}
	return arr
}
