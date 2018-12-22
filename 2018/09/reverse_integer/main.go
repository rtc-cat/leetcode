package main

import (
	"fmt"
)

// 给定一个 32 位有符号整数，将整数中的数字进行反转。
// 注意:
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。
// 根据这个假设，如果反转后的整数溢出，则返回 0。
func main() {
	fmt.Println(re(11111))
}

func re(x int) int {
	// 除以10,商为下一次的被除数,余数为下一个反转数字
	// 获取每一位数字
	const MaxInt = 0x7FFFFFFF
	// const MinInt = 0x80000000
	var (
		negative  bool  //	是否是负数
		numbers   []int //	最后的结果数字
		remainder int   //	余数
	)
	// 先将x转为正数
	if x < 0 {
		negative = true
		x = -x
	}
	for {
		// 开始除以10
		remainder = x % 10
		numbers = append(numbers, remainder)
		x = x / 10
		// 如果商是0说明已经是最后一位了
		if x == 0 {
			break
		}
	}
	// 根据数组里面的数字算出来结果
	var result int
	for i := len(numbers) - 1; i >= 0; i-- {
		var exponent = 1
		count := len(numbers) - 1 - i
		if count == 0 {
			result += numbers[i]
			continue
		}
		for j := 0; j < count; j++ {
			exponent *= 10
		}
		result += numbers[i] * exponent
		if result > MaxInt {
			return 0
		}
	}
	if negative {
		result = -result
	}
	return result
}
