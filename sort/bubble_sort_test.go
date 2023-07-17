package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := make([]int, len(testCase))
	copy(arr, testCase)
	BubbleSort(arr)
	t.Log(arr)
}

// 每次对比两个相邻的数,交换位置
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
