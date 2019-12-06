package stack

import (
	ds "github.com/krizss/leetcode/datastructure"
	"testing"
)

// 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

// 示例 1:

// 输入: "(()"
// 输出: 2
// 解释: 最长有效括号子串为 "()"
// 示例 2:

// 输入: ")()())"
// 输出: 4
// 解释: 最长有效括号子串为 "()()"

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-valid-parentheses

func TestLongestValidParentheses(t *testing.T) {
	testCases := []struct{ input string }{
		{"())"},
		{"()"},
		{")()())"},
		{"(()())"},
		{"(()"},
		{")()())"},
		{"()(())"},
	}
	for _, tc := range testCases {
		t.Log(longestValidParentheses(tc.input))
	}
}

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	var max int
	// 每一个数字表示当前位置之前的最优结果
	dp := make([]int, len(s))
	// 状态转移方程
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if s[i-1] == ')' {
				// 计算上一个括号的index
				lastLeft := i - dp[i-1] - 1
				if lastLeft >= 0 {
					if s[lastLeft] == '(' {
						if lastLeft >= 1 {
							// 那就应该是这个index左边的值加上现在的值
							dp[i] = dp[i-1] + dp[lastLeft-1] + 2
						} else {
							dp[i] = dp[i-1] + 2
						}
					}
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func longestValidParenthesesWithStack(s string) int {
	stack := &ds.Stack{}
	var max int
	stack.Push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.Empty() {
				stack.Push(i)
			} else {
				if max < i-stack.Peek() {
					max = i - stack.Peek()
				}
			}
		}
	}
	return max
}
