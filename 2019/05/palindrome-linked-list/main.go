package main

// 请判断一个链表是否为回文链表。

// 示例 1:

// 输入: 1->2
// 输出: false
// 示例 2:

// 输入: 1->2->2->1
// 输出: true
// 进阶：
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

func main() {

}

// ListNode  链表节点
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
func isPalindrome(head *ListNode) bool {
	// 定义两个指针
	if head == nil || head.Next == nil {
		return true
	}
	pre := &ListNode{Next: head}
	left := head
	slow, fast := pre, pre
	for {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
		// 如果fast到达终点,说明slow已经是中间节点
		if fast == nil || fast.Next == nil {
			right := reverse(slow.Next)
			slow.Next = nil
			// 然后比较左右两边的值
			for left != nil && right != nil {
				if left.Val != right.Val {
					return false
				}
				left = left.Next
				right = right.Next
			}
			return true
		}
	}
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	// 到达链表末端
	if head == nil || head.Next == nil {
		return head
	}
	// 把之后的先反转好
	p := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
