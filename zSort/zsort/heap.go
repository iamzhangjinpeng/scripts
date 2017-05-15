package zsort

func HeapSort(arr []int) []int {
	l := len(arr)
	for i := l/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, l-1)
	}
	for j := l - 1; j >= 0; j-- {
		arr[j], arr[0] = arr[0], arr[j]
		maxHeapify(arr, 0, j-1)
	}
	return arr
}

func maxHeapify(arr []int, left int, right int) {
	root := left
	for {
		child := root*2 + 1
		if child > right {
			break
		}
		if child+1 <= right && arr[child] < arr[child+1] {
			child = child + 1
		}
		if arr[root] < arr[child] {
			arr[root], arr[child] = arr[child], arr[root]
			root = child
		} else {
			break
		}
	}
}

func minHeapify(arr []int, left int, right int) {
	root := left
	for {
		child := root*2 + 1
		if child > right {
			break
		}
		if child+1 <= right && arr[child] > arr[child+1] {
			child = child + 1
		}
		if arr[root] > arr[child] {
			arr[root], arr[child] = arr[child], arr[root]
			root = child
		} else {
			break
		}
	}
}
