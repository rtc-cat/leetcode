package main

import (
	"fmt"

	. "github.com/krizss/leetcode/utils/treeutil"
)

// 给定一个二叉树，返回所有从根节点到叶子节点的路径。

// 说明: 叶子节点是指没有子节点的节点。

// 示例:

// 输入:

//    1
//  /   \
// 2     3
//  \
//   5

// 输出: ["1->2->5", "1->3"]

// 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

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
func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0, 0)
	path(&result, "", root)
	return result
}

func path(result *[]string, current string, root *TreeNode) {
	if root == nil {
		return
	}

	//
	if len(current) == 0 {
		current = fmt.Sprintf("%d", root.Val)
	} else {
		current = fmt.Sprintf("%s->%d", current, root.Val)
	}
	// 递归条件
	if root.Left == nil && root.Right == nil {
		*result = append(*result, current)
		return
	}
	path(result, current, root.Left)
	path(result, current, root.Right)
}
