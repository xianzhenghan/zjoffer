package zj034_二叉树中和为某一值的路径

import (
	"fmt"
	"testing"
)

// go test -v zj034.go zj034_test.go
func TestPath(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}
	r := pathTarget(tree, 12)
	fmt.Printf("r=%v", r)
}
