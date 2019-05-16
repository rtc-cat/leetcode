package main

// 删除链表中等于给定值 val 的所有节点。

// 示例:

// 输入: 1->2->6->3->4->5->6, val = 6
// 输出: 1->2->3->4->5

func main() {

}

// ListNode 节点
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
func removeElements(head *ListNode, val int) *ListNode {
	first := new(ListNode)
	first.Next = head
	pre := first

	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
		} else {
			pre = head
		}
		head = head.Next
	}

	return first.Next
}
