package main

// 给定两个数组，编写一个函数来计算它们的交集。

// 示例 1:

// 输入: nums1 = [1,2,2,1], nums2 = [2,2]
// 输出: [2]
// 示例 2:

// 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出: [9,4]
// 说明:

// 输出结果中的每个元素一定是唯一的。
// 我们可以不考虑输出结果的顺序。

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool, len(nums1))
	m2 := make(map[int]bool, len(nums2))
	result := make([]int, 0, len(nums1))
	// 分别记录nums1和nums2
	for _, n := range nums1 {
		m1[n] = true
	}
	for _, n := range nums2 {
		m2[n] = true
	}
	for k := range m1 {
		if m2[k] {
			result = append(result, k)
		}
	}
	return result
}
