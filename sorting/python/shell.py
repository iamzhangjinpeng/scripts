#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def shell_sort(arr):
    n = len(arr)
    gap = n / 2
    while gap > 0:
        for i in range(gap, n):
            temp = arr[i]
            index = i
            while (index >= gap and arr[index - gap] > temp):
                arr[index] = arr[index - gap]
                index = index - gap
            arr[index] = temp
        gap = gap / 2
    return arr

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7]

    print arr, "---default"
    print shell_sort(arr), "---shell"
