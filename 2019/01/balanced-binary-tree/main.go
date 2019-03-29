package main

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

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	// 递归退出条件 如果两边的节点的子节点都没有节点
	if root == nil {
		return true
	}
	// 如果两个子节点都是空
	if root.Left == nil && root.Right == nil {
		return true
	}
	// 如果两个都不为空
	if root.Left != nil && root.Right != nil {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	// 这里就是一个为空,另一个不为空的情况
	left := check(root.Left)
	right := check(root.Right)
	return left == right
}

// 判断节点是否有子节点
func check(node *TreeNode) bool {
	if node != nil {
		if node.Left == nil && node.Right == nil {
			return false
		}
		return true
	}
	return false
}
