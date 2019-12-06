package datastructure

type Stack struct {
	length int
	arr    []int
}

func (stack *Stack) Empty() bool {
	return stack.length == 0
}
func (stack *Stack) Len() int {
	return stack.length
}
func (stack *Stack) Push(i int) {
	stack.arr = append(stack.arr, i)
	stack.length++
}
func (stack *Stack) Pop() int {
	if stack.length == 0 {
		return 0
	}
	i := stack.arr[stack.length-1]
	stack.arr = stack.arr[0 : stack.length-1]
	stack.length--
	return i
}

func (stack *Stack) Peek() int {
	if stack.length == 0 {
		return 0
	}
	return stack.arr[stack.length-1]
}
