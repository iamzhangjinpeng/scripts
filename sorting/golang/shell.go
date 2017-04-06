package main

import (
	"fmt"
)

func shell_sort(s []int) []int {
	n := len(s)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := s[i]
			index := i
			for index >= gap && s[index-gap] > temp {
				s[index] = s[index-gap]
				index = index - gap
			}
			s[index] = temp
		}
		gap = gap / 2
	}
	return s
}

func main() {
	s := []int{1, 3, 9, 0, 2, 4, 7}
	fmt.Println(s, "---default")
	fmt.Println(shell_sort(s), "---shell")
}
