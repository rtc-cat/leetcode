package main

// 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

// push(x) -- 将元素 x 推入栈中。
// pop() -- 删除栈顶的元素。
// top() -- 获取栈顶元素。
// getMin() -- 检索栈中的最小元素。
// 示例:

// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.

func main() {

}

// MinStack 最小栈
// 使用链表实现,每个节点保存当前最小值
type MinStack struct {
	Head *Node // 最后的那个节点
}

// Node 链表节点
type Node struct {
	Val  int
	Min  int
	Next *Node
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{
		Head: nil,
	}
}

// Push 压栈
func (ms *MinStack) Push(x int) {
	if ms.Head == nil {
		ms.Head = &Node{Val: x, Min: x}
		return
	}
	node := &Node{Val: x}
	min := ms.Head.Min
	if x < min {
		node.Min = x
	} else {
		node.Min = min
	}
	// 新节点的下一个是之前的节点
	node.Next = ms.Head
	// 把指针指向新节点
	ms.Head = node
}

// Pop 出栈
func (ms *MinStack) Pop() {
	if ms.Head == nil {
		return
	}
	// 把指针指向下一个
	ms.Head = ms.Head.Next
}

// Top 获取栈顶元素
func (ms *MinStack) Top() int {
	if ms.Head == nil {
		return 0
	}
	return ms.Head.Val
}

// GetMin 检索栈中最小元素
func (ms *MinStack) GetMin() int {
	if ms.Head == nil {
		return 0
	}
	return ms.Head.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
