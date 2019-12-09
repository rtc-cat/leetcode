package tree

import (
	"math"
	"testing"
)

// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
// 假设一个二叉搜索树具有如下特征：
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
// 示例 1:

// 输入:
//     2
//    / \
//   1   3
// 输出: true
// 示例 2:

// 输入:
//     5
//    / \
//   1   4
//      / \
//     3   6
// 输出: false
// 解释: 输入为: [5,1,4,null,null,3,6]。
//      根节点的值为 5 ，但是其右子节点值为 4 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/validate-binary-search-tree

func TestIsValidBST(t *testing.T) {
	testCase := &TreeNode{
		Val:  10,
		Left: &TreeNode{Val: 5},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 20},
		},
	}
	t.Log(isValidBST(testCase))
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	val := root.Val
	if val <= min {
		return false
	}
	if val >= max {
		return false
	}
	if !isValid(root.Right, val, max) {
		return false
	}
	if !isValid(root.Left, min, val) {
		return false
	}
	return true
}
