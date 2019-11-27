package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := make([]int, len(testCase))
	copy(arr, testCase)
	QuickSort(arr)
	t.Log(arr)
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(arr, low, high)
	quickSort(arr, low, p-1)
	quickSort(arr, p+1, high)
}

func partition(arr []int, low, high int) int {
	// 取最后一个
	pivot := arr[high]
	p := low // 指向需要交换的位置,也就是第一个比pivot大的值
	for i := low; i < high; i++ {
		if arr[i] < pivot {
			arr[i], arr[p] = arr[p], arr[i]
			p++
		}
	}
	arr[p], arr[high] = arr[high], arr[p]
	return p
}
