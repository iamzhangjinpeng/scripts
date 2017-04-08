#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def bubble_sort(arr):
    n = len(arr)
    for i in range(n - 1):
        for j in range(1, n - i):
            if arr[j - 1] > arr[j]:
                arr[j - 1], arr[j] = arr[j], arr[j - 1]
    return arr

def bubble_sort2(arr):
    n = len(arr)
    for i in range(n - 1):
        flag = True
        for j in range(1, n - i):
            if arr[j - 1] > arr[j]:
                arr[j - 1], arr[j] = arr[j], arr[j - 1]
                flag = False
        if flag:
            break
    return arr

def bubble_sort3(arr):
    n = len(arr)
    last = n
    for i in range(n - 1):
        flag = True
        for j in range(1, last):
            if arr[j - 1] > arr[j]:
                arr[j - 1], arr[j] = arr[j], arr[j - 1]
                flag = False
                last = j
        if flag:
            break
    return arr

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7, 8, 6, 5]

    print arr, "---default"
    # print bubble_sort(arr), "---bubble"
    # print bubble_sort2(arr), "---bubble2"
    print bubble_sort3(arr), "---bubble3"
