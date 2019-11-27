package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := make([]int, len(testCase))
	copy(arr, testCase)
	SelectionSort(arr)
	t.Log(arr)
}

// 每次循环寻找最小的和index交换
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
