package main

func main() {

}

// MyStack 栈
type MyStack struct {
	arr []int
}

// Constructor Initialize your data structure here.
func Constructor() MyStack {
	return MyStack{
		arr: make([]int, 0, 0),
	}
}

// Push element x onto stack.
func (s *MyStack) Push(x int) {
	s.arr = append(s.arr, x)
}

// Pop the element on top of the stack and returns that element.
func (s *MyStack) Pop() int {
	result := s.Top()
	s.arr = s.arr[:len(s.arr)-1]
	return result
}

// Top Get the top element.
func (s *MyStack) Top() int {
	if s.Empty() {
		return 0
	}
	return s.arr[len(s.arr)-1]
}

// Empty Returns whether the stack is empty.
func (s *MyStack) Empty() bool {
	if len(s.arr) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
