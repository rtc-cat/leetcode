package main

// 使用栈实现队列的下列操作：

// push(x) -- 将一个元素放入队列的尾部。
// pop() -- 从队列首部移除元素。
// peek() -- 返回队列首部的元素。
// empty() -- 返回队列是否为空。
// 示例:

// MyQueue queue = new MyQueue();

// queue.push(1);
// queue.push(2);
// queue.peek();  // 返回 1
// queue.pop();   // 返回 1
// queue.empty(); // 返回 false
// 说明:

// 你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。

func main() {

}

// MyQueue  队列
type MyQueue struct {
	arr []int
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		arr: make([]int, 0, 0),
	}
}

// Push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	q.arr = append(q.arr, x)
}

// Pop Removes the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	result := q.Peek()
	q.arr = q.arr[1:]
	return result
}

// Peek Get the front element.
func (q *MyQueue) Peek() int {
	if q.Empty() {
		return 0
	}
	return q.arr[0]
}

// Empty Returns whether the queue is empty.
func (q *MyQueue) Empty() bool {
	if len(q.arr) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
