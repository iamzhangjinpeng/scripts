#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def select_sort(arr):
    n = len(arr)
    for i in range(0, n):
        min = i
        for j in range(i + 1, n):
            if arr[min] > arr[j]:
                min = j
        arr[min], arr[i] = arr[i], arr[min]
    return arr

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7]

    print arr, "---default"
    print select_sort(arr), "---select"
