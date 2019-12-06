package stack

import (
	"strconv"
	"testing"
)

// 根据逆波兰表示法，求表达式的值。

// 有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

// 说明：

// 整数除法只保留整数部分。
// 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
// 示例 1：

// 输入: ["2", "1", "+", "3", "*"]
// 输出: 9
// 解释: ((2 + 1) * 3) = 9
// 示例 2：

// 输入: ["4", "13", "5", "/", "+"]
// 输出: 6
// 解释: (4 + (13 / 5)) = 6
// 示例 3：

// 输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
// 输出: 22
// 解释:
//   ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/evaluate-reverse-polish-notation

func TestEvalRPN(t *testing.T) {
	testCases := []struct{ input []string }{
		{[]string{"2", "1", "+", "3", "*"}},
		{[]string{"4", "13", "5", "/", "+"}},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
	}
	for _, tc := range testCases {
		t.Log(evalRPN(tc.input))
	}
}

func evalRPN(tokens []string) int {
	// 记录数字
	digit := &Stack{}
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			bStr := digit.Pop()
			b, _ := strconv.Atoi(bStr)
			aStr := digit.Pop()
			a, _ := strconv.Atoi(aStr)
			switch token {
			case "+":
				digit.Push(strconv.Itoa(a + b))
			case "-":
				digit.Push(strconv.Itoa(a - b))
			case "*":
				digit.Push(strconv.Itoa(a * b))
			case "/":
				digit.Push(strconv.Itoa(a / b))
			}
		default:
			digit.Push(token)
		}
	}
	resultStr := digit.Pop()
	result, _ := strconv.Atoi(resultStr)
	return result
}

type Stack struct {
	length int
	arr    []string
}

func (stack *Stack) Empty() bool {
	return stack.length == 0
}
func (stack *Stack) Len() int {
	return stack.length
}
func (stack *Stack) Push(i string) {
	stack.arr = append(stack.arr, i)
	stack.length++
}
func (stack *Stack) Pop() string {
	if stack.length == 0 {
		return ""
	}
	i := stack.arr[stack.length-1]
	stack.arr = stack.arr[0 : stack.length-1]
	stack.length--
	return i
}

func (stack *Stack) Peek() string {
	if stack.length == 0 {
		return ""
	}
	return stack.arr[stack.length-1]
}
