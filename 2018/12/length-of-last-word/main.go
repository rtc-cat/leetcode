package main

import (
	"fmt"
)

// 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。

// 如果不存在最后一个单词，请返回 0 。

// 说明：一个单词是指由字母组成，但不包含任何空格的字符串。

// 示例:

// 输入: "Hello World"
// 输出: 5
func main() {
	fmt.Println(lengthOfLastWord("b  a  "))
}

func lengthOfLastWord(s string) int {
	// 遍历每个字符,如果是空格则清零
	var (
		result int
		next   bool // 这个单词是否结束
	)
	for _, r := range s {
		if r != ' ' {
			if next {
				result = 0
			}
			next = false
			result++
		} else {
			next = true
		}
	}
	return result
}
