package main

import "fmt"

// 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
// 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 8))
	fmt.Println(hashTwoSum([]int{1, 2, 3, 4, 5}, 8))
	fmt.Println(hashOnceTwoSum([]int{3, 3}, 6))
}

// 两次循环
// 返回index
func twoSum(nums []int, target int) []int {
	// 特殊情况判断
	// 空数组
	if len(nums) == 0 {
		return []int{}
	}
	// 只有一个元素
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0}
		}
	}
	// 正常判断,两次循环,i从0开始,j从i+1开始
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// 符合条件就停止循环
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 两次遍历hash
// 第一次遍历,建立索引
// 第二次遍历,根据索引查询是否有匹配
func hashTwoSum(nums []int, target int) []int {
	// ? 优化,初始化指定大小
	m := make(map[int]int, len(nums))
	// 第一次遍历
	for i := range nums {
		m[nums[i]] = i
	}
	// 第二次遍历
	for i := range nums {
		v := target - nums[i]
		// 如果有对应的index,则返回
		if index, ok := m[v]; ok && i != index {
			return []int{i, index}
		}
	}
	return []int{}
}

// 上面算法的改进
func hashOnceTwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		// 检查map是否已有
		if i, ok := m[v]; ok {
			// 如果已经有
			return []int{i, k}
		}
		m[target-v] = k
	}
	return []int{}
}
