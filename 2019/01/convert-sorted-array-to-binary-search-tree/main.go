package main

import "fmt"

// 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

// 示例:

// 给定有序数组: [-10,-3,0,5,9],

// 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

//       0
//      / \
//    -3   9
//    /   /
//  -10  5

func main() {
	a := []int{-10, -3, 0, 5, 9}
	result := sortedArrayToBST(a)
	fmt.Println(result)
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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return build(nums, 0, len(nums)-1)
}

func build(arr []int, low, high int) *TreeNode {
	// 递归退出条件
	if low > high {
		return nil
	}
	mid := (low + high) / 2
	root := &TreeNode{Val: arr[mid]}
	root.Left = build(arr, low, mid-1)
	root.Right = build(arr, mid+1, high)
	return root
}
