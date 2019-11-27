package treeutil

import (
	"sort"

	"github.com/krizss/leetcode/util/arrayutil"
)

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Load 从数组变为二叉树
func Load(nums []int) *TreeNode {
	t := sortedArrayToBST(nums)
	return t
}

// Random 随机生成一个树
func Random(max, n int) *TreeNode {
	arr := arrayutil.Random(max, n)
	sort.Ints(arr)
	return Load(arr)
}

// Convert 转变格式
func (t *TreeNode) Convert() [][]int {
	return levelOrderBottom(t)
}

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

func levelOrderBottom(root *TreeNode) [][]int {
	var result = make([][]int, 0, 0)
	result = insert(result, 1, root)
	return result
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
