package zsort

func MergeSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	middle := l / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := []int{}
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l += 1
		} else {
			result = append(result, right[r])
			r += 1
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
