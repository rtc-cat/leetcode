package main

import (
	"fmt"
)

// 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

// 说明:

// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
// 示例:

// 输入:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// 输出: [1,2,2,3,5,6]
func main() {
	nums1 := []int{0, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	merge(nums1, 0, nums2, 3)
	fmt.Println(nums1)
}

// 条件是排序数组,而且nums1空间很大,所以可以直接从数组末尾开始往回写,这样就不涉及插入操作,不需要移位
func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		i = m - 1
		j = n - 1
		k = m + n - 1 //最后一位,也是最大的一位
	)
	for i >= 0 && j >= 0 {
		x := nums1[i]
		y := nums2[j]
		if x > y {
			nums1[k] = x
			i--
		} else {
			nums1[k] = y
			j--
		}
		k--
	}
	// 这里需要判断,如果是nums2的还没有结束,需要继续
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
	return
}
