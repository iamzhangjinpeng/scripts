#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def heap_sort(arr):
    n = len(arr)
    last = n / 2 - 1
    for i in range(last, -1, -1):
        max_heapify(arr, i, n - 1)
    for j in range(n - 1, 0, -1):
        arr[j], arr[0] = arr[0], arr[j]
        max_heapify(arr, 0, j - 1)
    return arr

def max_heapify(arr, left, right):
    root = left
    while True:
        child = root * 2 + 1
        if child > right:
            break
        if child + 1 <= right and arr[child] < arr[child + 1]:
            child = child + 1
        if arr[root] < arr[child]:
            arr[root], arr[child] = arr[child], arr[root]
            root = child
        else:
            break
    return

def min_heapify(arr, left, right):
    root = left
    while True:
        child = root * 2 + 1
        if child > right:
            break
        if child + 1 <= right and arr[child] > arr[child + 1]:
            child = child + 1
        if arr[root] > arr[child]:
            arr[root], arr[child] = arr[child], arr[root]
            root = child
        else:
            break
    return

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7, 8, 6, 5]

    print arr, "---default"
    print heap_sort(arr), "---heap"
