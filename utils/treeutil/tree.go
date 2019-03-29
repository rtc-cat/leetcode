package treeutil

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Load 从数组变为二叉树
func (t *TreeNode) Load(nums []int) {
	if len(nums) == 0 {
		t = nil
		return
	}
	t = build(nums, 0, len(nums)-1)
}

// Random 随机生成一个树
func Random() *TreeNode {
	return nil
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
