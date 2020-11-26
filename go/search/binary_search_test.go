package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 4, 9, 11, 12, 32, 39, 42, 53, 55, 64, 70, 77, 87, 94}
	t.Log(binarySearch(arr, 94))
}

// 查找第一个相等的
func binarySearch(arr []int, n int) int {
	var (
		low  = 0
		high = len(arr) - 1
	)
	mid := (low + (high - low)) / 2
	for low <= high {
		if arr[mid] > n {
			high = mid - 1
		} else if arr[mid] < n {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != n {
				return mid
			}
			high = mid - 1
		}
		mid = (low + (high - low)) / 2
	}

	return mid
}
