#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def insert_sort(arr):
    n = len(arr)
    for i in range(1, n):
        if arr[i - 1] > arr[i]:
            temp = arr[i]
            index = i
            for j in range(i - 1, -1, -1):
                if arr[j] > temp:
                    arr[j + 1] = arr[j]
                    index = j
                else:
                    break
            arr[index] = temp
    return arr

def insert_sort2(arr):
    n = len(arr)
    for i in range(1, n):
        if arr[i - 1] > arr[i]:
            temp = arr[i]
            j = i
            while (j - 1) >= 0 and arr[j - 1] > temp:
                arr[j] = arr[j - 1]
                j -= 1
            arr[j] = temp
    return arr

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7, 8, 6, 5]

    print arr, "---default"
    # print insert_sort(arr), "---insert"
    print insert_sort2(arr), "---insert2"
