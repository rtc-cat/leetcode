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
	// 分别从l1,和l2开始,对比,将更小的那个数字合并到新的结果中
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var (
		result  = &ListNode{}
		current = result
	)

	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil && l2 != nil {
			current.Next = l2
			current = current.Next
			l2 = l2.Next
			continue
		}
		if l1 != nil && l2 == nil {
			current.Next = l1
			current = current.Next
			l1 = l1.Next
			continue
		}
		// 对比大小
		if l1.Val < l2.Val {
			current.Next = l1
			current = current.Next
			l1 = l1.Next
		} else {
			current.Next = l2
			current = current.Next
			l2 = l2.Next
		}
	}
	return result.Next
}
