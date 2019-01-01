package main

import "fmt"

// 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

// 例如：
// 给定二叉树 [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其自底向上的层次遍历为：

// [
//   [15,7],
//   [9,20],
//   [3]
// ]

func main() {
	a := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	fmt.Println(levelOrderBottom(a))
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
// 递归
func levelOrderBottom(root *TreeNode) [][]int {
	var result = make([][]int, 0, 0)
	result = insert(result, 1, root)
	// 倒序
	var reverse = make([][]int, 0, len(result))
	for i := len(result) - 1; i >= 0; i-- {
		reverse = append(reverse, result[i])
	}
	return reverse
}

func insert(result [][]int, depth int, node *TreeNode) [][]int {
	if node == nil {
		return result
	}
	if len(result) < depth {
		result = append(result, []int{})
	}
	arr := result[depth-1]
	arr = append(arr, node.Val)
	result[depth-1] = arr
	next := depth + 1
	result = insert(result, next, node.Left)
	result = insert(result, next, node.Right)
	return result
}
