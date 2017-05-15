package main

import (
	"fmt"
	"zSort/zsort"
)

func main() {
	fmt.Println("hello world")

	arr := []int{3, 13, 9, 54, 4, 2, 21, 0, 1, 37, 7, 5, 8, 78, 6, 1, 43, 3, 7, 127, 44, 88}

	fmt.Println(arr, "---default")

	//fmt.Println(zsort.BubbleSort(arr), "---bubble")
	//fmt.Println(zsort.BubbleSort2(arr), "---bubble2")
	//fmt.Println(zsort.BubbleSort3(arr), "---bubble3")

	//fmt.Println(zsort.SelectSort(arr), "---select")

	//fmt.Println(zsort.InsertSort(arr), "---insert")

	//fmt.Println(zsort.ShellSort(arr), "---shell")

	//fmt.Println(zsort.MergeSort(arr), "---merge")

	//fmt.Println(zsort.QuickSort(arr), "---quick")

	fmt.Println(zsort.HeapSort(arr), "---heap")

	//fmt.Println(zsort.CountSort(arr), "---count")

	//fmt.Println(zsort.BucketSort(arr), "---bucket")

	//fmt.Println(zsort.RadixSort(arr), "---radix")
}
