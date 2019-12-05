package linklist

import (
	"testing"
)

// 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

// 示例:

// 输入:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 输出: 1->1->2->3->4->4->5->6

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-k-sorted-lists

func TestMergeKLists(t *testing.T) {
	t.Log(mergeKLists([]*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	}))
}

// 分治
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) < 2 {
		return lists[0]
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	m := (left + right) / 2
	return mergeTwoList(merge(lists, left, m), merge(lists, m+1, right))
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoList(l1, l2.Next)
		return l2
	}
}
