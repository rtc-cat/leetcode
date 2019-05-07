package main

import (
	"fmt"
)

// 给定一个链表，判断链表中是否有环。

// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

// 示例 1：

// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。

// 示例 2：

// 输入：head = [1,2], pos = 0
// 输出：true
// 解释：链表中有一个环，其尾部连接到第一个节点。

// 示例 3：

// 输入：head = [1], pos = -1
// 输出：false
// 解释：链表中没有环。

// 进阶：

// 你能用 O(1)（即，常量）内存解决此问题吗？

func main() {
	l3 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 2}
	l0 := &ListNode{Val: 0}
	l4 := &ListNode{Val: -4}
	l3.Next = l2
	l2.Next = l0
	l0.Next = l4
	l4.Next = l2
	fmt.Println(hasCycle(l3))
}

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 两个指针,一个每次前进1,一个每次前进2
	// 如果快=慢 并且 不是nil,说明有环
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		// 快指针如果到达末尾,说明没有
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
