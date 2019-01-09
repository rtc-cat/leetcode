package main

import "fmt"

// 选择排序（Selection sort）是一种简单直观的排序算法。
// 它的工作原理如下。首先在未排序序列中找到最小（大）元素，
// 存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，
// 然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。
func main() {
	arrs := [][]int{
		{6, 5, 4, 3, 1, 2},
		{22, 34, 3, 40, 18, 4},
		{},
		{1, 2, 3},
	}
	for _, arr := range arrs {
		selectionSort(arr)
		fmt.Println(arr)
	}
}

func selectionSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr); i++ {
		j := i
		min := i
		for ; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// 拿到最小,交换当前位置
		arr[i], arr[min] = arr[min], arr[i]
	}
}
