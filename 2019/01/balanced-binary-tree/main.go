package main

import (
	"math"

	. "github.com/krizss/leetcode/utils/treeutil"
)

// 给定一个二叉树，判断它是否是高度平衡的二叉树。

// 本题中，一棵高度平衡二叉树定义为：

// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

// 示例 1:

// 给定二叉树 [3,9,20,null,null,15,7]

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回 true 。

// 示例 2:

// 给定二叉树 [1,2,2,3,3,null,null,4,4]

//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4
// 返回 false 。
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
func isBalanced(root *TreeNode) bool {
	b, _ := check(root)
	return b
}

// 判断节点是否有子节点
func check(node *TreeNode) (bool, int) {
	// 退出条件
	if node == nil {
		return true, 0
	}
	// 检查两个节点是否平衡,以及深度
	lb, ld := check(node.Left)
	rb, rd := check(node.Right)
	// 如果这俩有一个不平衡,那么就返回false
	if !lb || !rb {
		return false, 0
	}
	// 检查这俩树的深度,如果差距不是1,就返回false
	if math.Abs(float64(ld-rd)) > 1 {
		return false, 0
	}
	max := ld
	if rd > ld {
		max = rd
	}
	return true, max + 1
}
