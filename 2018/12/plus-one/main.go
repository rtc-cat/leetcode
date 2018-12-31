package main

import "fmt"

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

// 最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

// 你可以假设除了整数 0 之外，这个整数不会以零开头。

// 示例 1:

// 输入: [1,2,3]
// 输出: [1,2,4]
// 解释: 输入数组表示数字 123。
// 示例 2:

// 输入: [4,3,2,1]
// 输出: [4,3,2,2]
// 解释: 输入数组表示数字 4321。

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}
func plusOne(digits []int) []int {
	// 中间数记录
	flag := true
	for index := len(digits) - 1; index >= 0; index-- {
		if flag {
			// +1
			n := digits[index] + 1
			if n >= 10 {
				flag = true
				digits[index] = n % 10
				if index == 0 {
					result := []int{1}
					result = append(result, digits...)
					return result
				}
			} else {
				flag = false
				digits[index] = n
			}
		}
	}
	return digits
}
