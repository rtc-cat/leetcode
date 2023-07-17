package treeutil

import (
	"testing"
)

func TestTreeNode(t *testing.T) {
	tree := Load([]int{1, 2, 3, 4})
	t.Logf("%+v", tree)
	t.Logf("%v", tree.Convert())
}

func TestRandom(t *testing.T) {
	tree := Random(100, 10)
	t.Logf("%v", tree.Convert())
}
