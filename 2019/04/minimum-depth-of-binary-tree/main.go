package main

import (
	. "github.com/krizss/leetcode/utils/treeutil"
)

// 给定一个二叉树，找出其最小深度。

// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

// 说明: 叶子节点是指没有子节点的节点。

// 示例:

// 给定二叉树 [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回它的最小深度  2.

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	// 递归退出条件
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	ld := minDepth(root.Left)
	rd := minDepth(root.Right)
	if ld == 0 {
		return rd + 1
	}
	if rd == 0 {
		return ld + 1
	}
	if ld < rd {
		return ld + 1
	}
	return rd + 1
}
