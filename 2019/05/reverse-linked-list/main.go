package main

// 反转一个单链表。

// 示例:

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

func main() {

}

// ListNode  节点
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
// 循环的方式
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	current := head
	for current != nil {
		// 保存next
		next := current.Next
		// 将current.Next指向pre
		current.Next = pre
		// 将pre指向current
		pre = current
		// 将current执行next
		current = next
	}
	return pre
}

// 递归
// 返回的是之后的节点已经被反转过,那么只需要将
func reverseListByRecursion(head *ListNode) *ListNode {
	// 结束条件
	if head == nil || head.Next == nil {
		return head
	}
	// 这个是已经反转好的
	p := reverseListByRecursion(head.Next)
	// 将下一个反转
	head.Next.Next = head
	// 断开下一个
	head.Next = nil
	return p
}
