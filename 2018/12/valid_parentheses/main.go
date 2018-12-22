package main

import "fmt"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。

// 示例 1:

// 输入: "()"
// 输出: true
// 示例 2:

// 输入: "()[]{}"
// 输出: true
// 示例 3:

// 输入: "(]"
// 输出: false
// 示例 4:

// 输入: "([)]"
// 输出: false
// 示例 5:

// 输入: "{[]}"
// 输出: true
func main() {
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	// 数据结构应该是栈,遇到一个左括号变压入栈中
	// 当遇到右括号时,与栈中最上面的括号类型判断,如果不匹配则直接返回false
	// 如果匹配的话,就将栈中最上面的弹出,之后继续匹配
	// 如果最后栈中没有括号,则返回true
	// 使用map替代,指定初始长度
	m := make(map[int]rune, len(s)/2)
	// index表示栈的最上层索引
	index := 0
	for _, r := range s {
		switch r {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			m[index] = r
			index++
		case ')':
			i := index - 1
			if m[i] != '(' {
				return false
			}
			delete(m, i)
			index--
		case '}':
			i := index - 1
			if m[i] != '{' {
				return false
			}
			delete(m, i)
			index--
		case ']':
			i := index - 1
			if m[i] != '[' {
				return false
			}
			delete(m, i)
			index--
		}
	}
	if len(m) == 0 {
		return true
	}
	return false
}
