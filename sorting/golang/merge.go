package main

import (
	"fmt"
)

func merge_sort(s []int) []int {
	n := len(s)
	if n <= 1 {
		return s
	}
	num := n / 2
	left := merge_sort(s[:num])
	right := merge_sort(s[num:])
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

func main() {
	s := []int{1, 3, 9, 0, 2, 4, 7}
	fmt.Println(s, "---default")
	fmt.Println(merge_sort(s), "---merge")
}
