package datastructure

type Stack struct {
	length int
	arr    []interface{}
}

func (stack *Stack) Empty() bool {
	return stack.length == 0
}
func (stack *Stack) Len() int {
	return stack.length
}
func (stack *Stack) Push(i interface{}) {
	stack.arr = append(stack.arr, i)
	stack.length++
}
func (stack *Stack) Pop() interface{} {
	if stack.length == 0 {
		return nil
	}
	i := stack.arr[stack.length-1]
	stack.arr = stack.arr[0 : stack.length-1]
	stack.length--
	return i
}

func (stack *Stack) Peek() interface{} {
	if stack.length == 0 {
		return nil
	}
	return stack.arr[stack.length-1]
}
