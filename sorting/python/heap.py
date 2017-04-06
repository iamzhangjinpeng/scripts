#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def heap_sort(arr):
    n = len(arr)
    first = n / 2 - 1
    for start in range(first, -1, -1):
        max_heapify(arr, start, n - 1)
    for end in range(n - 1, 0, -1):
        arr[end], arr[0] = arr[0], arr[end]
        max_heapify(arr, 0, end - 1)
    return arr

def max_heapify(arr, start, end):
    root = start
    while True:
        child = root * 2 + 1
        if child > end:
            break
        if child + 1 <= end and arr[child] < arr[child + 1]:
            child = child + 1
        if arr[root] < arr[child]:
            arr[root], arr[child] = arr[child], arr[root]
            root = child
        else:
            break

def min_heapify(arr, start, end):
    root = start
    while True:
        child = root * 2 + 1
        if child > end:
            break
        if child + 1 <= end and arr[child] > arr[child + 1]:
            child = child + 1
        if arr[root] > arr[child]:
            arr[root], arr[child] = arr[child], arr[root]
            root = child
        else:
            break

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7]

    print arr, "---default"
    print heap_sort(arr), "---heap"
