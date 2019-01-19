package tree

// Random 随机生成一个树
func Random() *Node {
	return nil
}

// Node 树节点
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func sortedArrayToBST(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	return build(nums, 0, len(nums)-1)
}

func build(arr []int, low, high int) *Node {
	// 递归退出条件
	if low > high {
		return nil
	}
	mid := (low + high) / 2
	root := &Node{Val: arr[mid]}
	root.Left = build(arr, low, mid-1)
	root.Right = build(arr, mid+1, high)
	return root
}
