package main

import "fmt"

// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

// 示例 1:

// 输入: 1->1->2
// 输出: 1->2
// 示例 2:

// 输入: 1->1->2->3->3
// 输出: 1->2->3

func main() {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	result := deleteDuplicates(a)
	for {
		if result == nil {
			break
		}
		fmt.Println(result.Val)
		result = result.Next
	}
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 创建一个已有数据的重复记录
	var (
		duplicated map[int]bool
		current    *ListNode
		last       *ListNode
	)
	duplicated = make(map[int]bool)
	duplicated[head.Val] = true
	last = head
	current = head.Next
	for current != nil {
		if !duplicated[current.Val] {
			duplicated[current.Val] = true
			last = current
			current = current.Next
			continue
		}
		//如果重复,将上一个的指针引用,直接指向下一个
		// 空指针判断
		current = current.Next
		last.Next = current
	}
	return head
}
