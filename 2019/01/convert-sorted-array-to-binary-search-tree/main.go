package main

import (
	"fmt"
)

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
	fmt.Println(levelOrderBottom(result))
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
// 因为是排序数组,所以中间位置的作为root,依次计算左右节点对应的index
func sortedArrayToBST(nums []int) *TreeNode {
	rootIndex := len(nums) / 2
	rootNode := &TreeNode{Val: nums[rootIndex]}
	var (
		left        = rootIndex / 2
		right       = (len(nums) - rootIndex) / 2
		currentNode = rootNode
	)
	// 计算左节点和右节点
	for {
		if left == 1 {
			break
		}
		currentNode.Left = &TreeNode{Val: nums[left]}
		currentNode.Right = &TreeNode{Val: nums[right]}
	}
	return rootNode
}

func setCurrentNode(current *TreeNode, currentIndex, length int) *TreeNode {

	return current
}

// 为了print
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
