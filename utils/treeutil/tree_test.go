package treeutil

import (
	"testing"
)

func TestTreeNode(t *testing.T) {
	tree := &TreeNode{}
	tree.Load([]int{1})
	t.Logf("%+v", tree)
}
