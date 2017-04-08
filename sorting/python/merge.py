#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def merge_sort(arr):
    n = len(arr)
    if n <= 1:
        return arr
    num = n / 2
    left = merge_sort(arr[:num])
    right = merge_sort(arr[num:])
    return merge(left, right)

def merge(left, right):
    result = []
    l, r = 0, 0
    while l < len(left) and r < len(right):
        if left[l] < right[r]:
            result.append(left[l])
            l += 1
        else:
            result.append(right[r])
            r += 1
    result += left[l:]
    result += right[r:]
    return result

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7, 8, 6, 5]

    print arr, "---default"
    print merge_sort(arr), "---merge"
