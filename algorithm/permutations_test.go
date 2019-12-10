package algorithm

import "testing"

// 给定一个没有重复数字的序列，返回其所有可能的全排列。

// 示例:

// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/permutations

func TestPermute(t *testing.T) {
	t.Log(permute([]int{1, 2, 3}))
}

// result = []
// def backtrack(路径, 选择列表):
//     if 满足结束条件:
//         result.add(路径)
//         return

//     for 选择 in 选择列表:
//         做选择
//         backtrack(路径, 选择列表)
//         撤销选择
func permute(nums []int) [][]int {
	var result [][]int
	var backtracking func(current []int, start int)
	backtracking = func(current []int, start int) {
		if len(current) == start { // 说明start已经是最后一个数字,无法排列
			var v = make([]int, len(current))
			copy(v, current)
			result = append(result, v)
			return
		}
		// 不是最后一个就开始从start遍历,start之前是已经固定好的
		for i := start; i < len(current); i++ {
			current[i], current[start] = current[start], current[i] // 在第i个位置替换其他的数字
			//
			backtracking(current, start+1)
			// 如果不行就回溯
			current[i], current[start] = current[start], current[i]
		}
	}
	backtracking(nums, 0)
	return result
}
