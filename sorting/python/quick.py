#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def quick_sort(arr):
    return quick(arr, 0, len(arr) - 1)

def quick(arr, left, right):
    if left >= right:
        return arr
    key = arr[left]
    l, r = left, right
    while l < r:
        while arr[r] >= key and l < r:
            r -= 1
        if l < r:
            arr[l] = arr[r]
            l += 1
        while arr[l] <= key and l < r:
            l += 1
        if l < r:
            arr[r] = arr[l]
            r -= 1
    arr[l] = key
    quick(arr, left, l - 1)
    quick(arr, r + 1, right)
    return arr

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7]

    print arr, "---default"
    print quick_sort(arr), "---quick"
