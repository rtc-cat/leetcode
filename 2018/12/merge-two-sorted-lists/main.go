package main

import (
	"fmt"
)

// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 示例：

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
func main() {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	b := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := mergeTwoLists(a, b)
	for {
		if result == nil {
			break
		}
		fmt.Println(result.Val)
		result = result.Next
	}
}

// ListNode 链表
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// result定义为空,从它的Next开始
	// current用来操作当前节点,指定下一个节点地址
	var (
		result  = &ListNode{}
		current = result
		v       = 0
	)

	for l1 != nil && l2 != nil {
		// 对比大小
		if l1.Val < l2.Val {
			v = l1.Val
			l1 = l1.Next
		} else {
			v = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return result.Next
}
