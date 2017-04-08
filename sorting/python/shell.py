#!/usr/local/bin/python2
# -*- coding: UTF-8 -*-

def shell_sort(arr):
    n = len(arr)
    gap = n / 2
    while gap > 0:
        for i in range(gap, n):
            temp = arr[i]
            index = i
            for j in range(i - gap, -1, -gap):
                if arr[j] > temp:
                    arr[j + gap] = arr[j]
                    index = j
                else:
                    break
            arr[index] = temp
        gap /= 2
    return arr

def shell_sort2(arr):
    n = len(arr)
    gap = n / 2
    while gap > 0:
        for i in range(gap, n):
            temp = arr[i]
            j = i
            while (j - gap) >= 0 and arr[j - gap] > temp:
                arr[j] = arr[j - gap]
                j -= gap
            arr[j] = temp
        gap /= 2
    return arr

if __name__ == "__main__":
    arr = [1, 3, 9, 0, 2, 4, 7, 8, 6, 5]

    print arr, "---default"
    # print shell_sort(arr), "---shell"
    print shell_sort2(arr), "---shell2"
