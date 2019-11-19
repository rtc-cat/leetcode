package array

import "testing"

// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

// 你可以假设 nums1 和 nums2 不会同时为空。

// 示例 1:

// nums1 = [1, 3]
// nums2 = [2]

// 则中位数是 2.0
// 示例 2:

// nums1 = [1, 2]
// nums2 = [3, 4]

// 则中位数是 (2 + 3)/2 = 2.5

// 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return getMid(nums2)
	}
	if len(nums2) == 0 {
		return getMid(nums1)
	}
	// 归并排序
	var result []int
	for i, j := 0, 0; ; {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] > nums2[j] {
				result = append(result, nums2[j])
				j++
			} else {
				result = append(result, nums1[i])
				i++
			}
		} else if i < len(nums1) && j >= len(nums2) {
			result = append(result, nums1[i])
			i++
		} else if i >= len(nums1) && j < len(nums2) {
			result = append(result, nums2[j])
			j++
		} else {
			break
		}
	}
	return getMid(result)
}

func getMid(nums []int) float64 {
	l := len(nums)
	if l%2 == 0 {
		var left, right int
		right = l / 2
		left = right - 1
		return float64(nums[left]+nums[right]) / 2.0
	} else {
		return float64(nums[l/2])
	}
}
