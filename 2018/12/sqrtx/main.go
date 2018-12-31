package main

import (
	"fmt"
	"math"
)

// 实现 int sqrt(int x) 函数。

// 计算并返回 x 的平方根，其中 x 是非负整数。

// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

// 示例 1:

// 输入: 4
// 输出: 2
// 示例 2:

// 输入: 8
// 输出: 2
// 说明: 8 的平方根是 2.82842...,
//      由于返回类型是整数，小数部分将被舍去。
func main() {
	// pi := 3.54159
	// 向下取整
	// fmt.Println(math.Floor(pi))
	// 向上取整
	// fmt.Println(math.Ceil(pi))
	// 四舍五入
	// fmt.Println(math.Round(pi))
	// fmt.Println(math.Trunc(pi))
	// fmt.Println(mySqrt(16))
	// fmt.Println(mySqrt(2147395599))
	fmt.Println(bitSqrt(4))
}

// 牛顿迭代法
// https://www.matongxue.com/madocs/205.html
// 计算a的平方根,相当于求 f(x) = x² - a 的解
// next = last - (last² - a) / (2*last)
func mySqrt(x int) int {
	// 定义精度
	var prec = 0.00001
	var (
		last float64
		// result float64
	)
	last = 1.0
	a := float64(x)
	for {
		next := next(last, a)
		if math.Abs(next-last) < prec {
			break
		}
		last = next
	}
	return int(last)
}

func next(last float64, a float64) float64 {
	return last - (last*last-a)/(last*2)
}

// 二进制算法
// 通过二进制来计算,以256位例
func bitSqrt(x int) int {
	if x == 0 {
		return 0
	}
	// 第一步,找出占比最大的一位
	var h uint
	for {
		b := 1 << h
		if b*b > x {
			break
		}
		h++

	}
	// 这里减一是获得结果比x小的那一位的位置
	h-- // 4
	var (
		b      = int(h - 1) // 3
		result = 1 << h     // 0000 0000 0001 0000
	)
	// 获得剩下的每一位
	for b >= 0 {
		// 1 << b = 1000
		n := result | (1 << uint(b))
		// 0000 0000 0001 1000 24
		// 0000 0000 0001 0100 20
		// 0000 0000 0001 0010 18
		// 0000 0000 0001 0001 17
		// 0000 0000 0001 0000 16
		if n*n <= x {
			result |= (1 << uint(b))
		}
		b--
	}
	return int(result)
}
