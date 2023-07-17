package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	arr := make([]int, len(testCase))
	copy(arr, testCase)
	InsertionSort(arr)
	t.Log(arr)
}

// 找到合适的位置,其余的全部右移一位
func InsertionSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		tmp := arr[i] // 需要插入的数字
		// 寻找k, 当arr[k] > tmp时 右移一位
		var k int
		for k = i - 1; k >= 0 && arr[k] > tmp; k-- {
			arr[k+1] = arr[k]
		}
		arr[k+1] = tmp
	}
}
