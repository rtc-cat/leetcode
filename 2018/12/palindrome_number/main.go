package main

import "fmt"

// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

// 示例 1:

// 输入: 121
// 输出: true
// 示例 2:

// 输入: -121
// 输出: false
// 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 示例 3:

// 输入: 10
// 输出: false
// 解释: 从右向左读, 为 01 。因此它不是一个回文数。
// 进阶:

// 你能不将整数转为字符串来解决这个问题吗？
func main() {
	fmt.Println(isPalindrome(112233))
}

// ? 其实可以只翻转一半的数字
func isPalindrome(x int) bool {
	// 如果是负数,那一定不是
	if x < 0 {
		return false
	}
	// 如果这个数字只有一位那肯定是
	if x < 10 {
		return true
	}
	// 大于10需要将每一位拆出来
	var (
		numbers   []int // 每一位数字
		remainder int   //余数
	)
	for {
		remainder = x % 10
		numbers = append(numbers, remainder)
		x = x / 10
		if x == 0 {
			break
		}
	}
	// 定义两个指针,从两边开始,同时向中遍历
	var (
		left  = 0
		right = len(numbers) - 1
	)
	for i := 0; i < len(numbers)/2; i++ {
		if numbers[left] != numbers[right] {
			return false
		}
		left++
		right--
	}
	return true
}
