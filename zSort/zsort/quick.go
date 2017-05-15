package zsort

func QuickSort(arr []int) []int {
	return quick(arr, 0, len(arr)-1)
}

func quick(arr []int, left int, right int) []int {
	if left >= right {
		return arr
	}

	pivot := arr[left]
	l, r := left, right
	for l < r {
		for arr[r] > pivot && l < r {
			r -= 1
		}
		if l < r {
			arr[l] = arr[r]
			l += 1
		}
		for arr[l] <= pivot && l < r {
			l += 1
		}
		if l < r {
			arr[r] = arr[l]
			r -= 1
		}
	}
	arr[l] = pivot

	quick(arr, left, l-1)
	quick(arr, r+1, right)

	return arr
}
