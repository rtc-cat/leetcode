package sort

import "testing"

func TestShellSort(t *testing.T) {
	arr := make([]int, len(testCase))
	copy(arr, testCase)
	ShellSort(arr)
	t.Log(arr)
}

func ShellSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	// 指定一个间隔,然后使用插入排序
	for step := len(arr) / 2; step > 0; step /= 2 {
		for i := step; i < len(arr); i++ {
			// 插入排序
			insert(arr, step, i)
		}
	}
}

func insert(arr []int, step, i int) {
	tmp := arr[i]
	var k int
	for k = i - step; k >= 0 && arr[k] > tmp; k -= step {
		arr[k+step] = arr[k]
	}
	arr[k+step] = tmp
}
