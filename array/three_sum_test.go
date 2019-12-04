package array

import (
	"testing"
)

// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

// 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

// 满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{[]int{-2, 0, 0, 2, 2}},
		{[]int{-1, 0, 1, 2, -1, -4}},
		{[]int{1, -1, -1, 0}},
	}
	for _, tc := range testCases {
		t.Log(threeSum(tc.input))
	}
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	// 排序
	quickSort(nums, 0, len(nums)-1)
	var result [][]int
	var visited = make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return result
		}
		if visited[nums[i]] {
			continue
		}
		visited[nums[i]] = true // 标记为已判定
		left := i + 1
		right := len(nums) - 1
		var localVisited = make(map[int]bool, right-left)
		for left < right {
			if localVisited[nums[left]] {
				left++
				continue
			}
			if localVisited[nums[right]] {
				right--
				continue
			}
			r := nums[i] + nums[left] + nums[right]
			if r == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				localVisited[nums[left]] = true
				localVisited[nums[right]] = true
				left++
				right--
			} else if r > 0 {
				right--
			} else if r < 0 {
				left++
			}
		}
	}
	return result
}

func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(nums, low, high)
	quickSort(nums, low, p-1)
	quickSort(nums, p+1, high)
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	// 找到第一个比pivot大的位置
	p := low
	for i := low; i < high; i++ {
		if nums[i] < pivot {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	nums[p], nums[high] = nums[high], nums[p]
	return p
}
