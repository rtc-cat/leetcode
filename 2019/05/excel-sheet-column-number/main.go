package main

import "fmt"

// 给定一个Excel表格中的列名称，返回其相应的列序号。

// 例如，

//     A -> 1
//     B -> 2
//     C -> 3
//     ...
//     Z -> 26
//     AA -> 27
//     AB -> 28
//     ...
// 示例 1:

// 输入: "A"
// 输出: 1
// 示例 2:

// 输入: "AB"
// 输出: 28
// 示例 3:

// 输入: "ZY"
// 输出: 701

func main() {
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(s string) int {
	// 26 进制数
	result := 0
	for i := range s {
		b := int(s[i]-'A') + 1
		// 指数运算
		n := 1
		exponent := len(s) - 1 - i
		for exponent > 0 {
			n *= 26
			exponent--
		}
		result += n * b
	}
	return result
}
