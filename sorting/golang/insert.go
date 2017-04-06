package main

import (
	"fmt"
)

func insert_sort(s []int) []int {
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] < s[i - 1] {
			temp := s[i]
			index := i
			for j := i - 1; j >= 0; j-- {
				if s[j] > temp {
					s[j + 1] = s[j]
					index = j
				} else {
					break
				}
			}
			s[index] = temp
		}
	}
	return s
}

func main() {
	s := []int{1, 3, 9, 0, 2, 4, 7}
	fmt.Println(s, "---default")
	fmt.Println(insert_sort(s), "---insert")
}
