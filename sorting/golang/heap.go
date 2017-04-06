package main

import (
	"fmt"
)

func heap_sort(s []int) []int {
	n := len(s)
	first := n/2 - 1
	for start := first; start >= 0; start-- {
		// max_heapify(s, start, n-1)
		min_heapify(s, start, n-1)
	}
	for end := n - 1; end >= 0; end-- {
		s[end], s[0] = s[0], s[end]
		// max_heapify(s, 0, end-1)
		min_heapify(s, 0, end-1)
	}
	return s
}

func max_heapify(s []int, start int, end int) {
	root := start
	for {
		child := root*2 + 1
		if child > end {
			break
		}
		if child+1 < end && s[child] < s[child+1] {
			child = child + 1
		}
		if s[root] < s[child] {
			s[root], s[child] = s[child], s[root]
			root = child
		} else {
			break
		}
	}
}

func min_heapify(s []int, start int, end int) {
	root := start
	for {
		child := root*2 + 1
		if child > end {
			break
		}
		if child+1 < end && s[child] > s[child+1] {
			child = child + 1
		}
		if s[root] > s[child] {
			s[root], s[child] = s[child], s[root]
			root = child
		} else {
			break
		}
	}
}

func main() {
	s := []int{1, 3, 9, 0, 2, 4, 7}
	fmt.Println(s, "---default")
	fmt.Println(heap_sort(s), "---heap")
}
