#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def quick_sort(arr):
    return quick(arr, 0, len(arr) - 1)

def quick(arr, left, right):
    if left >= right:
        return arr
    l, r = left, right
    key = arr[left]
    while l < r:
        while l < r and arr[r] > key:
            r -= 1
        if l < r:
            arr[l] = arr[r]
            l += 1
        while l < r and arr[l] <= key:
            l += 1
        if l < r:
            arr[r] = arr[l]
            r -= 1
    arr[l] = key
    quick(arr, left, l - 1)
    quick(arr, r + 1, right)
    return arr

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7, 8, 6, 5]

    print arr, "---default"
    print quick_sort(arr), "---quick"
