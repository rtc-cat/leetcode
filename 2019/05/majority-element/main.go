package main

import "fmt"

// 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在众数。

// 示例 1:

// 输入: [3,2,3]
// 输出: 3
// 示例 2:

// 输入: [2,2,1,1,1,2,2]
// 输出: 2

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

// 摩尔投票算法
// 从第一个数开始count=1，遇到相同的就加1，遇到不同的就减1，减到0就重新换个数开始计数，总能找到最多的那个
func majorityElement(nums []int) int {
	count := 1
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == last {
			count++
		} else {
			count--
			if count == 0 {
				last = nums[i]
				count = 1
			}
		}
	}
	return last
}
