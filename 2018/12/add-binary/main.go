package main

import "fmt"

// 给定两个二进制字符串，返回他们的和（用二进制表示）。

// 输入为非空字符串且只包含数字 1 和 0。

// 示例 1:

// 输入: a = "11", b = "1"
// 输出: "100"
// 示例 2:

// 输入: a = "1010", b = "1011"
// 输出: "10101"

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	// 从最低位开始计算
	var (
		as     = []rune(a)
		bs     = []rune(b)
		ai     = len(as) - 1 // a的index
		bi     = len(bs) - 1 // b的index
		carry  bool          // 是否进位
		result []rune        // 结果
	)
	for ai >= 0 || bi >= 0 {
		var x, y rune
		if ai < 0 {
			x = '0'
		} else {
			x = as[ai]
		}
		if bi < 0 {
			y = '0'
		} else {
			y = bs[bi]
		}
		// 加法计算,并且写入slice中
		var r rune
		r, carry = plus(x, y, carry)
		result = append(result, r)
		ai--
		bi--
	}
	if carry {
		result = append(result, '1')
	}
	// 翻转
	var reverse = make([]rune, 0, len(result))
	for i := len(result) - 1; i >= 0; i-- {
		reverse = append(reverse, result[i])
	}
	return string(reverse)
}

// plus 计算结果以及是否进位
func plus(a, b rune, carry bool) (rune, bool) {
	if a == '0' && b == '0' {
		if carry {
			return '1', false
		}
		return '0', false
	}
	if a == '1' && b == '1' {
		if carry {
			return '1', true
		}
		return '0', true
	}
	if carry {
		return '0', true
	}
	return '1', false
}
