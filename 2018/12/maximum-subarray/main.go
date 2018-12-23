package main

import "fmt"

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// 示例:

// 输入: [-2,1,-3,4,-1,2,1,-5,4],
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 进阶:

// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
func maxSubArray(nums []int) int {
	max := -(1 << 63) // 先将最大值设置为 最小的负数
	crt := 0          // current 当前最大子数组的和
	// 遍历数组
	for _, v := range nums {
		// 如果加上之前的和比自己小,说明之前的是负数
		if crt+v < v {
			// 那就把当前数字设置为新的数组的开始
			crt = v
		} else {
			// 否则就加上当前数字
			crt += v
		}
		// 如果当前和比之前的max大,那么就替换掉max
		if crt > max {
			max = crt
		}
	}

	return max
}
