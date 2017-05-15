package zsort

func BubbleSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func BubbleSort2(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		flag := true
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return arr
}

func BubbleSort3(arr []int) []int {
	l := len(arr)
	last := l - 1
	for i := 0; i < l-1; i++ {
		flag := true
		temp := last
		for j := 0; j < temp; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
				last = j
			}
		}
		if flag {
			break
		}
	}
	return arr
}
