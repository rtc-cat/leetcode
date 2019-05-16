package main

import "fmt"

// 编写一个算法来判断一个数是不是“快乐数”。

// 一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。

// 示例:

// 输入: 19
// 输出: true
// 解释:
// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 02 = 1

func main() {
	fmt.Println(isHappy(19))
}

// 不管是不是快乐数,最终都会进入循环
// 只需要判断进入循环时,结果是不是1就行了
func isHappy(n int) bool {
	// 保存出现过的所有数字
	m := make(map[int]bool)
	m[n] = true
	for {
		n = getNumbers(n)
		if n == 1 {
			return true
		}
		if m[n] {
			return false
		}
		m[n] = true
	}
}

// 解析数字的每一位,然后计算平方和
// 商: quotient
// 余数: remainder
func getNumbers(n int) int {
	var result = 0
	for n > 0 {
		remainder := n % 10
		result += remainder * remainder
		n /= 10
	}
	return result
}
