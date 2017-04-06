package main

import (
	"fmt"
)

func bubble_sort(s []int) []int {
	count := 0
	n := len(s)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			count++
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
	fmt.Println(count)
	return s
}

func bubble_sort2(s []int) []int {
	count := 0
	n := len(s)
	for i := 0; i < n; i++ {
		flag := true
		for j := 1; j < n-i; j++ {
			count++
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	fmt.Println(count)
	return s
}

func bubble_sort3(s []int) []int {
	count := 0
	n := len(s)
	last := n
	for i := 0; i < n; i++ {
		flag := true
		temp := last
		for j := 1; j < temp; j++ {
			count++
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				flag = false
				last = j
			}
		}
		if flag {
			break
		}
	}
	fmt.Println(count)
	return s
}

func main() {
	s := []int{1, 3, 9, 0, 2, 4, 7}
	fmt.Println(s, "---default")
	// fmt.Println(bubble_sort(s), "---bubble")
	// fmt.Println(bubble_sort2(s), "---bubble2")
	fmt.Println(bubble_sort3(s), "---bubble3")
}
