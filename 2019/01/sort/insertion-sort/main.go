package main

import "fmt"

// 插入排序
// 一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：

// 1. 从第一个元素开始，该元素可以认为已经被排序
// 2. 取出下一个元素，在已经排序的元素序列中从后向前扫描
// 3. 如果该元素（已排序）大于新元素，将该元素移到下一位置
// 4. 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
// 5. 将新元素插入到该位置后
// 6. 重复步骤2~5
func main() {
	arrs := [][]int{
		{6, 5, 4, 3, 1, 2},
		{22, 34, 3, 40, 18, 4},
		{},
		{1, 2, 3},
	}
	for _, arr := range arrs {
		insertionSort(arr)
		fmt.Println(arr)
	}
}

func insertionSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		value := arr[i] // 这个是准备插入的数字
		j := i - 1
		for ; j >= 0; j-- {
			if value < arr[j] {
				arr[j+1] = arr[j] // 移动所有数据
			} else {
				break
			}
		}
		arr[j+1] = value
	}
}
