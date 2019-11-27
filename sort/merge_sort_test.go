package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := make([]int, len(testCase))
	copy(arr, testCase)
	arr = MergeSort(arr)
	t.Log(arr)
}

func MergeSort(arr []int) []int {
	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	// 递归退出条件
	if len(arr) < 2 {
		return arr
	}
	m := len(arr) / 2
	left := mergeSort(arr[:m])
	right := mergeSort(arr[m:])
	return merge(left, right)
}

func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	var result []int
	for i < len(arr1) || j < len(arr2) {
		if i >= len(arr1) {
			result = append(result, arr2[j:]...)
			break
		}
		if j >= len(arr2) {
			result = append(result, arr1[i:]...)
			break
		}
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	return result
}
