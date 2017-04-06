package main

import (
	"fmt"
)

func select_sort(s []int) []int {
	n := len(s)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if s[min] > s[j] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
	return s
}

func main() {
	s := []int{1, 3, 9, 0, 2, 4, 7}
	fmt.Println(s, "---default")
	fmt.Println(select_sort(s), "---select")
}
