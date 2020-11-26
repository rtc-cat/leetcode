package array

import "testing"

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// 示例 1:

// 输入: [3,2,3]
// 输出: 3
// 示例 2:

// 输入: [2,2,1,1,1,2,2]
// 输出: 2

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/majority-element

func TestMajorityElement(t *testing.T) {
	testCases := []struct{ input []int }{
		{[]int{3, 2, 3}},
		{[]int{2, 2, 1, 1, 1, 2, 2}},
	}
	for _, tc := range testCases {
		t.Log(majorityElement(tc.input))
	}
}

// Boyer-Moore 投票算法
func majorityElement(nums []int) int {
	var result int
	var count int
	for i := range nums {
		if count == 0 {
			result = nums[i]
			count++
			continue
		}
		if result == nums[i] {
			count++
		} else {
			count--
		}
	}
	return result
}
