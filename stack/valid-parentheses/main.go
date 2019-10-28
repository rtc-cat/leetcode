package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/valid-parentheses/

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。

// 示例 1:

// 输入: "()"
// 输出: true
// 示例 2:

// 输入: "()[]{}"
// 输出: true
// 示例 3:

// 输入: "(]"
// 输出: false
// 示例 4:

// 输入: "([)]"
// 输出: false
// 示例 5:

// 输入: "{[]}"
// 输出: true

func main() {
	fmt.Println(isValid(""))
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	//
	var left []rune
	for _, r := range s {
		switch r {
		case '(':
			left = append(left, ')')
		case '[':
			left = append(left, ']')
		case '{':
			left = append(left, '}')
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			if len(left) == 0 {
				return false
			}
			if !(left[len(left)-1] == r) {
				return false
			}
			left = left[:len(left)-1]
		}
	}
	return len(left) == 0
}
