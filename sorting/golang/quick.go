package main

import (
	"fmt"
)

func quick_sort(s []int) []int {
	return quick(s, 0, len(s)-1)
}

func quick(s []int, left int, right int) []int {
	if left >= right {
		return s
	}
	key := s[left]
	l, r := left, right
	for l < r {
		for s[r] > key && l < r {
			r -= 1
		}
		if l < r {
			s[l] = s[r]
			l += 1
		}
		for s[l] < key && l < r {
			l += 1
		}
		if l < r {
			s[r] = s[l]
			r -= 1
		}
	}
	s[l] = key
	quick(s, left, l-1)
	quick(s, r+1, right)
	return s
}

func main() {
	s := []int{1, 3, 9, 0, 2, 4, 7}
	fmt.Println(s, "---default")
	fmt.Println(quick_sort(s), "---quick")
}
