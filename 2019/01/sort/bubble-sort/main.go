package main

import "fmt"

// 冒泡排序算法的运作如下：

// 1.比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 3.针对所有的元素重复以上的步骤，除了最后一个。
// 4.持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
func main() {
	// a := []int{6, 5, 4, 3, 1, 2}
	arrs := [][]int{
		{6, 5, 4, 3, 1, 2},
		{22, 34, 3, 40, 18, 4},
		{},
		{1, 2, 3},
	}
	for _, arr := range arrs {
		bubbleSort(arr)
		fmt.Println(arr)
	}
}

// 冒泡排序
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
