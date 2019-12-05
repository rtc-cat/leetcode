package array

import "testing"

// 给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

// 示例 1:

// 输入: [1,2,0]
// 输出: 3
// 示例 2:

// 输入: [3,4,-1,1]
// 输出: 2
// 示例 3:

// 输入: [7,8,9,11,12]
// 输出: 1
// 说明:

// 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/first-missing-positive

func TestFirstMissingPositive(t *testing.T) {
	testCases := []struct{ input []int }{
		{[]int{1, 2, 0}},
		{[]int{3, 4, -1, 1}},
		{[]int{7, 8, 9, 11, 12}},
		{[]int{-5, 10000}},
		{[]int{1}},
	}
	for _, tc := range testCases {
		t.Log(firstMissingPositive(tc.input))
	}
}

func firstMissingPositive(nums []int) int {
	if len(nums) < 1 {
		return 1
	}
	// 没有缺失的数组应该是1,2,3...,n
	// 将数字交换到它该去的位置,不属于这个数组的就不动了
	for i := 0; i < len(nums); {
		// 如果这里的数字没办法交换位置就不动了,寻找下一个
		// 小于零,大于n,或者对应的那个位置已经有数字
		if nums[i] <= 0 ||
			nums[i] > len(nums) ||
			nums[nums[i]-1] == nums[i] {
			i++
			continue
		}
		// 否则就交换
		nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
